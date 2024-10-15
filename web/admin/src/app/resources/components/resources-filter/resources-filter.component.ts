import { CommonModule } from '@angular/common';
import { Component, EventEmitter, Input, Output, ChangeDetectorRef, ChangeDetectionStrategy } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { TranslateModule } from '@ngx-translate/core';
import { DropdownModule } from 'primeng/dropdown';
import { OwnerType } from '../../../core/models/resources/ownerType';
import { ResourceType } from '../../../core/models/resources/resourceType';
import { Cluster } from '../../../core/models/cluster';
import { ResourceFilter } from '../../models/resourceFilter';
import { UtilsService } from '../../../shared/services/utils.service';
import { SharedModule } from '../../../shared/shared.module';

@Component({
  selector: 'app-resources-filter',
  standalone: true,
  imports: [CommonModule, FormsModule, TranslateModule, DropdownModule, FormsModule, SharedModule],
  templateUrl: './resources-filter.component.html',
  styleUrl: './resources-filter.component.scss',
})
export class ResourcesFilterComponent {
  @Output() filterChanged = new EventEmitter<ResourceFilter>();

  @Input() owners: OwnerType[] = [];
  selectedOwner: OwnerType;

  @Input() resourceTypes: ResourceType[] = [];
  selectedResourceType: ResourceType;

  @Input() clusters: Cluster[] = [];
  selectedCluster: Cluster;

  @Input() clusterId: string = undefined;

  @Input() set filter(value: ResourceFilter | undefined) {
    if (!value) {
      return;
    }

    if (this.utilsService.isEqual(value, this.filter)) {
      return;
    }

    this.selectedResourceType = this.resourceTypes.find((x) => x.apiVersion === value.apiVersion && x.kind === value.kind);
    this.selectedOwner = this.owners.find((x) => x.scope === value.scope);
    if (this.selectedOwner?.clusterSpesific) {
      this.showClusterDropdown = true;
    }
    if (this.clusterId) {
      this.selectedCluster = this.clusters.find((x) => x.clusterId === this.clusterId);
    } else {
      this.selectedCluster = this.clusters.find((x) => x.clusterId === value.clusterId);
    }
  }

  showClusterDropdown = false;

  constructor(private utilsService: UtilsService) {}

  selectedResourceTypeChange(event: any): void {
    this.selectedResourceType = event?.value;
    this.emit();
  }

  selectedOwnerChange(event: any): void {
    this.selectedOwner = event?.value;
    if (this.selectedOwner?.clusterSpesific) {
      this.showClusterDropdown = true;
      this.selectedCluster = undefined;
    } else {
      this.showClusterDropdown = false;
      this.emit();
    }
  }

  selectedClusterIdChange(event: any): void {
    this.selectedCluster = event?.value;
    this.emit();
  }

  resetFilter(): void {
    this.selectedResourceType = undefined;
    this.selectedOwner = undefined;
    this.selectedCluster = undefined;
    this.showClusterDropdown = false;
    this.emit();
  }

  private emit(): void {
    const filter: ResourceFilter = {
      apiVersion: this.selectedResourceType?.apiVersion,
      kind: this.selectedResourceType?.kind,
      scope: this.selectedOwner?.scope,
      subject: this.selectedOwner?.subject,
      clusterId: this.selectedCluster?.clusterId,
    };

    if (this.clusterId) {
      filter.clusterId = this.clusterId;
      filter.scope = this.owners.find((x) => x.scope === 'cluster')?.scope;
    }

    this.filterChanged.emit(filter);
  }
}

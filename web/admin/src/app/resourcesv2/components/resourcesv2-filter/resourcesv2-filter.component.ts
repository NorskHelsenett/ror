import { Component, EventEmitter, inject, Input, Output } from '@angular/core';
import { ResourcesV2QueryService } from '../../services/resources-v2-query.service';
import { DropdownModule } from 'primeng/dropdown';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { TranslateModule } from '@ngx-translate/core';
import { Cluster } from '../../../core/models/cluster';
import { ResourceType } from '../../../core/models/resources/resourceType';
import { OwnerType } from '../../../core/models/resources/ownerType';
import { ResourceQuery } from '../../../core/models/resource-query';

@Component({
  selector: 'app-resourcesv2-filter',
  standalone: true,
  imports: [CommonModule, FormsModule, ReactiveFormsModule, TranslateModule, DropdownModule],
  templateUrl: './resourcesv2-filter.component.html',
  styleUrl: './resourcesv2-filter.component.scss',
})
export class Resourcesv2FilterComponent {
  @Input() clusterId: string = undefined;
  @Input() clusters: Cluster[] = [];
  selectedCluster: Cluster;

  @Input() resourceTypes: ResourceType[] = [];
  selectedResourceType: ResourceType;
  showClusterDropdown = false;

  @Output() filterChange = new EventEmitter<ResourceQuery>();
  @Output() reset = new EventEmitter<void>();

  private resourcesV2QueryService = inject(ResourcesV2QueryService);

  selectedResourceTypeChange(event: any): void {
    this.selectedResourceType = event?.value;
    var filterchange = {
      versionkind: {
        Version: event?.value?.apiVersion,
        Kind: event?.value?.kind,
      },
    };
    this.updateAndNotify(filterchange);
  }

  selectedOwnerChange(event: any): void {
    var filterchange = {
      ownerrefs: [
        {
          scope: event?.value?.scope,
          subject: event?.value?.subject,
        },
      ],
    };
    this.updateAndNotify(filterchange);
  }

  selectedClusterIdChange(event: any): void {
    this.selectedCluster = event?.value;
    var filterchange = {
      ownerrefs: {
        subject: this.selectedCluster?.clusterId,
      },
    };
    this.updateAndNotify(filterchange);
  }

  resetFilter(): void {
    this.selectedResourceType = undefined;
    this.selectedCluster = undefined;
    this.showClusterDropdown = false;
    this.resourcesV2QueryService.clearQuery();
    this.reset.emit();
  }

  private updateAndNotify(query: any): void {
    this.resourcesV2QueryService.updateQuery(query);
    this.filterChange.emit(query);
  }
}

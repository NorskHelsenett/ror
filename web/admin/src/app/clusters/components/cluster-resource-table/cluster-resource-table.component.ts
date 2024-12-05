import { TypesService } from './../../../resources/services/types.service';
import { Component, EventEmitter, inject, Input, OnInit, Output } from '@angular/core';
import { TranslateModule } from '@ngx-translate/core';
import { ResourcesV2ListComponent } from '../../../resourcesv2/components/resources-v2-list/resources-v2-list.component';
import { Resourcesv2FilterComponent } from '../../../resourcesv2/components/resourcesv2-filter/resourcesv2-filter.component';
import { ResourceType } from '../../../core/models/resources/resourceType';
import { ResourcesV2QueryService } from '../../../resourcesv2/services/resources-v2-query.service';

@Component({
  selector: 'app-cluster-resource-table',
  standalone: true,
  imports: [TranslateModule, Resourcesv2FilterComponent, ResourcesV2ListComponent],
  templateUrl: './cluster-resource-table.component.html',
  styleUrl: './cluster-resource-table.component.scss',
  providers: [TypesService],
})
export class ClusterResourceTableComponent implements OnInit {
  @Input() cluster: any | undefined;
  @Output() resourceSelected = new EventEmitter<any>();

  resourceTypes: ResourceType[] | undefined;

  private typesService = inject(TypesService);
  private resourcesv2QueryService = inject(ResourcesV2QueryService);

  ngOnInit() {
    this.resourcesv2QueryService.clearQuery();
    this.resourceTypes = this.typesService.getResourceTypes();
    if (this.resourceTypes?.length > 0) {
      this.resourceTypes = this.resourceTypes?.sort((a: ResourceType, b: ResourceType) => a.kind.localeCompare(b.kind));
    }
  }

  showSelectedResource(resource: any) {
    this.resourceSelected.emit(resource);
  }
}

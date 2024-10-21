import { ResourcesV2QueryService } from './../../services/resources-v2-query.service';
import { Component, inject, OnInit } from '@angular/core';
import { TranslateModule } from '@ngx-translate/core';
import { Resourcesv2FilterComponent } from '../../components/resourcesv2-filter/resourcesv2-filter.component';
import { ClustersService } from '../../../core/services/clusters.service';
import { catchError, Observable } from 'rxjs';
import { PaginationResult } from '../../../core/models/paginatedResult';
import { Cluster } from '../../../core/models/cluster';
import { TypesService } from '../../../resources/services/types.service';
import { ResourceType } from '../../../core/models/resources/resourceType';
import { OwnerType } from '../../../core/models/resources/ownerType';
import { ResourcesV2ListComponent } from '../../components/resources-v2-list/resources-v2-list.component';
import { AsyncPipe } from '@angular/common';
import { SidebarModule } from 'primeng/sidebar';
import { ResourceV2DetailsComponent } from '../../components/resource-v2-details/resource-v2-details.component';

@Component({
  selector: 'app-resources',
  standalone: true,
  imports: [TranslateModule, ResourcesV2ListComponent, Resourcesv2FilterComponent, AsyncPipe, SidebarModule, ResourceV2DetailsComponent],
  templateUrl: './resources.component.html',
  styleUrl: './resources.component.scss',
})
export class ResourcesComponent implements OnInit {
  clusters$: Observable<PaginationResult<Cluster>>;
  resourceTypes: ResourceType[];
  owners: OwnerType[];

  sidebarVisible = false;
  selectedResource: any;

  resourcesV2QueryService = inject(ResourcesV2QueryService);

  constructor(
    private clustersService: ClustersService,
    private typesService: TypesService,
  ) {}

  ngOnInit() {
    this.resourceTypes = this.typesService.getResourceTypes();
    if (this.resourceTypes?.length > 0) {
      this.resourceTypes = this.resourceTypes?.sort((a: ResourceType, b: ResourceType) => a.kind.localeCompare(b.kind));
    }

    this.owners = this.typesService.getOwnerTypes();
    if (this.owners?.length > 0) {
      this.owners = this.owners?.sort((a: OwnerType, b: OwnerType) => a.scope.localeCompare(b.scope));
    }

    this.clusters$ = this.clustersService
      .getByFilter({
        limit: 200,
        skip: 0,
      })
      .pipe(
        catchError((error) => {
          throw error;
        }),
      );
  }

  showSelectedResource(resource: any) {
    this.selectedResource = resource;
    this.sidebarVisible = true;
  }
}

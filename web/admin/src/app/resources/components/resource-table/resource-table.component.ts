import { CommonModule } from '@angular/common';
import { ChangeDetectorRef, Component, Input, OnDestroy, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { ActivatedRoute, Params, Router, RouterModule } from '@angular/router';
import { TranslateModule } from '@ngx-translate/core';
import { HighlightModule } from 'ngx-highlightjs';
import { HighlightLineNumbers } from 'ngx-highlightjs/line-numbers';
import { FilterMatchMode } from 'primeng/api';
import { ButtonModule } from 'primeng/button';
import { MultiSelectModule } from 'primeng/multiselect';
import { RippleModule } from 'primeng/ripple';
import { SidebarModule } from 'primeng/sidebar';
import { TableModule } from 'primeng/table';
import { ResourcesFilterComponent } from '../resources-filter/resources-filter.component';
import { ErrorComponent } from '../../../shared/components/error/error.component';
import { ResourcesService } from '../../../core/services/resources.service';
import { ColumnFactoryService } from '../../services/column-factory.service';
import { TypesService } from '../../services/types.service';
import { ResourceFilter } from '../../models/resourceFilter';
import { Observable, Subscription, catchError, map, of, tap } from 'rxjs';
import { Cluster } from '../../../core/models/cluster';
import { PaginationResult } from '../../../core/models/paginatedResult';
import { OwnerType } from '../../../core/models/resources/ownerType';
import { ResourceType } from '../../../core/models/resources/resourceType';
import { ClustersService } from '../../../core/services/clusters.service';
import { ConfigService } from '../../../core/services/config.service';
import { SharedModule } from '../../../shared/shared.module';
import { ResourceQuery } from '../../../core/models/resources-v2';

@Component({
  selector: 'app-resource-table',
  standalone: true,
  imports: [
    CommonModule,
    SharedModule,
    TranslateModule,
    RouterModule,
    TableModule,
    ErrorComponent,
    ButtonModule,
    RippleModule,
    ResourcesFilterComponent,
    MultiSelectModule,
    FormsModule,
    SidebarModule,
    HighlightModule,
    HighlightLineNumbers,
  ],
  providers: [ResourcesService, ColumnFactoryService, TypesService, Location],
  templateUrl: './resource-table.component.html',
  styleUrl: './resource-table.component.scss',
})
export class ResourceTableComponent implements OnInit, OnDestroy {
  @Input() clusterId: string = undefined;

  resourcesFetchError: any;
  resources$: Observable<any[]>;
  selectedResource: any;

  owners: OwnerType[];
  resourceTypes: ResourceType[];

  clusters$: Observable<PaginationResult<Cluster>>;
  clusters: any[] = [];
  selectedClusterId: string;
  filter: ResourceFilter = {
    scope: undefined,
    subject: undefined,
    kind: undefined,
    apiVersion: undefined,
    clusterId: undefined,
  };
  rowsPerPage = this.configService.config.rowsPerPage;
  rows = this.configService.config.rows;
  matchModeOptions: any[] = [
    {
      label: 'Contains',
      value: FilterMatchMode.CONTAINS,
    },
    {
      label: 'Equals',
      value: FilterMatchMode.EQUALS,
    },
  ];
  columnDefinitions: any[] = [];
  selectedColumns: any[] = [];

  sidebarVisible = false;

  private subscriptions = new Subscription();

  constructor(
    private changeDetector: ChangeDetectorRef,
    private activeRoute: ActivatedRoute,
    private router: Router,
    private resourcesService: ResourcesService,
    private columnFactoryService: ColumnFactoryService,
    private typesService: TypesService,
    private clustersService: ClustersService,
    private configService: ConfigService,
  ) {}

  ngOnInit(): void {
    this.owners = this.typesService.getOwnerTypes();
    if (this.owners?.length > 0) {
      this.owners = this.owners?.sort((a: OwnerType, b: OwnerType) => a.scope.localeCompare(b.scope));
    }

    this.resourceTypes = this.typesService.getResourceTypes();
    if (this.resourceTypes?.length > 0) {
      this.resourceTypes = this.resourceTypes?.sort((a: ResourceType, b: ResourceType) => a.kind.localeCompare(b.kind));
    }

    if (this.clusterId) {
      this.resourceTypes = this.resourceTypes.filter((x) => x.clusterSpecific);
    }

    this.getClusters();

    const queries = this.activeRoute.snapshot.queryParams;
    this.filter.scope = queries['scope'];
    this.filter.subject = queries['subject'];
    this.filter.clusterId = queries['clusterId'];
    this.filter.kind = queries['kind'];
    this.filter.apiVersion = queries['apiVersion'];

    this.fetchResources();
    this.changeDetector.detectChanges();
  }

  ngOnDestroy(): void {
    this.subscriptions.unsubscribe();
  }

  fetchResources(): void {
    this.selectedResource = undefined;
    this.resourcesFetchError = undefined;
    if (!this.filter) {
      this.resources$ = of([]);
      return;
    }
    let subject = this.filter.subject ? this.filter.subject : this.filter?.clusterId;
    let query: ResourceQuery = {
      versionkind: {
        Kind: this.filter?.kind,
        Group: '',
        Version: this.filter?.apiVersion,
      },
      ownerrefs: [
        {
          scope: this.filter?.scope,
          subject: subject,
        },
      ],
    };
    this.resources$ = this.resourcesService.getResources(query).pipe(
      map((rs) => {
        return rs?.resources;
      }),
      tap({
        next: (_) => {
          this.columnDefinitions = this.columnFactoryService.getColumnDefinitions(this.filter?.apiVersion, this.filter?.kind);
          this.selectedColumns = this.columnDefinitions;
          this.changeDetector.detectChanges();
        },
        error: (error) => {
          this.resourcesFetchError = error;
          this.changeDetector.detectChanges();
        },
      }),
      catchError((error) => {
        this.resourcesFetchError = error;
        this.changeDetector.detectChanges();
        throw error;
      }),
    );
  }

  getClusters(): void {
    this.clusters$ = this.clustersService
      .getByFilter({
        limit: 100,
        skip: 0,
      })
      .pipe(
        catchError((error) => {
          this.changeDetector.detectChanges();
          throw error;
        }),
      );
  }

  extractData(data: any, field: string) {
    if (!data || !field) {
      return;
    }
    return field.includes('.') ? field.split('.').reduce((acc: any, obj: any) => acc[obj], data) : data[field];
  }

  export(): any {
    const exportObjects: any[] = [];
    return exportObjects;
  }

  filterChanged(event: ResourceFilter): void {
    this.filter = event;
    this.fetchResources();
    this.updateRoute();
    this.changeDetector.detectChanges();
  }

  getGlobalSearchFields(): string[] {
    return this.columnDefinitions.map((x) => x.field);
  }

  showDetails(resource: any): void {
    this.selectedResource = resource;
    this.sidebarVisible = true;
    this.changeDetector.detectChanges();
  }

  private updateRoute(): void {
    const queryParams: Params = {
      scope: this.filter?.scope,
      subject: this.filter?.subject,
      kind: this.filter?.kind,
      apiVersion: this.filter?.apiVersion,
      clusterId: this.filter?.clusterId,
    };
    this.router.navigate([], {
      relativeTo: this.activeRoute,
      queryParams: this.filter,
      queryParamsHandling: 'merge',
    });
  }
}

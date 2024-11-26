import { ResourceQuery } from './../../../core/models/resources-v2';
import { catchError, finalize, map, Observable, share } from 'rxjs';
import { Component, OnInit, ChangeDetectorRef, inject, ChangeDetectionStrategy, Output, EventEmitter, Input } from '@angular/core';
import { Resourcesv2Service } from '../../../core/services/resourcesv2.service';
import { Resource, ResourceSet } from '../../../core/models/resources-v2';
import { SharedModule } from '../../../shared/shared.module';
import { TranslateModule } from '@ngx-translate/core';
import { CommonModule } from '@angular/common';
import { TableModule } from 'primeng/table';
import { ConfigService } from '../../../core/services/config.service';
import { FilterMatchMode } from 'primeng/api';
import { ColumnFactoryService } from '../../services/column-factory.service';
import { ResourcesV2QueryService } from '../../services/resources-v2-query.service';
import { MultiSelectModule } from 'primeng/multiselect';
import { FormsModule } from '@angular/forms';
import { AclService } from '../../../core/services/acl.service';
import { AclAccess, AclScopes } from '../../../core/models/acl-scopes';
import { DropdownModule } from 'primeng/dropdown';
import { FilterService } from '../../services/filter.service';
import { ColumnDefinition } from '../../../resources/models/columnDefinition';
import { TooltipModule } from 'primeng/tooltip';

@Component({
  selector: 'app-resources-v2-list',
  standalone: true,
  imports: [CommonModule, TranslateModule, SharedModule, TableModule, MultiSelectModule, FormsModule, DropdownModule, TooltipModule],
  templateUrl: './resources-v2-list.component.html',
  styleUrl: './resources-v2-list.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ResourcesV2ListComponent implements OnInit {
  @Input() clusterId: string | undefined;
  @Output() resourceSelected = new EventEmitter<any>();

  resourceQuery: any;
  loading = false;
  showLoadMore = true;

  resources: Resource[] = [];
  resourceSet$: Observable<ResourceSet> | undefined = undefined;
  resourceSetFetchError: any;
  matchModeOptionsString: any[] = [
    {
      label: 'Contains',
      value: FilterMatchMode.CONTAINS,
    },
    {
      label: 'Equals',
      value: FilterMatchMode.EQUALS,
    },
    {
      label: 'Not equals',
      value: FilterMatchMode.NOT_EQUALS,
    },
  ];

  matchModeOptionsInt: any[] = [
    {
      label: 'Equal',
      value: FilterMatchMode.EQUALS,
    },
    {
      label: 'Greater than',
      value: FilterMatchMode.GREATER_THAN,
    },
    {
      label: 'Less than',
      value: FilterMatchMode.LESS_THAN,
    },
    {
      label: 'Grather than or equal to',
      value: FilterMatchMode.GREATER_THAN_OR_EQUAL_TO,
    },
    {
      label: 'Less than or equal to',
      value: FilterMatchMode.LESS_THAN_OR_EQUAL_TO,
    },
  ];

  matchModeOptionsBool: any[] = [
    {
      label: 'Equal',
      value: FilterMatchMode.EQUALS,
    },
  ];

  exportFilename = 'resources';

  selectedColumns: any[] = [];
  rows = 10;
  rowsPerPage = [5, 10, 25, 50, 100];
  columnDefinitions: ColumnDefinition[] = [];

  lastLazyEvent: any;

  private changeDetector = inject(ChangeDetectorRef);
  private resourcesv2Service = inject(Resourcesv2Service);
  private columnFactoryService = inject(ColumnFactoryService);
  private configService = inject(ConfigService);
  private resourcesV2QueryService = inject(ResourcesV2QueryService);
  private aclService = inject(AclService);
  private filterService = inject(FilterService);

  adminCreate$: Observable<boolean> | undefined;
  aclFetchError: any;

  ngOnInit() {
    this.rows = this.configService.config.rows;
    this.rowsPerPage = this.configService.config.rowsPerPage;
    this.fetchAcl();

    if (this.clusterId) {
      this.exportFilename = 'resources-' + this.clusterId;
    }
  }

  refreshData() {
    this.resourceQuery = this.resourcesV2QueryService.getQuery();
    this.resourceQuery.limit = this.resources?.length;
    this.resourceQuery.offset = 0;
    this.resources = [];
    this.fetchResourceSet();
    this.changeDetector.detectChanges();
  }

  filterChanged(event: any) {
    this.showLoadMore = true;
    this.resourceQuery = this.resourcesV2QueryService.getQuery();
    this.resourceQuery.limit = this.rows;
    this.resourceQuery.offset = 0;
    this.resources = [];
    this.fetchResourceSet();
    this.changeDetector.detectChanges();
  }

  extractData(data: any, field: string) {
    if (!data || !field) {
      return;
    }
    try {
      return field.includes('.') ? field.split('.').reduce((acc: any, obj: any) => acc[obj], data) : data[field];
    } catch (error) {
      return;
    }
  }

  export(): any[] {
    const exportObjects: any[] = [];
    this.resources?.forEach((resource) => {
      const exportObject: any = {};
      this.selectedColumns.forEach((column) => {
        exportObject[column.field] = this.extractData(resource, column.field);
      });
      exportObjects.push(exportObject);
    });
    return exportObjects;
  }

  setSelectedResource(resource: any) {
    this.resourceSelected.emit(resource);
    this.changeDetector.detectChanges();
  }

  selectedRowsChange(event: any) {
    this.rows = event?.value;
    this.reset();
    this.changeDetector.detectChanges();
  }

  loadLazy(event: any) {
    if (event) {
      this.lastLazyEvent = event;
      this.resources = [];
    }
    this.resourceQuery = this.resourcesV2QueryService.getQuery();

    this.resourceQuery.offset = this.lastLazyEvent?.first;
    this.resourceQuery.limit = this.lastLazyEvent?.rows;

    this.fetchResourceSet();
    this.changeDetector.detectChanges();
  }

  loadMore(): void {
    this.resourceQuery.offset = this.resources?.length;
    this.fetchResourceSet();
    this.changeDetector.detectChanges();
  }

  reset(): void {
    this.resourceQuery = this.resourcesV2QueryService.getQuery();
    this.resourceQuery.limit = this.rows;
    this.resourceQuery.offset = 0;
    this.resources = [];
    this.fetchResourceSet();
    this.changeDetector.detectChanges();
  }

  private fetchResourceSet() {
    this.loading = true;
    this.resourceSet$ = undefined;
    this.resourceSetFetchError = undefined;

    let showOwner = true;
    if (this.clusterId) {
      showOwner = false;
    }
    this.columnDefinitions = this.columnFactoryService.getColumnDefinitions(
      this.resourceQuery?.versionkind?.Version,
      this.resourceQuery?.versionkind?.Kind,
      showOwner,
    );
    this.selectedColumns = this.columnDefinitions.filter((column) => column.enabled);
    this.resourceQuery.fields = this.getQueryFields(this.columnDefinitions);
    if (!this.resourceQuery.limit) {
      this.resourceQuery.limit = this.rows;
    }

    const filters = this.filterService.getFilters(this.lastLazyEvent, this.columnDefinitions);
    const order = this.filterService.getOrder(this.lastLazyEvent, this.columnDefinitions);
    this.resourceQuery.filters = filters;
    this.resourceQuery.order = order;

    if (this.clusterId) {
      this.resourceQuery.filters.push({
        field: 'rormeta.ownerref.subject',
        value: this.clusterId,
        operator: 'eq',
        type: 'string',
      });
    }

    this.resourceSet$ = this.resourcesv2Service.getResources(this.resourceQuery).pipe(
      share(),
      map((resourceSet: ResourceSet) => {
        if (!resourceSet || !resourceSet.resources) {
          this.showLoadMore = false;
        } else if (resourceSet.resources.length < this.resourceQuery.limit) {
          this.resources = [...this.resources, ...resourceSet?.resources];
          this.showLoadMore = false;
        } else {
          this.resources = [...this.resources, ...resourceSet?.resources];
        }

        this.changeDetector.detectChanges();
        return resourceSet;
      }),
      catchError((error: any) => {
        this.resourceSetFetchError = error;
        this.changeDetector.detectChanges();
        this.loading = false;
        throw error;
      }),
      finalize(() => {
        this.loading = false;
        this.changeDetector.detectChanges();
      }),
    );
  }

  private getQueryFields(fields: any[]): string[] {
    if (!fields) {
      return [];
    }

    let f = [];
    fields.filter((field: any) => {
      if (
        field &&
        field.field.indexOf('metadata') < 0 &&
        field.field.indexOf('kind') < 0 &&
        field.field.indexOf('apiversion') < 0 &&
        field.field.indexOf('typemeta') < 0 &&
        field.field.indexOf('rormeta') < 0
      ) {
        f.push(field.field.toLowerCase());
      }
    });
    if (f.length === 0) {
      f.push('metadata');
    }
    return f;
  }

  private fetchAcl(): void {
    this.adminCreate$ = this.aclService.check(AclScopes.ROR, AclScopes.Global, AclAccess.Create).pipe(
      share(),
      catchError((error: any) => {
        this.aclFetchError = error;
        this.changeDetector.detectChanges();
        throw error;
      }),
    );
  }
}

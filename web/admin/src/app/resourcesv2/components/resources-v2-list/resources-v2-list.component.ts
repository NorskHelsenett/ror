import { ResourceQuery } from './../../../core/models/resource-query';
import { catchError, delay, finalize, map, Observable, share } from 'rxjs';
import { Component, OnInit, ChangeDetectorRef, inject, ChangeDetectionStrategy, Output, EventEmitter } from '@angular/core';
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
import { Resourcesv2FilterComponent } from '../resourcesv2-filter/resourcesv2-filter.component';
import { MultiSelectModule } from 'primeng/multiselect';
import { FormsModule } from '@angular/forms';
import { RouterLink } from '@angular/router';
import { AclService } from '../../../core/services/acl.service';
import { AclAccess, AclScopes } from '../../../core/models/acl-scopes';
import { DropdownModule } from 'primeng/dropdown';

@Component({
  selector: 'app-resources-v2-list',
  standalone: true,
  imports: [
    CommonModule,
    TranslateModule,
    SharedModule,
    TableModule,
    Resourcesv2FilterComponent,
    MultiSelectModule,
    FormsModule,
    RouterLink,
    DropdownModule,
  ],
  templateUrl: './resources-v2-list.component.html',
  styleUrl: './resources-v2-list.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ResourcesV2ListComponent implements OnInit {
  @Output() resourceSelected = new EventEmitter<any>();

  resourceQuery = new ResourceQuery({
    limit: 10,
  });

  loading = false;
  showLoadMore = true;

  resources: Resource[] = [];
  resourceSet$: Observable<ResourceSet> | undefined = undefined;
  resourceSetFetchError: any;
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

  selectedColumns: any[] = [];
  rows = 10;
  rowsPerPage = [5, 10, 25, 50, 100];
  columnDefinitions: any[] = [];

  private changeDetector = inject(ChangeDetectorRef);
  private resourcesv2Service = inject(Resourcesv2Service);
  private columnFactoryService = inject(ColumnFactoryService);
  private configService = inject(ConfigService);
  private resourcesV2QueryService = inject(ResourcesV2QueryService);
  private aclService = inject(AclService);

  adminCreate$: Observable<boolean> | undefined;
  aclFetchError: any;

  ngOnInit() {
    this.rows = this.configService.config.rows;
    this.rowsPerPage = this.configService.config.rowsPerPage;
    this.fetchAcl();
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

  export(): any {
    const exportObjects: any[] = [];
    return exportObjects;
  }

  setSelectedResource(resource: any) {
    this.resourceSelected.emit(resource);
    this.changeDetector.detectChanges();
  }

  selectedRowsChange(event: any) {
    console.log('rows event: ', event);
    this.rows = event?.value;
    this.resources = [];
    this.fetchResourceSet();
    this.changeDetector.detectChanges();
  }

  loadLazy(event: any) {
    this.resourceQuery.offset = event.first;
    this.resourceQuery.limit = event.rows;
    this.fetchResourceSet();
    this.changeDetector.detectChanges();
  }

  loadMore(): void {
    this.resourceQuery.offset = this.resources?.length;
    this.fetchResourceSet();
    this.changeDetector.detectChanges();
  }

  reset(): void {
    this.resourceQuery = new ResourceQuery({
      limit: this.rows,
    });
    this.resourceQuery.offset = 0;
    this.resources = [];
    this.fetchResourceSet();
    this.changeDetector.detectChanges();
  }

  private fetchResourceSet() {
    this.loading = true;
    this.resourceSet$ = undefined;
    this.resourceSetFetchError = undefined;

    this.columnDefinitions = this.columnFactoryService.getColumnDefinitions(
      this.resourceQuery?.versionkind?.Version,
      this.resourceQuery?.versionkind?.Kind,
    );
    this.selectedColumns = this.columnDefinitions.filter((column) => column.enabled);
    if (this.resourceQuery === undefined) {
      this.resourceQuery = new ResourceQuery({
        limit: this.rows,
      });
    }
    this.resourceQuery.fields = this.getQueryFields(this.columnDefinitions);
    if (!this.resourceQuery.limit) {
      this.resourceQuery.limit = this.rows;
    }
    this.resourceSet$ = this.resourcesv2Service.getResources(this.resourceQuery).pipe(
      share(),
      map((resourceSet: ResourceSet) => {
        if (!resourceSet) {
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

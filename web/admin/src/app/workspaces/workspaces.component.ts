import { WorkspacesService } from '../core/services/workspaces.service';
import { ChangeDetectionStrategy, ChangeDetectorRef, Component, OnInit } from '@angular/core';
import { catchError, finalize, map, Observable, share, tap } from 'rxjs';
import { MetricsService } from '../core/services/metrics.service';
import { Filter } from '../core/models/apiFilter';
import { PaginationResult } from '../core/models/paginatedResult';
import { FilterService } from '../core/services/filter.service';
import { ConfigService } from '../core/services/config.service';

@Component({
  selector: 'app-workspaces',
  templateUrl: './workspaces.component.html',
  styleUrls: ['./workspaces.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class WorkspacesComponent implements OnInit {
  workspaces$: Observable<any> | undefined;
  workspaces: any[];
  workspacesError: any;

  workspaceMetrics$: Observable<any> | undefined;
  workspaceMetricsError: any;

  filter: Filter = {
    skip: 0,
    limit: 25,
    sort: [
      {
        sortField: '_id',
        sortOrder: 1,
      },
    ],
    filters: [],
  };
  lastFilter: Filter;
  lastTableEvent: any;

  currentPage = 1;
  itemsPerPage = 25;
  totalItems: number;
  pageCount: number;
  pages: number[] = [];
  loading: boolean;

  rowsPerPage = this.configService.config.rowsPerPage;
  rows = this.configService.config.rows;

  constructor(
    private changeDetector: ChangeDetectorRef,
    private workspaceService: WorkspacesService,
    private metricsService: MetricsService,
    private filterService: FilterService,
    private configService: ConfigService,
  ) {}

  ngOnInit(): void {
    this.fetchWorkspaces();
  }

  fetchWorkspaces(): void {
    this.workspacesError = undefined;
    this.workspaces$ = this.workspaceService.get().pipe(
      share(),
      tap((workspaces: any) => {
        this.workspaces = workspaces;
      }),
      catchError((error) => {
        this.workspacesError = error;
        this.changeDetector.detectChanges();
        return error;
      }),
    );
  }

  fetchMetrics(event: any): void {
    if (event) {
      this.filter = this.filterService.mapFilter(event);
      this.lastTableEvent = event;
    }

    if (!this.workspaces || this.workspaces?.length === 0) {
      this.fetchWorkspaces();
    }

    if (this.filter?.sort?.length === 0) {
      this.filter = {
        ...this.filter,
        sort: [
          {
            sortField: '_id',
            sortOrder: 1,
          },
        ],
      };
    }

    this.loading = true;
    this.lastFilter = this.filter;

    this.workspaceMetricsError = undefined;
    this.workspaceMetrics$ = this.metricsService.getForWorkspaces(this.filter).pipe(
      share(),
      map((data: PaginationResult<any>) => {
        if (!data || !data?.data) {
          return data;
        }

        let metricsData = [];
        data?.data?.forEach((workspaceMetric: any) => {
          workspaceMetric.workspace = this.workspaces.find((x: any) => x?.name === workspaceMetric?.id);
          metricsData.push(workspaceMetric);
        });
        return { ...data, data: metricsData };
      }),
      catchError((error) => {
        this.workspacesError = error;
        return error;
      }),
      finalize(() => {
        this.loading = false;
        this.changeDetector.detectChanges();
      }),
    );
  }

  extractDatacenter(workspaceName: any): any {
    this.workspaces.forEach((workspace) => {
      if (workspace?.name === workspaceName) {
        return workspace?.datacenter;
      }
    });
  }
}

import { MetricsService } from '../../../core/services/metrics.service';
import { Component, OnInit, ChangeDetectionStrategy, ChangeDetectorRef } from '@angular/core';
import { Observable, Subscription, catchError, tap } from 'rxjs';
import { ActivatedRoute, Router } from '@angular/router';
import { Filter } from '../../../core/models/apiFilter';
import { PaginationResult } from '../../../core/models/paginatedResult';
import { ConfigService } from '../../../core/services/config.service';

@Component({
  selector: 'app-datacenter-detail',
  templateUrl: './datacenter-detail.component.html',
  styleUrls: ['./datacenter-detail.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class DatacenterDetailComponent implements OnInit {
  datacenterName: string = undefined;

  datacenterMetrics$: Observable<any> | undefined;
  datacenterMetricsError: any;

  workspaceMetrics$: Observable<any> | undefined;
  workspaceMetricsError: any;
  workspaceCount = 0;

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

  currentPage = 1;
  itemsPerPage = 25;
  totalItems: number;
  pageCount: number;
  pages: number[] = [];

  pageSizes = this.configService.config.rowsPerPage;

  private subscriptions = new Subscription();

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private changeDetector: ChangeDetectorRef,
    private metricsService: MetricsService,
    private configService: ConfigService,
  ) {}

  ngOnInit(): void {
    this.subscriptions.add(
      this.route.params
        .pipe(
          tap((data: any) => {
            this.datacenterName = data?.datacenterName;
            this.fetch();

            this.changeDetector.detectChanges();
          }),
        )
        .subscribe(),
    );

    this.datacenterName = this.route.snapshot.params['datacenterName'];
  }

  fetch(): void {
    this.fetchDatacenterMetrics();
    this.fetchWorkspaceMetrics();
  }

  fetchDatacenterMetrics(): void {
    this.datacenterMetricsError = undefined;
    this.datacenterMetrics$ = this.metricsService.getForDatacenter(this.datacenterName).pipe(
      catchError((error) => {
        switch (error?.status) {
          case 401: {
            this.router.navigateByUrl('/error/401');
            break;
          }
          case 404: {
            this.router.navigateByUrl('/error/404');
            break;
          }
        }
        this.datacenterMetricsError = error;
        this.changeDetector.detectChanges();
        return error;
      }),
    );
  }

  fetchWorkspaceMetrics(): void {
    this.workspaceMetricsError = undefined;
    this.filter = { ...this.filter, limit: this.itemsPerPage, skip: this.itemsPerPage * (this.currentPage - 1) };
    this.workspaceMetrics$ = this.metricsService.getForWorkspacesByDatacenter(this.datacenterName, this.filter).pipe(
      tap((data: PaginationResult<any>) => {
        if (data.totalCount > 0) {
          this.pageCount = Math.ceil(data.totalCount / this.itemsPerPage);
          if (this.pageCount === 0) {
            this.pageCount = 1;
          }
          this.pages = Array.from(Array(this.pageCount), (_, x) => x + 1);
        } else {
          this.pages = [0];
          this.pageCount = 0;
        }
        this.workspaceCount = data?.totalCount;
        this.changeDetector.detectChanges();
      }),
      catchError((error) => {
        switch (error?.status) {
          case 401: {
            this.router.navigateByUrl('/error/401');
            break;
          }
          case 404: {
            this.router.navigateByUrl('/error/404');
            break;
          }
        }
        this.workspaceMetricsError = error;
        this.changeDetector.detectChanges();
        return error;
      }),
    );
  }

  itemPerPageChanged(itemPerPage: number): void {
    this.itemsPerPage = itemPerPage;
    this.currentPage = 1;
    this.fetchWorkspaceMetrics();
    this.changeDetector.detectChanges();
  }

  setPage(pageNumber: number): void {
    if (pageNumber < 1) {
      pageNumber = 1;
    }

    if (pageNumber > this.pages.length) {
      pageNumber = this.pages.length;
    }

    this.currentPage = pageNumber;
    this.fetchWorkspaceMetrics();
    this.changeDetector.detectChanges();
  }

  setPageTo(param: any): void {
    return;
  }

  changeSort(field: string): void {
    if (!field || field.length === 0) {
      return;
    }
    const oldSort = this.filter.sort[0];

    let asc = 1;
    if (oldSort.sortField === field) {
      asc = oldSort.sortOrder === 1 ? -1 : 1;
    }

    this.filter.sort = [
      {
        sortField: field,
        sortOrder: asc,
      },
    ];

    this.fetchWorkspaceMetrics();
  }
}

import { ClustersService } from '../../../core/services/clusters.service';
import { ChangeDetectorRef, Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { catchError, Observable, Subscription, tap } from 'rxjs';
import { User } from '../../../core/models/user';
import { Filter } from '../../../core/models/apiFilter';
import { MetricsService } from '../../../core/services/metrics.service';
import { UserService } from '../../../core/services/user.service';
import { ConfigService } from '../../../core/services/config.service';
import { PaginationResult } from '../../../core/models/paginatedResult';

@Component({
  selector: 'app-workspace-details',
  templateUrl: './workspace-details.component.html',
  styleUrls: ['./workspace-details.component.scss'],
})
export class WorkspaceDetailsComponent implements OnInit, OnDestroy {
  user$: Observable<User>;
  workspaceName: string;

  workspaceMetrics$: Observable<any> | undefined;
  workspaceMetricsError: any;

  clusters$: Observable<any> | undefined;
  clustersError: any;

  filter: Filter = {
    skip: 0,
    limit: 25,
    sort: [
      {
        sortField: 'clustername',
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
    private clusterService: ClustersService,
    private userService: UserService,
    private configService: ConfigService,
  ) {}

  ngOnInit(): void {
    this.user$ = this.userService.user;
    this.subscriptions.add(
      this.route.params
        .pipe(
          tap((data: any) => {
            this.workspaceName = data?.workspaceName;
            this.fetch();

            this.changeDetector.detectChanges();
          }),
        )
        .subscribe(),
    );

    this.workspaceName = this.route.snapshot.params['workspaceName'];
  }

  ngOnDestroy(): void {
    this.subscriptions.unsubscribe();
  }

  fetch(): void {
    this.fetchWorkspaceMetrics();
    this.fetchClusters();
  }

  fetchWorkspaceMetrics(): void {
    this.workspaceMetricsError = undefined;
    this.workspaceMetrics$ = this.metricsService.getForWorkspace(this.workspaceName).pipe(
      tap(() => {
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

  fetchClusters(): void {
    this.clustersError = undefined;
    this.filter = { ...this.filter, limit: this.itemsPerPage, skip: this.itemsPerPage * (this.currentPage - 1) };
    this.clusters$ = this.clusterService.getByWorkspace(this.workspaceName, this.filter).pipe(
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
        this.clustersError = error;
        this.changeDetector.detectChanges();
        return error;
      }),
    );
  }

  itemPerPageChanged(itemPerPage: number): void {
    this.itemsPerPage = itemPerPage;
    this.currentPage = 1;
    this.fetchClusters();
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
    this.fetchClusters();
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

    this.fetchClusters();
  }
}

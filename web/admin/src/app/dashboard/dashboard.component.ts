import { ConfirmationService, MessageService } from 'primeng/api';
import { LangChangeEvent, TranslateService } from '@ngx-translate/core';
import { MetricsTotal } from '../core/models/metricsTotal';
import { ThemeService } from '../core/services/theme.service';
import { ChangeDetectionStrategy, ChangeDetectorRef, Component, OnDestroy, OnInit } from '@angular/core';
import { catchError, forkJoin, map, Observable, of, share, Subscription, tap } from 'rxjs';
import { Filter } from '../core/models/apiFilter';
import { Datacenter } from '../core/models/datacenter';
import { ClustersService } from '../core/services/clusters.service';

import { PaginationResult } from '../core/models/paginatedResult';
import { DatacenterService } from '../core/services/datacenter.service';
import { MetricsService } from '../core/services/metrics.service';
import { OrderService } from '../core/services/order.service';
import { SignalService } from '../create/create-cluster/services/signal.service';

import dayjs from 'dayjs';
import dayjsDuration from 'dayjs/plugin/duration';
import dayjsRelativeTime from 'dayjs/plugin/relativeTime';
import { AclService } from '../core/services/acl.service';
import { AclAccess, AclScopes } from '../core/models/acl-scopes';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class DashboardComponent implements OnInit, OnDestroy {
  isDark: boolean;

  metricsCollection$: Observable<any> | undefined;
  metricsForUser$: Observable<any> | undefined;
  metricsForUser: MetricsTotal | undefined;
  metricsTotal$: Observable<any> | undefined;
  metricsTotal: MetricsTotal | undefined;
  metricsError: any;

  workspaces$: Observable<any> | undefined;
  workspaceData: PaginationResult<any> | undefined;
  workspacesError: any;
  workspacesFilter: Filter = {
    skip: 0,
    limit: 10,
    sort: [
      {
        sortField: '_id',
        sortOrder: 1,
      },
    ],
    filters: [],
  };

  datacenters$: Observable<Datacenter[]> | undefined;
  datacenterData: any;
  datacentersError: any;

  clusters$: Observable<any> | undefined;
  clustersError: any;
  clusterFilter: Filter = {
    skip: 0,
    limit: 25,
  };

  orders$: Observable<any> | undefined;
  ordersError: any;

  totalCpuPercentage = '';
  totalMemoryPercentage = '';

  workspacesCurrentPage = 1;
  workspacesItemsPerPage = 25;
  workspacesPageCount: number;
  workspacesPages: number[] = [];

  clustersCurrentPage = 1;
  clustersItemsPerPage = 25;
  clustersPageCount: number;
  clustersPages: number[] = [];
  clusterCount: number = 0;
  clusterData: PaginationResult<any> | undefined;

  clusterCreated$: Observable<string> | undefined;
  clusterOrderUpdated$: Observable<string> | undefined;

  adminDelete$: Observable<boolean> | undefined;
  adminRead$: Observable<boolean> | undefined;
  aclFetchError: any;

  private subscriptions = new Subscription();

  constructor(
    private changeDetector: ChangeDetectorRef,
    private metricsService: MetricsService,
    private clusterService: ClustersService,
    private datacenterService: DatacenterService,
    private themeService: ThemeService,
    private orderService: OrderService,
    private signalService: SignalService,
    private translateService: TranslateService,
    private messageService: MessageService,
    private aclService: AclService,
    private confirmationService: ConfirmationService,
  ) {
    dayjs.extend(dayjsRelativeTime);
    dayjs.extend(dayjsDuration);
    dayjs.locale('en');
  }

  ngOnInit(): void {
    this.fetch();
    this.subscriptions.add(
      this.themeService.isDark.subscribe((value) => {
        this.isDark = value;
        this.changeDetector.detectChanges();
      }),
    );

    this.subscriptions.add(
      this.translateService.onLangChange.subscribe((lang: LangChangeEvent) => {
        dayjs.locale(lang?.lang);
        this.changeDetector.detectChanges();
      }),
    );

    this.setupEvents();

    this.changeDetector.detectChanges();
  }

  ngOnDestroy(): void {
    this.subscriptions.unsubscribe();
  }

  fetch(): void {
    this.fetchDatacenters();
    this.fetchOverallMetrics();
    this.fetchClusters();
    this.fetchWorkspaces();
    this.fetchOrders();
    this.fetchAcl();
    this.changeDetector.detectChanges();
  }

  fetchAcl(): void {
    this.adminDelete$ = this.aclService.check(AclScopes.ROR, AclScopes.Global, AclAccess.Delete).pipe(
      share(),
      catchError((error: any) => {
        this.aclFetchError = error;
        this.changeDetector.detectChanges();
        throw error;
      }),
    );
    this.adminRead$ = this.aclService.check(AclScopes.ROR, AclScopes.Global, AclAccess.Read).pipe(
      share(),
      catchError((error: any) => {
        this.aclFetchError = error;
        this.changeDetector.detectChanges();
        throw error;
      }),
    );
  }

  fetchOverallMetrics(): void {
    this.metricsError = undefined;

    this.metricsForUser$ = this.metricsService.getTotalForUser();
    this.metricsTotal$ = this.metricsService.getTotal();

    this.metricsCollection$ = forkJoin([this.metricsForUser$, this.metricsTotal$]).pipe(
      tap(([userMetrics, totalMetrics]) => {
        this.metricsForUser = userMetrics;
        this.metricsTotal = totalMetrics;

        const cpuValue = totalMetrics?.cpuConsumed / 10 / totalMetrics?.cpu;
        this.totalCpuPercentage = cpuValue.toFixed(0);
        if (this.totalCpuPercentage == 'NaN') {
          this.totalCpuPercentage = '';
        }
        const memoryValue = (totalMetrics?.memoryConsumed * 100) / totalMetrics?.memory;
        this.totalMemoryPercentage = memoryValue.toFixed(0);
        if (this.totalMemoryPercentage == 'NaN') {
          this.totalMemoryPercentage = '';
        }
        this.changeDetector.detectChanges();
      }),
      catchError((error) => {
        this.metricsError = error;
        this.changeDetector.detectChanges();
        return error;
      }),
    );
  }

  fetchClusters(): void {
    this.clustersError = undefined;
    this.clusterFilter = {
      ...this.clusterFilter,
      limit: this.clustersItemsPerPage,
      skip: this.clustersItemsPerPage * (this.clustersCurrentPage - 1),
    };
    this.clusters$ = this.clusterService.getByFilter(this.clusterFilter).pipe(
      tap((data: PaginationResult<any>) => {
        if (data) {
          if (data.totalCount > 0) {
            this.clustersPageCount = Math.ceil(data.totalCount / this.clustersItemsPerPage);
            if (this.clustersPageCount === 0) {
              this.clustersPageCount = 1;
            }
            this.clustersPages = Array.from(Array(this.clustersPageCount), (_, x) => x + 1);
          } else {
            this.clustersPages = [0];
            this.clustersPageCount = 0;
          }
          this.clusterCount = data?.totalCount;
          this.clusterData = data;
        }
        this.changeDetector.detectChanges();
      }),
      catchError((error) => {
        this.clustersError = error;
        this.changeDetector.detectChanges();
        return error;
      }),
    );
  }

  fetchWorkspaces(): void {
    this.workspacesError = undefined;
    this.workspacesFilter = {
      ...this.workspacesFilter,
      limit: this.workspacesItemsPerPage,
      skip: this.workspacesItemsPerPage * (this.workspacesCurrentPage - 1),
    };
    this.workspaces$ = this.metricsService.getForWorkspaces(this.workspacesFilter).pipe(
      share(),
      tap((data: PaginationResult<any>) => {
        if (data) {
          if (data?.totalCount > 0) {
            this.workspacesPageCount = Math.ceil(data.totalCount / this.workspacesItemsPerPage);
            if (this.workspacesPageCount === 0) {
              this.workspacesPageCount = 1;
            }
            this.workspacesPages = Array.from(Array(this.workspacesPageCount), (_, x) => x + 1);
          } else {
            this.workspacesPages = [0];
            this.workspacesPageCount = 0;
          }
          this.workspaceData = data;
        }
        this.changeDetector.detectChanges();
      }),
      catchError((error) => {
        this.workspacesError = error;
        this.changeDetector.detectChanges();
        return error;
      }),
    );
  }

  fetchDatacenters(): void {
    this.datacentersError = undefined;
    this.datacenters$ = this.datacenterService.get().pipe(
      tap((data: any) => {
        this.datacenterData = data;
      }),
      catchError((error) => {
        this.datacentersError = error;
        this.changeDetector.detectChanges();
        return error;
      }),
    );
  }

  fetchOrders(): void {
    this.ordersError = undefined;
    this.orders$ = this.orderService.getOrders().pipe(
      map((data: any) => {
        if (!data) {
          return [];
        }
        return data?.clusterorders?.filter(
          (order: any) => order?.status?.phase === 'Recieved' || order?.status?.phase === 'Creating' || order?.status?.phase === 'Error',
        );
      }),
      catchError((error) => {
        this.ordersError = error;
        this.changeDetector.detectChanges();
        return of(false);
      }),
    );
  }

  timeSince(date: string): string {
    return dayjs(date).locale(this.translateService.currentLang).fromNow();
  }

  deleteOrder(orderUid: any): void {
    this.confirmationService.confirm({
      header: this.translateService.instant('pages.dashboard.orders.delete.title'),
      message: this.translateService.instant('pages.dashboard.orders.delete.details'),
      accept: () => {
        this.subscriptions.add(
          this.orderService.deleteOrder(orderUid).subscribe((orderDeleted: boolean) => {
            if (orderDeleted) {
              this.messageService.add({
                severity: 'success',
                summary: this.translateService.instant('pages.dashboard.orders.delete.success'),
              });
            } else {
              this.messageService.add({
                severity: 'error',
                summary: this.translateService.instant('pages.dashboard.orders.delete.error'),
              });
            }
            this.fetchOrders();
            this.changeDetector.detectChanges();
          }),
        );
      },
    });
  }

  private setupEvents(): void {
    this.clusterOrderUpdated$ = this.signalService?.clusterOrderUpdated$.pipe(
      tap((event: any) => {
        if (event) {
          this.fetchOrders();
          this.changeDetector.detectChanges();
        }
      }),
    );

    this.clusterCreated$ = this.signalService?.clusterCreated$.pipe(
      tap((event: any) => {
        if (event) {
          this.fetchOrders();
          this.changeDetector.detectChanges();
        }
      }),
    );
  }
}

import { TranslateService } from '@ngx-translate/core';
import { ChangeDetectionStrategy, ChangeDetectorRef, Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Observable, Subscription, catchError, tap } from 'rxjs';
import dayjs from 'dayjs';
import dayjsRelativeTime from 'dayjs/plugin/relativeTime';
import { environment } from '../../../../environments/environment';
import { ClusterOrder, TanzuConfig } from '../../../core/models/clusterOrder';
import { OrderService } from '../../../core/services/order.service';
import { WorkspacesService } from '../../../core/services/workspaces.service';
import { DatacenterService } from '../../../core/services/datacenter.service';
import { Datacenter } from '../../../core/models/datacenter';
import { Workspace } from '../../../core/models/workspace';

@Component({
  selector: 'app-order-detail',
  templateUrl: './order-detail.component.html',
  styleUrls: ['./order-detail.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class OrderDetailComponent implements OnInit {
  order$: Observable<ClusterOrder> | undefined;
  orderFetchError: any;

  workspace$: Observable<Workspace> | undefined;
  workspaceFetchError: any;

  datacenter$: Observable<Datacenter> | undefined;
  datacenterFetchError: any;

  uid: string = undefined;
  environment = environment;
  showRaw = false;

  private subscriptions = new Subscription();
  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private changeDetector: ChangeDetectorRef,
    private orderService: OrderService,
    private workspaceService: WorkspacesService,
    private datacenterService: DatacenterService,
    private translateService: TranslateService,
  ) {
    dayjs.locale(this.translateService.currentLang);
    dayjs.extend(dayjsRelativeTime);
  }

  ngOnInit(): void {
    this.subscriptions.add(
      this.route.params
        .pipe(
          tap((data: any) => {
            this.uid = data?.uid;
            this.fetchOrder();

            this.changeDetector.detectChanges();
          }),
        )
        .subscribe(),
    );

    this.uid = this.route.snapshot.params['uid'];
  }

  fetchOrder(): void {
    this.orderFetchError = undefined;
    this.order$ = this.orderService.getOrder(this.uid).pipe(
      tap((order) => {
        if (order?.spec.provider?.toLowerCase() === 'tanzu') {
          let providerConfig = order?.spec?.providerConfig as TanzuConfig;
          this.fetchWorkspace(providerConfig?.namespaceId);
          this.fetchDatacenter(providerConfig?.datacenterId);
        }

        this.changeDetector.detectChanges();
      }),
      catchError((error) => {
        this.orderFetchError = error;
        this.changeDetector.detectChanges();
        throw error;
      }),
    );
  }

  deleteOrder(): void {
    this.orderService.deleteOrder(this.uid).subscribe((result: boolean) => {
      if (result) {
        this.router.navigate(['/']);
      }
    });
  }

  getTime(date: string): Date {
    let result = dayjs(date).locale(this.translateService.currentLang);
    return result.toDate();
  }

  timeSince(date: string): string {
    return dayjs(date).locale(this.translateService.currentLang).fromNow();
  }

  private fetchWorkspace(workspaceId: string): void {
    this.workspace$ = this.workspaceService.getById(workspaceId).pipe(
      tap((workspace) => {
        this.changeDetector.detectChanges();
      }),
      catchError((error) => {
        this.workspaceFetchError = error;
        this.changeDetector.detectChanges();
        throw error;
      }),
    );
  }

  private fetchDatacenter(datacenterId: string): void {
    this.datacenter$ = this.datacenterService.getById(datacenterId).pipe(
      tap((datacenter) => {
        this.changeDetector.detectChanges();
      }),
      catchError((error) => {
        this.datacenterFetchError = error;
        this.changeDetector.detectChanges();
        throw error;
      }),
    );
  }
}

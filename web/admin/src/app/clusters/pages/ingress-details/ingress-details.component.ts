import { OAuthService } from 'angular-oauth2-oidc';
import { Component, OnInit, OnDestroy, ChangeDetectorRef, ChangeDetectionStrategy } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Observable, Subscription, tap, catchError } from 'rxjs';
import { ClustersService } from '../../../core/services/clusters.service';

@Component({
  selector: 'app-ingress-details',
  templateUrl: './ingress-details.component.html',
  styleUrls: ['./ingress-details.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class IngressDetailsComponent implements OnInit, OnDestroy {
  clusterId: string | undefined;
  ingressId: string | undefined;

  cluster$: Observable<any> | undefined;
  clusterError: any;

  ingresses: any[] = [];
  ingress: any;

  metrics$: Observable<any> | undefined;
  metricsError: any;
  userClaims: any;

  private subscriptions = new Subscription();

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private changeDetector: ChangeDetectorRef,
    private clusterService: ClustersService,
    private oauthService: OAuthService,
  ) {}

  ngOnInit(): void {
    this.subscriptions.add(
      this.route.params
        .pipe(
          tap((data: any) => {
            this.clusterId = data?.id;
            this.ingressId = data?.ingressid;
            if (this.clusterId?.length < 1) {
              this.router.navigateByUrl('/error/401');
              return;
            }
            this.fetchCluster(this.clusterId);
            this.changeDetector.detectChanges();
          }),
        )
        .subscribe(),
    );
    this.userClaims = this.oauthService.getIdentityClaims();
    this.clusterId = this.route.snapshot.params['id'];
  }

  ngOnDestroy(): void {
    this.subscriptions.unsubscribe();
  }
  fetch(): void {
    this.fetchCluster(this.clusterId);
  }
  fetchCluster(clusterId: string): void {
    this.clusterError = undefined;
    this.cluster$ = this.clusterService.getById(clusterId).pipe(
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
        this.clusterError = error;
        return error;
      }),
    );
  }
}

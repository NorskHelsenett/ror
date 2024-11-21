import { ResourceQueryFilter } from './../../../core/models/resources-v2';
import { ResourceFilter } from './../../../resources/models/resourceFilter';
import { Resourcesv2Service } from './../../../core/services/resourcesv2.service';
import { CommonModule } from '@angular/common';
import { TranslateModule } from '@ngx-translate/core';
import { Component, OnInit, ChangeDetectionStrategy, ChangeDetectorRef, inject, OnDestroy } from '@angular/core';
import { catchError, finalize, map, Observable, Subscription, tap } from 'rxjs';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { ClustersService } from '../../../core/services/clusters.service';
import { SharedModule } from '../../../shared/shared.module';
import { HighlightModule } from 'ngx-highlightjs';
import { TabViewModule } from 'primeng/tabview';
import { ClusterIngressAnnotationsComponent } from '../../components/cluster-ingress-annotations/cluster-ingress-annotations.component';
import { ClusterIngressService } from '../../services/cluster-ingress.service';
import { Resource, ResourceSet, ResourceQuery } from '@rork8s/ror-resources/models';
import { ClusterIngressMetaComponent } from '../../components/cluster-ingress-meta/cluster-ingress-meta.component';
import { ClusterIngressDetailsComponent } from '../../components/cluster-ingress-details/cluster-ingress-details.component';
import { ClusterIngressChartComponent } from '../../components/cluster-ingress-chart/cluster-ingress-chart.component';

@Component({
  selector: 'app-ingress',
  standalone: true,
  imports: [
    CommonModule,
    TranslateModule,
    SharedModule,
    RouterLink,
    HighlightModule,
    TabViewModule,
    ClusterIngressAnnotationsComponent,
    ClusterIngressMetaComponent,
    ClusterIngressDetailsComponent,
    ClusterIngressChartComponent,
  ],
  templateUrl: './ingress.component.html',
  styleUrl: './ingress.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class IngressComponent implements OnInit, OnDestroy {
  clusterId: string | undefined;
  cluster$: Observable<any> | undefined;
  clusterFetchError: any;

  ingressId: string | undefined;
  ingress$: Observable<Resource> | undefined;
  ingressFetchError: any;

  certificates$: Observable<Resource> | undefined;
  certificatesFetchError: any;

  services: Resource[] | undefined;
  servicesFetchError: any;

  pods: Resource[] | undefined;
  podsFetchError: any;

  private subscriptions = new Subscription();

  private changeDetector = inject(ChangeDetectorRef);
  private clustersServices = inject(ClustersService);
  private resourcesv2Service = inject(Resourcesv2Service);
  private route = inject(ActivatedRoute);
  private clusterIngressService = inject(ClusterIngressService);

  constructor() {}

  ngOnInit(): void {
    this.subscriptions.add(
      this.route.params
        .pipe(
          tap((data: any) => {
            this.clusterId = data?.id;
            this.ingressId = data?.ingressid;
            this.fetchCluster();
            this.fetchIngress();
          }),
        )
        .subscribe(),
    );
  }

  ngOnDestroy(): void {
    this.subscriptions.unsubscribe();
  }

  refresh() {
    this.fetchCluster();
    this.fetchIngress();
  }

  fetchCluster() {
    this.cluster$ = undefined;
    this.clusterFetchError = undefined;
    if (this.clusterId) {
      this.cluster$ = this.clustersServices.getById(this.clusterId).pipe(
        map((data: any) => {
          this.clusterIngressService.setCluster(data);
          return data;
        }),
        catchError((error) => {
          this.clusterFetchError = error;
          this.changeDetector.detectChanges();
          throw error;
        }),
        finalize(() => {
          this.changeDetector.detectChanges();
        }),
      );
    }
  }

  fetchIngress() {
    this.ingress$ = undefined;
    this.ingressFetchError = undefined;
    if (!this.ingressId) {
      return;
    }

    const query: ResourceQuery = {
      versionkind: {
        Group: '',
        Kind: 'Ingress',
        Version: 'networking.k8s.io/v1',
      },
      filters: [
        {
          field: 'metadata.name',
          type: 'string',
          operator: 'eq',
          value: this.ingressId,
        },
        {
          field: 'rormeta.ownerref.subject',
          type: 'string',
          operator: 'eq',
          value: this.clusterId,
        },
      ],
    };

    this.ingress$ = this.resourcesv2Service.getResources(query).pipe(
      map((data: ResourceSet) => {
        if (data.resources.length === 1) {
          this.clusterIngressService.setIngress(data.resources[0]);
          let serviceNames = data.resources[0].ingress?.spec?.rules?.map((rule) => rule.http?.paths[0].backend.service.name);

          this.fetchServices(serviceNames);
          this.fetchCerficates(data.resources[0].ingress?.spec?.tls?.map((tls) => tls.secretName));
          return data.resources[0];
        } else {
          return null;
        }
      }),
      catchError((error) => {
        this.ingressFetchError = error;
        this.changeDetector.detectChanges();
        throw error;
      }),
      finalize(() => {
        this.changeDetector.detectChanges();
      }),
    );
  }

  fetchServices(serviceNames: string[]) {
    this.services = undefined;
    this.servicesFetchError = undefined;
    if (!serviceNames || serviceNames.length === 0) {
      return;
    }

    const query: ResourceQuery = {
      versionkind: {
        Group: '',
        Kind: 'Service',
        Version: 'v1',
      },
      filters: [
        {
          field: 'rormeta.ownerref.subject',
          type: 'string',
          operator: 'eq',
          value: this.clusterId,
        },
      ],
    };

    let serviceFilters: ResourceQueryFilter[] = [];
    serviceNames.forEach((serviceName) => {
      serviceFilters.push({
        field: 'metadata.name',
        type: 'string',
        operator: 'eq',
        value: serviceName,
      });
    });

    query.filters = query.filters.concat(serviceFilters);

    this.subscriptions.add(
      this.resourcesv2Service
        .getResources(query)
        .pipe(
          map((data: ResourceSet) => {
            if (data.resources.length === 1) {
              this.clusterIngressService.setServices(data.resources);

              let namespaces = data.resources.map((service) => service.metadata.namespace);
              let serviceNames = data.resources.map((service) => service.metadata.name);

              this.fetchPodsByNamespaceAndService(namespaces, serviceNames);
              return data.resources;
            } else {
              return null;
            }
          }),
          catchError((error) => {
            this.servicesFetchError = error;
            this.changeDetector.detectChanges();
            throw error;
          }),
          finalize(() => {
            this.changeDetector.detectChanges();
          }),
        )
        .subscribe(),
    );
  }

  fetchPodsByNamespaceAndService(namespaces: string[], serviceNames: string[]) {
    this.pods = undefined;
    this.podsFetchError = undefined;
    if (!serviceNames || serviceNames.length === 0 || !namespaces || namespaces.length === 0) {
      return;
    }

    const query: ResourceQuery = {
      versionkind: {
        Group: '',
        Kind: 'Pod',
        Version: 'v1',
      },
      filters: [
        {
          field: 'rormeta.ownerref.subject',
          type: 'string',
          operator: 'eq',
          value: this.clusterId,
        },
      ],
    };

    let namespaceFilters: ResourceQueryFilter[] = [];
    namespaces.forEach((namespace) => {
      if (!namespace) {
        return;
      }

      if (namespaceFilters.filter((filter) => filter.value === namespace).length === 0) {
        namespaceFilters.push({
          field: 'metadata.namespace',
          type: 'string',
          operator: 'eq',
          value: namespace,
        });
      }
    });

    query.filters = query.filters.concat(namespaceFilters);

    this.subscriptions.add(
      this.resourcesv2Service
        .getResources(query)
        .pipe(
          map((data: ResourceSet) => {
            if (data.resources) {
              let result = data.resources.filter((pod) => {
                let podServiceName = pod.metadata.labels['app.kubernetes.io/instance'];
                return serviceNames.includes(podServiceName);
              });

              this.clusterIngressService.setPods(result);
              return result;
            } else {
              return null;
            }
          }),
          catchError((error) => {
            this.servicesFetchError = error;
            this.changeDetector.detectChanges();
            throw error;
          }),
          finalize(() => {
            this.changeDetector.detectChanges();
          }),
        )
        .subscribe(),
    );
  }

  fetchCerficates(certNames: string[]) {
    this.certificates$ = undefined;
    this.certificatesFetchError = undefined;
    if (!certNames || certNames.length === 0) {
      return;
    }

    const query: ResourceQuery = {
      versionkind: {
        Group: '',
        Kind: 'Certificate',
        Version: 'cert-manager.io/v1',
      },
    };

    let certFilters: ResourceQueryFilter[] = [];
    certNames.forEach((certName) => {
      certFilters.push({
        field: 'metadata.name',
        type: 'string',
        operator: 'eq',
        value: certName,
      });
    });

    query.filters = certFilters;

    this.subscriptions.add(
      this.resourcesv2Service
        .getResources(query)
        .pipe(
          map((data: ResourceSet) => {
            if (data.resources) {
              this.clusterIngressService.setCertificates(data.resources);
              return data.resources;
            } else {
              return null;
            }
          }),
          catchError((error) => {
            this.ingressFetchError = error;
            this.changeDetector.detectChanges();
            throw error;
          }),
          finalize(() => {
            this.changeDetector.detectChanges();
          }),
        )
        .subscribe(),
    );
  }
}

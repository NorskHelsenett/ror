import { TranslateService } from '@ngx-translate/core';
import { inject, Injectable, signal } from '@angular/core';
import { Resource } from '@rork8s/ror-resources/models';
import { HealtStatus } from '../../core/models/healthstatus';

@Injectable({
  providedIn: 'root',
})
export class ClusterIngressService {
  private cluster = signal<any>(undefined);
  private ingress = signal<Resource>(undefined);
  private services = signal<Resource[]>(undefined);
  private endpoints = signal<Resource[]>(undefined);
  private pods = signal<Resource[]>(undefined);
  private certificates = signal<Resource[]>(undefined);

  private translateService = inject(TranslateService);

  setCluster(cluster: any) {
    this.cluster.set(cluster);
  }

  getCluster() {
    return this.cluster();
  }

  setIngress(resource: Resource) {
    this.ingress.set(resource);
  }

  getIngress() {
    return this.ingress();
  }

  setServices(resources: Resource[]) {
    this.services.update(() => resources);
  }

  setEndpoints(resources: Resource[]) {
    this.endpoints.update(() => resources);
  }

  getEndpoints() {
    return this.endpoints();
  }

  getServices() {
    return this.services();
  }

  setPods(resources: Resource[]) {
    this.pods.update(() => resources);
  }

  getPods() {
    return this.pods();
  }

  setCertificates(resources: Resource[]) {
    this.certificates.update(() => resources);
  }

  getCertificates() {
    return this.certificates();
  }

  getHealthStatusForIngress(): HealtStatus {
    return {
      healthy: true,
      messages: ['Healthy'],
    };
  }

  getHealthStatusForServices(): HealtStatus {
    let result: HealtStatus = {
      healthy: false,
      messages: [],
    };

    if (!this.services()) {
      result.healthy = false;
      result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.service.noservice'));
    }

    this.services()?.forEach((service) => {
      if (!service || !service?.service) {
        result.healthy = false;
        result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.service.noservice'));
      }

      if (!service?.metadata?.name) {
        result.healthy = false;
        result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.service.noname'));
      }

      if (!service?.service?.spec?.type || service?.service?.spec?.type === '' || service?.service?.spec?.type.trim() === '') {
        result.healthy = false;
        result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.service.notype'));
      }

      let supportedTypes = ['ClusterIP', 'NodePort', 'LoadBalancer', 'ExternalName'];
      let supportedTypesWithAko = ['NodePort'];
      if (this.isAkoIngress()) {
        if (!supportedTypesWithAko.includes(service?.service?.spec?.type)) {
          result.healthy = false;
          result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.service.wrongakotype'));
        }
      } else {
        if (!supportedTypes.includes(service?.service?.spec?.type)) {
          result.healthy = false;
          result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.service.wrongtype'));
        }
      }

      if (!service?.service?.spec.ipFamilyPolicy) {
        result.healthy = false;
        result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.service.noipfamiliypolicy'));
      }

      if (!service?.service?.spec.clusterIPs || service?.service?.spec.clusterIPs.length === 0) {
        result.healthy = false;
        result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.service.noclusterips'));
      }

      if (!service?.service?.spec?.ipFamilies || service?.service?.spec?.ipFamilies?.length === 0) {
        result.healthy = false;
        result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.service.noipfamilies'));
      }

      if (!service?.service?.spec.ports) {
        result.healthy = false;
        result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.services.noports'));
      }
    });

    if (result.messages.length === 0) {
      result.healthy = true;
    }

    return result;
  }

  getHealthStatusForEndpoints(): HealtStatus {
    let result: HealtStatus = {
      healthy: false,
      messages: [],
    };

    if (!this.endpoints() || this.endpoints().length === 0) {
      result.healthy = false;
      result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.endpoints.noendpoint'));
    }

    this.endpoints()?.forEach((endpoint) => {
      if (!endpoint) {
        result.healthy = false;
        result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.endpoints.noendpoint'));
      }

      if (!endpoint?.endpoints?.subsets || endpoint?.endpoints?.subsets?.length === 0) {
        result.healthy = false;
        result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.endpoints.nosubsets'));
      }

      endpoint?.endpoints?.subsets?.forEach((subset) => {
        // if (subset?.notReadyAddresses?.length > 0) {
        //   result.healthy = false;
        //   result.message = 'Not all endpoints are ready';
        // }

        if (!subset?.addresses || subset?.addresses?.length === 0) {
          result.healthy = false;
          result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.endpoints.noaddressess'));
        }
      });
    });

    if (result.messages?.length === 0) {
      result.healthy = true;
    }

    return result;
  }

  getHealthStatusForPods(): HealtStatus {
    let result: HealtStatus = {
      healthy: false,
      messages: [],
    };

    if (!this.pods() || this.pods().length === 0) {
      result.healthy = false;
      result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.pods.nopods'));
    }

    this.pods()?.forEach((pod) => {
      if (!pod) {
        result.healthy = false;
        result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.pods.nopods'));
      }

      if (!pod?.metadata?.name) {
        result.healthy = false;
        result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.pods.noname'));
      }

      if (!pod?.pod?.status?.phase || pod?.pod?.status?.phase === '' || pod?.pod?.status?.phase?.trim()?.toLocaleLowerCase() !== 'Running') {
        result.healthy = false;
        result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.pods.nophase'));
      }
    });
    return result;
  }

  isCertManagerIngress(): boolean {
    return this.ingress()?.metadata?.annotations?.['cert-manager.io/cluster-issuer'] !== undefined;
  }

  isAkoIngress(): boolean {
    return this.ingress()?.metadata?.annotations?.['ako.vmware.com/controller-cluster-uuid'] !== undefined;
  }
}

import { TranslateService } from '@ngx-translate/core';
import { inject, Injectable, signal } from '@angular/core';
import { Resource, ResourceIngressSpecRules, ResourceIngressSpecRulesHttpPaths } from '@rork8s/ror-resources/models';
import { HealthStatus } from '../../core/models/healthstatus';

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

  getHealthStatus(): HealthStatus {
    let result: HealthStatus = {
      healthy: false,
      messages: [],
    };

    let healthStatusForIngress = this.getHealthStatusForIngress();
    let healthStatusForServices = this.getHealthStatusForServices();
    let healthStatusForEndpoints = this.getHealthStatusForEndpoints();
    let healthStatusForPods = this.getHealthStatusForPods();
    let healthStatusForIngressPaths: HealthStatus = {
      healthy: false,
      messages: [],
    };

    if (this.ingress()?.ingress?.spec?.rules) {
      this.ingress()?.ingress?.spec?.rules?.forEach((rule: ResourceIngressSpecRules) => {
        rule?.http?.paths?.forEach((path) => {
          let healthStatusForIngressPath = this.getHealthStatusForIngressPath(path);
          if (!healthStatusForIngressPath.healthy) {
            healthStatusForIngressPaths.healthy = false;
            healthStatusForIngressPaths.messages.push(...healthStatusForIngressPath.messages);
          }
        });
      });
    }

    if (healthStatusForIngress.healthy) {
      result.healthy = true;
    } else {
      result.healthy = false;
      result.messages.push(...healthStatusForIngress.messages);
    }

    if (healthStatusForServices.healthy) {
      result.healthy = true;
    } else {
      result.healthy = false;
      result.messages.push(...healthStatusForServices.messages);
    }

    if (healthStatusForEndpoints.healthy) {
      result.healthy = true;
    } else {
      result.healthy = false;
      result.messages.push(...healthStatusForEndpoints.messages);
    }

    if (healthStatusForPods.healthy) {
      result.healthy = true;
    } else {
      result.healthy = false;
      result.messages.push(...healthStatusForPods.messages);
    }

    if (healthStatusForIngressPaths.healthy) {
      result.healthy = true;
    } else {
      result.healthy = false;
      result.messages.push(...healthStatusForIngressPaths.messages);
    }

    if (result.messages.length === 0) {
      result.healthy = true;
    }

    return result;
  }

  getHealthStatusForIngress(): HealthStatus {
    let result: HealthStatus = {
      healthy: false,
      messages: [],
    };

    if (!this.ingress()) {
      return {
        healthy: false,
        messages: [this.translateService.instant('pages.clusters.details.ingresses.errors.ingress.noingress')],
      };
    }

    if (!this.ingress()?.ingress?.spec?.ingressClassName || this.ingress()?.ingress?.spec?.ingressClassName === '') {
      return {
        healthy: false,
        messages: [this.translateService.instant('pages.clusters.details.ingresses.errors.ingress.noingressclassname')],
      };
    }

    if (!this.ingress()?.metadata?.name) {
      return {
        healthy: false,
        messages: [this.translateService.instant('pages.clusters.details.ingresses.errors.ingress.noingressname')],
      };
    }

    if (!this.ingress()?.metadata?.namespace) {
      return {
        healthy: false,
        messages: [this.translateService.instant('pages.clusters.details.ingresses.errors.ingress.noingressnamespace')],
      };
    }

    if (!this.ingress()?.metadata?.annotations) {
      return {
        healthy: false,
        messages: [this.translateService.instant('pages.clusters.details.ingresses.errors.ingress.noingressannotations')],
      };
    }

    if (!this.ingress()?.ingress?.spec?.rules || this.ingress()?.ingress?.spec?.rules.length === 0) {
      return {
        healthy: false,
        messages: [this.translateService.instant('pages.clusters.details.ingresses.errors.ingress.norules')],
      };
    }

    this.ingress()?.ingress?.spec?.rules?.forEach((rule) => {
      if (!rule?.http?.paths || rule?.http?.paths.length === 0) {
        result.healthy = false;
        result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.ingress.nopaths'));
      }
    });

    if (result.messages.length === 0) {
      result.healthy = true;
    }

    return result;
  }

  getHealthStatusForIngressPath(path: ResourceIngressSpecRulesHttpPaths): HealthStatus {
    let result: HealthStatus = {
      healthy: true,
      messages: [],
    };

    if (!path) {
      result.healthy = false;
      result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.ingress.nopath'));
    }

    if (!path?.path || path?.path === '') {
      result.healthy = false;
      result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.ingress.nopath'));
    }

    if (!path?.pathType || path?.pathType === '') {
      result.healthy = false;
      result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.ingress.noPathType'));
    }

    let supportedPathTypes = ['Exact', 'Prefix', 'ImplementationSpecific'];
    if (!supportedPathTypes.includes(path?.pathType)) {
      result.healthy = false;
      result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.ingress.wrongPathType'));
    }

    if (!path?.backend) {
      result.healthy = false;
      result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.ingress.nopath'));
    }

    if (!path?.backend?.service?.name) {
      result.healthy = false;
      result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.ingress.nobackendservice'));
    }

    if (
      (!path?.backend?.service?.port?.name || path?.backend?.service?.port?.name === '') &&
      (!path?.backend?.service?.port?.number || path?.backend?.service?.port?.number === 0)
    ) {
      result.healthy = false;
      result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.ingress.nobackendportnameornumber'));
    }

    if (result.messages.length === 0) {
      result.healthy = true;
    }

    return result;
  }

  getHealthStatusForServices(): HealthStatus {
    let result: HealthStatus = {
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

  getHealthStatusForEndpoints(): HealthStatus {
    let result: HealthStatus = {
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

  getHealthStatusForPods(): HealthStatus {
    let result: HealthStatus = {
      healthy: false,
      messages: [],
    };

    if (!this.pods() || this.pods().length === 0) {
      result.healthy = false;
      result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.pods.nopods'));
    }

    this.pods()?.forEach((pod) => {
      let healthStatusForPod = this.getHealthStatusForPod(pod);
      if (!healthStatusForPod.healthy) {
        result.healthy = false;
        result.messages.push(...healthStatusForPod.messages);
      }
    });

    if (result.messages.length === 0) {
      result.healthy = true;
    }
    return result;
  }

  getHealthStatusForPod(pod: Resource): HealthStatus {
    let result: HealthStatus = {
      healthy: false,
      messages: [],
    };

    if (!pod) {
      result.healthy = false;
      result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.pods.nopods'));
    }

    if (!pod?.metadata?.name) {
      result.healthy = false;
      result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.pods.noname'));
    }

    if (!pod?.pod?.status?.phase || pod?.pod?.status?.phase === '' || pod?.pod?.status?.phase?.trim()?.toLocaleLowerCase() !== 'running') {
      result.healthy = false;
      result.messages.push(this.translateService.instant('pages.clusters.details.ingresses.errors.pods.nophase'));
    }

    if (result.messages.length === 0) {
      result.healthy = true;
    }

    return result;
  }

  isCertManagerIngress(): boolean {
    return this.ingress()?.metadata?.annotations?.['cert-manager.io/cluster-issuer'] !== undefined;
  }

  isAkoIngress(): boolean {
    return this.ingress()?.metadata?.annotations?.['ako.vmware.com/controller-cluster-uuid'] !== undefined;
  }
}

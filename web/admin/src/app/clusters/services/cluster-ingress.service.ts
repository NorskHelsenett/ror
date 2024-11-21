import { Injectable, signal } from '@angular/core';
import { Resource } from '../../core/models/resources-v2';

@Injectable({
  providedIn: 'root',
})
export class ClusterIngressService {
  private cluster = signal<any>(undefined);
  private ingress = signal<Resource>(undefined);
  private services = signal<Resource[]>(undefined);
  private pods = signal<Resource[]>(undefined);
  private certificates = signal<Resource[]>(undefined);

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

  // fetchIngress() {
  //     this.ingress$ = undefined;
  //     this.ingressFetchError = undefined;
  //     if (!this.ingressId) {
  //       return;
  //     }

  //     const query: ResourceQuery = {
  //       versionkind: {
  //         Group: '',
  //         Kind: 'Ingress',
  //         Version: 'networking.k8s.io/v1',
  //       },
  //       filters: [
  //         {
  //           field: 'metadata.name',
  //           type: 'string',
  //           operator: 'eq',
  //           value: this.ingressId,
  //         },
  //         {
  //           field: 'rormeta.ownerref.subject',
  //           type: 'string',
  //           operator: 'eq',
  //           value: this.clusterId,
  //         },
  //       ],
  //     };

  //     this.ingress$ = this.resourcesv2Service.getResources(query).pipe(
  //       map((data: ResourceSet) => {
  //         if (data.resources.length === 1) {
  //           this.clusterIngressService.setResource(data.resources[0]);
  //           return data.resources[0];
  //         } else {
  //           return null;
  //         }
  //       }),
  //       catchError((error) => {
  //         this.ingressFetchError = error;
  //         this.changeDetector.detectChanges();
  //         throw error;
  //       }),
  //       finalize(() => {
  //         this.changeDetector.detectChanges();
  //       }),
  //     );
  //   }
}

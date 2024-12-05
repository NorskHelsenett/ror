import { Injectable, signal } from '@angular/core';
import { Resource } from '@rork8s/ror-resources/models';

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
}

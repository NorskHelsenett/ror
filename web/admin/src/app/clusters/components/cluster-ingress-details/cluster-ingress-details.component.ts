import { Resource } from '@rork8s/ror-resources/models';
import { ClusterIngressService } from './../../services/cluster-ingress.service';
import { Component, effect, inject, ChangeDetectorRef, ChangeDetectionStrategy } from '@angular/core';
import { HighlightModule } from 'ngx-highlightjs';
import { ClusterIngressMetaComponent } from '../cluster-ingress-meta/cluster-ingress-meta.component';

@Component({
  selector: 'app-cluster-ingress-details',
  standalone: true,
  imports: [HighlightModule, ClusterIngressMetaComponent],
  templateUrl: './cluster-ingress-details.component.html',
  styleUrl: './cluster-ingress-details.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ClusterIngressDetailsComponent {
  cluster: any | undefined;
  ingress: Resource | undefined;
  services: Resource[] | undefined;
  pods: Resource[] | undefined;
  certificates: Resource[] | undefined;

  private changeDetectorRef = inject(ChangeDetectorRef);
  private clusterIngressService = inject(ClusterIngressService);

  constructor() {
    effect(() => {
      this.cluster = this.clusterIngressService.getCluster();
      this.ingress = this.clusterIngressService.getIngress();
      this.services = this.clusterIngressService.getServices();
      this.pods = this.clusterIngressService.getPods();
      this.certificates = this.clusterIngressService.getCertificates();
      this.changeDetectorRef.detectChanges();
    });
  }
}

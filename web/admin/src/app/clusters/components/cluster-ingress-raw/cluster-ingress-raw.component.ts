import { ChangeDetectionStrategy, Component, effect, ChangeDetectorRef, inject } from '@angular/core';
import { TranslateModule } from '@ngx-translate/core';
import { HighlightModule } from 'ngx-highlightjs';

import { ClusterIngressService } from '../../services/cluster-ingress.service';
import { Resource } from '@rork8s/ror-resources/models';
import { JsonPipe } from '@angular/common';

@Component({
  selector: 'app-cluster-ingress-raw',
  standalone: true,
  imports: [TranslateModule, HighlightModule, JsonPipe],
  templateUrl: './cluster-ingress-raw.component.html',
  styleUrl: './cluster-ingress-raw.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ClusterIngressRawComponent {
  showIngress: boolean = true;
  showCertificates: boolean = false;
  showServices: boolean = false;
  showEndpoints: boolean = false;
  showPods: boolean = false;

  ingress: Resource | undefined;
  certificates: Resource[] | undefined;
  services: Resource[] | undefined;
  endpoints: Resource[] | undefined;
  pods: Resource[] | undefined;

  private changeDetector = inject(ChangeDetectorRef);
  private clusterIngressService = inject(ClusterIngressService);

  constructor() {
    effect(() => {
      this.ingress = this.clusterIngressService.getIngress();
      this.certificates = this.clusterIngressService.getCertificates();
      this.services = this.clusterIngressService.getServices();
      this.endpoints = this.clusterIngressService.getEndpoints();
      this.pods = this.clusterIngressService.getPods();
      this.changeDetector.detectChanges();
    });
  }

  show(type: string) {
    this.showIngress = type === 'ingress';
    this.showCertificates = type === 'certificates';
    this.showServices = type === 'services';
    this.showEndpoints = type === 'endpoints';
    this.showPods = type === 'pods';
    this.changeDetector.detectChanges();
  }
}

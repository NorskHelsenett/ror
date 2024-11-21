import { Resource } from '@rork8s/ror-resources/models';
import { ClusterIngressService } from './../../services/cluster-ingress.service';
import { ChangeDetectionStrategy, ChangeDetectorRef, Component, effect, inject } from '@angular/core';
import { HighlightModule } from 'ngx-highlightjs';
import { JsonPipe } from '@angular/common';

@Component({
  selector: 'app-cluster-ingress-meta',
  standalone: true,
  imports: [HighlightModule, JsonPipe],
  templateUrl: './cluster-ingress-meta.component.html',
  styleUrl: './cluster-ingress-meta.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ClusterIngressMetaComponent {
  cluster: any | undefined;
  ingress: Resource | undefined;

  private changeDetectorRef = inject(ChangeDetectorRef);
  private clusterIngressService = inject(ClusterIngressService);

  constructor() {
    effect(() => {
      this.cluster = this.clusterIngressService.getCluster();
      this.ingress = this.clusterIngressService.getIngress();
      this.changeDetectorRef.detectChanges();
    });
  }
}

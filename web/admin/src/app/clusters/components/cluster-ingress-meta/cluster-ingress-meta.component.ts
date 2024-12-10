import { Resource } from '@rork8s/ror-resources/models';
import { ClusterIngressService } from './../../services/cluster-ingress.service';
import { ChangeDetectionStrategy, ChangeDetectorRef, Component, effect, inject } from '@angular/core';
import { HighlightModule } from 'ngx-highlightjs';
import { SharedModule } from '../../../shared/shared.module';
import { HealthStatus } from '../../../core/models/healthstatus';
import { CommonModule } from '@angular/common';
import { TableModule } from 'primeng/table';
import { TranslateModule } from '@ngx-translate/core';
import { RouterModule } from '@angular/router';

@Component({
  selector: 'app-cluster-ingress-meta',
  standalone: true,
  imports: [CommonModule, HighlightModule, SharedModule, TableModule, TranslateModule, RouterModule],
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

  getIngressStatus(): HealthStatus {
    return this.clusterIngressService.getHealthStatus();
  }
}

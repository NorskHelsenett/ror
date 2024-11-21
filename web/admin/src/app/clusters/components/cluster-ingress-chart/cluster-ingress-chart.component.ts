import { ChangeDetectionStrategy, ChangeDetectorRef, Component, inject } from '@angular/core';
import { ClusterIngressService } from '../../services/cluster-ingress.service';

@Component({
  selector: 'app-cluster-ingress-chart',
  standalone: true,
  imports: [],
  templateUrl: './cluster-ingress-chart.component.html',
  styleUrl: './cluster-ingress-chart.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ClusterIngressChartComponent {
  private changeDetector = inject(ChangeDetectorRef);
  private clusterIngressService = inject(ClusterIngressService);

  constructor() {}
}

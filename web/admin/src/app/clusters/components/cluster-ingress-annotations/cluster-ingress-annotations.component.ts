import { ChangeDetectorRef, Component, computed, effect, inject, OnInit } from '@angular/core';
import { ClusterIngressService } from '../../services/cluster-ingress.service';
import { CommonModule, KeyValuePipe } from '@angular/common';

@Component({
  selector: 'app-cluster-ingress-annotations',
  standalone: true,
  imports: [CommonModule, KeyValuePipe],
  templateUrl: './cluster-ingress-annotations.component.html',
  styleUrl: './cluster-ingress-annotations.component.scss',
})
export class ClusterIngressAnnotationsComponent implements OnInit {
  resource: any | undefined;

  private changeDetector = inject(ChangeDetectorRef);
  private clusterIngressService = inject(ClusterIngressService);

  constructor() {
    effect(() => {
      this.resource = this.clusterIngressService.getIngress();
      this.changeDetector.detectChanges();
    });
    computed(() => {
      this.resource = this.clusterIngressService.getIngress();
      this.changeDetector.detectChanges();
    });
  }

  ngOnInit(): void {}
}

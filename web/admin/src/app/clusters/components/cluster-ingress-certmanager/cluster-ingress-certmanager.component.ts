import { ChangeDetectionStrategy, ChangeDetectorRef, Component, effect, inject } from '@angular/core';
import { ClusterIngressService } from '../../services/cluster-ingress.service';
import { Resource } from '@rork8s/ror-resources/models';

@Component({
  selector: 'app-cluster-ingress-certmanager',
  standalone: true,
  imports: [],
  templateUrl: './cluster-ingress-certmanager.component.html',
  styleUrl: './cluster-ingress-certmanager.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ClusterIngressCertmanagerComponent {
  certificates: Resource[] = [];

  private changeDetectorRef = inject(ChangeDetectorRef);
  private clusterIngressService = inject(ClusterIngressService);

  constructor() {
    effect(() => {
      this.certificates = this.clusterIngressService.getCertificates();
      this.changeDetectorRef.detectChanges();
    });
  }
}

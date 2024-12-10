import { ChangeDetectionStrategy, ChangeDetectorRef, Component, effect, inject } from '@angular/core';
import { ClusterIngressService } from '../../services/cluster-ingress.service';
import { Resource } from '@rork8s/ror-resources/models';
import { TableModule } from 'primeng/table';
import { SharedModule } from '../../../shared/shared.module';
import { ButtonModule } from 'primeng/button';
import { TranslateModule } from '@ngx-translate/core';

@Component({
  selector: 'app-cluster-ingress-certmanager',
  standalone: true,
  imports: [TableModule, ButtonModule, SharedModule, TranslateModule],
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

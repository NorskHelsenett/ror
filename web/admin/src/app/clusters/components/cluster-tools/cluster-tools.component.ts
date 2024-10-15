import { ClipboardService } from 'ngx-clipboard';
import { Component, Input, OnInit } from '@angular/core';
import { MessageService } from 'primeng/api';
import { TranslateService } from '@ngx-translate/core';
import { Ingress } from '../../../core/models/ingress';

@Component({
  selector: 'app-cluster-tools',
  templateUrl: './cluster-tools.component.html',
  styleUrls: ['./cluster-tools.component.scss'],
})
export class ClusterToolsComponent implements OnInit {
  @Input() cluster: any = undefined;
  @Input() userClaims: any = undefined;

  argoIngress: any;
  grafanaIngress: any;

  serverUrl: string;

  showLogin = false;

  constructor(
    private clipboardService: ClipboardService,
    private messageService: MessageService,
    private translateService: TranslateService,
  ) {}

  ngOnInit(): void {
    if (!this.cluster) {
      return;
    }

    this.cluster?.ingresses?.forEach((ingress: Ingress) => {
      if (ingress?.name === 'argocd-server') {
        this.argoIngress = ingress;
      }
      if (ingress?.name === 'grafana-helsenett') {
        this.grafanaIngress = ingress;
      }
    });

    if (this.cluster?.workspace?.datacenter?.apiEndpoint && this.cluster?.workspace?.datacenter?.apiEndpoint?.length > 0) {
      this.serverUrl = this.cluster?.workspace?.datacenter?.apiEndpoint;
    } else {
      this.serverUrl = '<missing>';
    }
  }

  copyLogin(): void {
    this.clipboardService.copyFromContent(
      `kubectl vsphere login --server=${this.serverUrl} -u ${this.userClaims?.email} --insecure-skip-tls-verify --tanzu-kubernetes-cluster-namespace ${this.cluster?.workspace?.name} --tanzu-kubernetes-cluster-name ${this.cluster?.clusterName}`,
    );
    this.messageService.add({ severity: 'success', summary: this.translateService.instant('pages.clusters.details.tools.kubectlLoginCopied') });
  }

  copyRorLogin(): void {
    this.clipboardService.copyFromContent(`ror login ${this.cluster?.clusterId}`);
    this.messageService.add({ severity: 'success', summary: this.translateService.instant('pages.clusters.details.tools.rorlogincopied') });
  }

  toggleShowLogin(): void {
    this.showLogin = !this.showLogin;
  }
}

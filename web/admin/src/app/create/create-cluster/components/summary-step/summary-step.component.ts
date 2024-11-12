import { TranslateService } from '@ngx-translate/core';
import { MessageService } from 'primeng/api';
import { ChangeDetectionStrategy, ChangeDetectorRef, Component, Input, OnInit } from '@angular/core';
import { FormGroup } from '@angular/forms';
import { ClusterFormService } from '../../services/cluster-form.service';
import { ActivatedRoute, Router } from '@angular/router';
import { OAuthService } from 'angular-oauth2-oidc';
import { Subscription, catchError } from 'rxjs';
import { environment } from '../../../../../environments/environment';
import { ClusterProvider } from '../../../../clusters/models/clusterProvider';
import { ClusterOrderModel, ClusterOrderType, ProviderConfig } from '../../../../core/models/clusterOrder';
import { OrderService } from '../../../../core/services/order.service';

@Component({
  selector: 'app-summary-step',
  templateUrl: './summary-step.component.html',
  styleUrls: ['./summary-step.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class SummaryStepComponent implements OnInit {
  @Input() clusterForm: FormGroup = this.clusterFormService.clusterForm;

  account: any | undefined;
  createError: any | undefined;

  private subscriptions = new Subscription();

  constructor(
    private changeDetector: ChangeDetectorRef,
    private router: Router,
    private route: ActivatedRoute,
    private clusterFormService: ClusterFormService,
    private oauthService: OAuthService,
    private orderService: OrderService,
    private messageService: MessageService,
    private translateService: TranslateService,
  ) {}

  ngOnInit(): void {
    if (this.clusterFormService.clusterForm.pristine) {
      this.router.navigate(['../'], { relativeTo: this.route });
    }

    this.account = this.oauthService.getIdentityClaims();
  }
  createCluster(): void {
    this.createError = undefined;
    this.changeDetector.detectChanges();

    let clusterFormValue = this.clusterForm.value;
    if (!this.validForm()) {
      this.clusterForm.markAllAsTouched();
      this.changeDetector.detectChanges();
      this.messageService.add({
        severity: 'error',
        summary: this.translateService.instant('pages.create.cluster.steps.summary.errorSubmit'),
      });
      return;
    }

    let nodepools = [];
    let count = 0;
    for (let nodepool of clusterFormValue?.capasity) {
      count += count + 1;
      nodepools.push({
        name: `workers-${count}`,
        machineClass: nodepool?.machineClass,
        count: +nodepool?.count,
      });
    }

    const tags = this.createTagArray();
    let availability = false;
    if (clusterFormValue?.availability?.value) {
      availability = true;
    }

    const providerValue = clusterFormValue?.provider?.type;
    let providerConfig: any = clusterFormValue?.providerConfig;
    if (!providerConfig) {
      return;
    }

    let providerConfigInput = {};
    if (providerValue?.toLowerCase() === 'tanzu') {
      providerConfigInput = {
        namespaceId: providerConfig?.tanzu?.workspace?.id,
        datacenterId: providerConfig?.tanzu?.workspace?.datacenter?.id,
      };
    } else if (providerValue?.toLowerCase() === 'aks') {
      providerConfigInput = {
        subscriptionId: providerConfig?.azure?.subscription?.id,
        resourceGroupName: providerConfig?.azure?.resourceGroup?.name,
        region: providerConfig?.region?.azure?.name,
      };
    } else if (providerValue?.toLowerCase() === 'k3d' || providerValue?.toLowerCase() === 'kind') {
      providerConfigInput = {};
    }

    let clusterOrder: ClusterOrderModel = {
      orderType: ClusterOrderType.Create,
      provider: providerValue,
      projectId: clusterFormValue?.project?.id,
      cluster: clusterFormValue?.environment?.prefix + '-' + clusterFormValue?.clusterName,
      criticality: clusterFormValue?.criticality?.value,
      environment: clusterFormValue?.environment?.value,
      orderBy: clusterFormValue?.orderBy,
      highAvailability: availability,
      ownerGroup: clusterFormValue?.ownergroup,
      sensitivity: clusterFormValue?.sensitivity?.value,
      serviceTags: Object.fromEntries(tags),
      nodePools: nodepools,
      providerconfig: providerConfigInput,
      k8sVersion: clusterFormValue?.k8sVersion?.version,
    };

    this.subscriptions.add(
      this.orderService
        .orderCluster(clusterOrder)
        .pipe(
          catchError((error: any) => {
            this.createError = error;
            this.messageService.add({
              severity: 'error',
              summary: this.translateService.instant('pages.create.cluster.steps.summary.errorSubmit'),
            });
            this.changeDetector.detectChanges();
            throw error;
          }),
        )
        .subscribe((data: boolean) => {
          this.router.navigate(['../../../'], { relativeTo: this.route });
          this.messageService.add({
            severity: 'success',
            closable: true,
            summary: this.translateService.instant('pages.create.cluster.steps.summary.submitSuccess'),
          });
          this.changeDetector.detectChanges();
        }),
    );
  }

  validForm(): boolean {
    const providerValid = this.clusterFormService?.clusterForm?.get('provider')?.valid;
    const ownerGroupValid = this.clusterFormService?.clusterForm?.get('ownergroup').valid;

    const validMap = this.getValidMap();
    let hasFalse = Array.from(validMap.values()).some((value) => value === false);

    let providerConfigValid = false;
    if (this.clusterFormService?.selectedProvider?.type?.toLocaleLowerCase() === ClusterProvider.PrivatSky?.toLocaleLowerCase()) {
      providerConfigValid = this.clusterFormService?.clusterForm?.get('providerConfig')?.get('tanzu')?.valid;
    } else if (this.clusterFormService?.selectedProvider?.type?.toLocaleLowerCase() === ClusterProvider.Talos?.toLocaleLowerCase()) {
      providerConfigValid = true;
    } else if (this.clusterFormService?.selectedProvider?.type?.toLocaleLowerCase() === ClusterProvider.K3d?.toLocaleLowerCase()) {
      providerConfigValid = true;
      hasFalse = false;
    } else if (this.clusterFormService?.selectedProvider?.type?.toLocaleLowerCase() === ClusterProvider.Kind?.toLocaleLowerCase()) {
      providerConfigValid = true;
      hasFalse = false;
    } else if (
      this.clusterFormService?.selectedProvider?.type?.toLocaleLowerCase() === ClusterProvider.Unknown?.toLocaleLowerCase() &&
      !environment.production
    ) {
      providerConfigValid = this.clusterFormService?.clusterForm?.get('providerConfig')?.get('tanzu')?.valid;
    }

    return providerValid && providerConfigValid && !hasFalse && ownerGroupValid;
  }

  private createTagArray(): Map<string, string> {
    let tags: Map<string, string> = new Map();
    const formTags = this.clusterForm?.get('tags').getRawValue();
    if (!formTags || formTags?.length == 0) {
      return tags;
    }

    formTags.forEach((tag: string) => {
      tags.set(tag, '');
    });

    return tags;
  }

  private getValidMap(): Map<string, boolean> {
    let map = new Map<string, boolean>();
    Object.keys(this.clusterForm.controls).forEach((key) => {
      if (key !== 'providerConfig') {
        map.set(key, this.clusterForm.controls[key].valid);
      }
    });
    return map;
  }
}

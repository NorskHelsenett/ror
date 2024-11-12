import { ChangeDetectionStrategy, Component, Input, OnInit, ChangeDetectorRef } from '@angular/core';
import { FormGroup } from '@angular/forms';
import { ClusterFormService } from '../../services/cluster-form.service';
import { ActivatedRoute, Router } from '@angular/router';
import { Observable, tap, Subscription, map } from 'rxjs';
import { TranslateService } from '@ngx-translate/core';
import { ClusterCapasity } from '../../../../clusters/models/clusterCapasity';
import { ClusterEnvironment } from '../../../../core/models/clusterEnvironment';
import { Price } from '../../../../core/models/price';
import { ProviderKubernetesVersion } from '../../../../core/models/provider';
import { PriceService } from '../../../../core/services/price.service';
import { ProvidersService } from '../../../../core/services/providers.service';
import { ClusterProvider } from '../../../../clusters/models/clusterProvider';

@Component({
  selector: 'app-resources-step',
  templateUrl: './resources-step.component.html',
  styleUrls: ['./resources-step.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ResourcesStepComponent implements OnInit {
  @Input() clusterForm: FormGroup = this.clusterFormService.clusterForm;

  nodePools: ClusterCapasity[] = [];

  prices$: Observable<any> | undefined;
  pricesFiltered: any[] = [];
  selectedPrices: Price[] = [];

  haOptions: any[] = [];
  showEnvironmentWarning = false;

  clusterName: string;

  environments: any[] = [
    {
      name: ClusterEnvironment[ClusterEnvironment.Development],
      value: ClusterEnvironment.Development,
      prefix: 'd',
    },
    {
      name: ClusterEnvironment[ClusterEnvironment.Testing],
      value: ClusterEnvironment.Testing,
      prefix: 't',
    },
    {
      name: ClusterEnvironment[ClusterEnvironment.QA],
      value: ClusterEnvironment.QA,
      prefix: 'q',
    },
    {
      name: ClusterEnvironment[ClusterEnvironment.Production],
      value: ClusterEnvironment.Production,
      prefix: 'p',
    },
  ];

  k8sVersions: any[] = [];
  k8sVersions$: Observable<ProviderKubernetesVersion[]> | undefined;

  private pricesAll: any[] = [];
  private subscriptions = new Subscription();

  constructor(
    private changeDetector: ChangeDetectorRef,
    private router: Router,
    private route: ActivatedRoute,
    private clusterFormService: ClusterFormService,
    private priceService: PriceService,
    private translateService: TranslateService,
    private providersService: ProvidersService,
  ) {}

  ngOnInit(): void {
    if (this.clusterFormService?.clusterForm?.pristine) {
      this.router.navigate(['../'], { relativeTo: this.route });
      this.changeDetector.detectChanges();
    }
    this.fetchK8sVersions();

    this.haOptions = this.setupHaOptions();
    this.setHaSelection();

    this.subscriptions.add(
      this.translateService.onLangChange
        .pipe(
          tap((event: any) => {
            this.haOptions = this.setupHaOptions();
            this.setHaSelection();

            this.changeDetector.detectChanges();
          }),
        )
        .subscribe(),
    );

    this.fetchPrices();

    this.setK8sVersionSelection();

    this.changeDetector.detectChanges();
  }

  fetchK8sVersions(): void {
    this.k8sVersions$ = this.providersService.getKubernetesVersionByProviderType(this.clusterForm?.get('provider')?.value?.type).pipe(
      tap((data: ProviderKubernetesVersion[]) => {
        this.k8sVersions = data?.sort((a: any, b: any) => (a?.name > b?.name ? -1 : 1));

        if (this.k8sVersions?.length > 0) {
          const index = this.k8sVersions.findIndex((x: any) => x.disabled === false);
          this.clusterForm?.get('k8sVersion')?.setValue(this.k8sVersions[index]);
        }
        this.changeDetector.detectChanges();
        return this.k8sVersions;
      }),
    );
  }

  addNodePool(price: any): void {
    if (!this.nodePools) {
      this.nodePools = [];
    }
    this.nodePools.push({
      count: price.count,
      machineClass: price.machineClass,
      price: price.price,
    });
    this.clusterForm.patchValue({ capasity: this.nodePools });
    this.changeDetector.detectChanges();
  }

  removeNodePool(nodePool: any): void {
    this.nodePools = this.nodePools.filter((x: any) => {
      if (x != nodePool) {
        return x;
      }
    });

    if (this.nodePools && this.nodePools.length === 0) {
      this.nodePools = undefined;
      this.clusterForm.patchValue({ capasity: null });
    } else {
      this.clusterForm.patchValue({ capasity: this.nodePools });
    }

    this.changeDetector.detectChanges();
  }

  getNodePoolSum(): number {
    return this.clusterFormService.getNodePoolSum();
  }

  environmentChanged(event: any): void {
    const env = this.clusterForm?.get('environment')?.value;
    this.showEnvironmentWarning = env?.value == ClusterEnvironment.Production && event?.value?.value === false;
  }

  clusterNameChanged(): void {
    const clusterName = this.clusterForm.get('clusterName').value;
    let clusterEnv = this.clusterForm.get('environment').value;
    let clusterId: string;

    if (!clusterEnv) {
      clusterEnv = this.environments[0];
    }

    if (clusterName !== null && clusterName !== '') {
      clusterId = ` ${clusterEnv?.prefix}-${clusterName}`;
    } else {
      clusterId = '';
    }

    this.clusterName = clusterId;
    this.changeDetector.detectChanges();
  }

  updateEvent(event: any): void {
    this.changeDetector.detectChanges();
  }

  validform(): boolean {
    const environmentValid = this.clusterFormService?.clusterForm?.get('environment')?.valid;
    const clusterNameValid = this.clusterFormService?.clusterForm?.get('clusterName')?.valid;
    const capasityValid = this.clusterFormService?.clusterForm?.get('capasity')?.valid;
    const availabilityValid = this.clusterFormService?.clusterForm?.get('availability')?.valid;
    const k8sVersionValid = this.clusterFormService?.clusterForm?.get('k8sVersion')?.valid;

    var valid = false;
    if (
      this.clusterFormService?.selectedProvider?.type?.toLocaleLowerCase() === ClusterProvider.K3d?.toLocaleLowerCase() ||
      this.clusterFormService?.selectedProvider?.type?.toLocaleLowerCase() === ClusterProvider.Kind?.toLocaleLowerCase()
    ) {
      valid = clusterNameValid && capasityValid;
    } else {
      valid = environmentValid && clusterNameValid && capasityValid && availabilityValid && k8sVersionValid;
    }

    return valid;
  }

  private fetchPrices(): void {
    this.prices$ = this.priceService.getAll().pipe(
      tap((prices: Price[]) => {
        prices.forEach((price: any) => {
          if (price?.machineClass?.indexOf('small') >= 0) {
            return;
          }
          this.pricesAll.push({ ...price, count: 3 });
          this.pricesFiltered = this.pricesAll.filter((x: any) => {
            if (x?.provider === 'tanzu') {
              return x;
            }
          });
        });
        this.setupNodePools();
        this.changeDetector.detectChanges();
      }),
    );
  }

  private setupNodePools(): void {
    const capacity = this.clusterForm?.get('capasity')?.value;
    this.nodePools = [];
    if (capacity) {
      for (let i = 0; i < capacity?.length; i++) {
        const c = capacity[i];
        this.nodePools.push(c);
      }
    }
  }

  private setupHaOptions(): any[] {
    const env = this.clusterForm?.get('environment')?.value;
    let options: any[] = [];
    if (env?.value == ClusterEnvironment.Production) {
      options = [{ label: this.translateService.instant('pages.create.cluster.steps.resources.highAvailability'), value: true, price: 2114 }];
    } else {
      options = [
        { label: this.translateService.instant('pages.create.cluster.steps.resources.lowAvailability'), value: false, price: 0 },
        { label: this.translateService.instant('pages.create.cluster.steps.resources.highAvailability'), value: true, price: 2114 },
      ];
    }

    return options;
  }

  private setHaSelection(): void {
    const env = this.clusterFormService?.clusterForm?.get('environment')?.value;
    this.clusterFormService?.clusterForm?.get('availability')?.setValue(this.haOptions[0]);
    if (env?.value == ClusterEnvironment.Production) {
      this.clusterFormService?.clusterForm?.get('availability')?.setValue(this.haOptions[0]);
    }
  }

  private setK8sVersionSelection(): void {
    const selectedK8sVersion = this.clusterFormService?.clusterForm?.get('k8sVersion')?.value;
    if (selectedK8sVersion) {
      return;
    }

    let availableK8sversion = this.k8sVersions.filter((x: any) => x.disabled == false);
    if (availableK8sversion?.length === 0) {
      return;
    }
    this.clusterFormService?.clusterForm?.get('k8sVersion')?.setValue(availableK8sversion[0]);
  }
}

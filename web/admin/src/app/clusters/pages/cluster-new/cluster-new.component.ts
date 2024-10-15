import { ChangeDetectionStrategy, Component, OnInit, ChangeDetectorRef, ViewEncapsulation } from '@angular/core';
import { forkJoin, Observable, catchError, tap } from 'rxjs';
import { ClusterModel } from '../../models/clusterModel';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { OAuthService } from 'angular-oauth2-oidc';
import { ClusterCapasity } from '../../models/clusterCapasity';
import { Price } from '../../../core/models/price';
import { ClusterProvider } from '../../models/clusterProvider';
import { DatacenterService } from '../../../core/services/datacenter.service';
import { WorkspacesService } from '../../../core/services/workspaces.service';
import { PriceService } from '../../../core/services/price.service';
import { ClusterEnvironment } from '../../../core/models/clusterEnvironment';
import { ClustersService } from '../../../core/services/clusters.service';
import ClusterNameValidator from '../../../shared/validators/clusterNameValidator';

@Component({
  selector: 'app-cluster-new',
  templateUrl: './cluster-new.component.html',
  styleUrls: ['./cluster-new.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
  encapsulation: ViewEncapsulation.None,
})
export class ClusterNewComponent implements OnInit {
  clusterForm: FormGroup | undefined;

  data$: Observable<any> | undefined;
  dataError: any;

  datacenters$: Observable<any> | undefined;
  datacenters: any[] = [];
  workspaces$: Observable<any> | undefined;
  workspacesAll: any[] = [];
  workspacesFiltered: any[] = [];
  prices$: Observable<any> | undefined;

  pricesFiltered: any[] = [];
  selectedPrices: Price[] = [];
  environments: any[] = [
    {
      name: ClusterEnvironment[ClusterEnvironment.Development],
      value: ClusterEnvironment.Development,
    },
    {
      name: ClusterEnvironment[ClusterEnvironment.Testing],
      value: ClusterEnvironment.Testing,
    },
    {
      name: ClusterEnvironment[ClusterEnvironment.QA],
      value: ClusterEnvironment.QA,
    },
    {
      name: ClusterEnvironment[ClusterEnvironment.Production],
      value: ClusterEnvironment.Production,
    },
  ];

  account: any | undefined;

  submitted = false;
  clusterCreateModel: ClusterModel | undefined;

  createResponse$: Observable<any> | undefined;
  createError: any;

  providerOptions: any[] = [
    {
      name: ClusterProvider[ClusterProvider.PrivatSky].toString(),
      value: ClusterProvider.PrivatSky,
    },
    // {
    //   name: ClusterProvider[ClusterProvider.Azure].toString(),
    //   value: ClusterProvider.Azure,
    // },
  ];

  selectedClusterName: string = '';
  defaultNodeCountValue = 3;
  nodePools: ClusterCapasity[];

  private pricesAll: any[] = [];

  constructor(
    private changeDetector: ChangeDetectorRef,
    private fb: FormBuilder,
    private priceService: PriceService,
    private workspacesService: WorkspacesService,
    private clustersService: ClustersService,
    private datacenterService: DatacenterService,
    private oauthService: OAuthService,
  ) {}

  ngOnInit(): void {
    this.setupObservables();
    this.fetchData();
    this.account = this.oauthService.getIdentityClaims();

    this.clusterForm = this.fb.group({
      provider: [null, [Validators.required]],
      datacenter: [null, [Validators.required, Validators.min(2)]],
      workspace: [null, [Validators.required, Validators.min(2)]],
      clusterNameInput: [null, []],
      clusterName: [null, [Validators.required, Validators.min(2)], [ClusterNameValidator.validName(this.clustersService)]],
      tags: [[], []],
      environment: [null, [Validators.required]],
      accessGroups: [[], []],
      capasity: [null, [Validators.required, Validators.minLength(1)]],
      project: [null, []],
      responsible: [this.account?.email, [Validators.required]],
    });

    this.clusterForm.patchValue({
      provider: this.providerOptions[0]?.value,
      environment: this.environments[0]?.value,
    });

    this.changeDetector.detectChanges();
  }

  setupObservables(): void {
    this.datacenters$ = this.datacenterService.get();
    this.prices$ = this.priceService.getAll();
    this.workspaces$ = this.workspacesService.get();
  }

  fetchData(): void {
    this.dataError = undefined;
    this.data$ = forkJoin([this.datacenters$, this.workspaces$, this.prices$]).pipe(
      tap((results: any[]) => {
        if (!results || results.length === 0) {
          this.dataError = 'Empty result';
          return;
        }
        this.datacenters = results[0];
        this.workspacesAll = results[1];
        const prices = results[2];
        prices.forEach((price: any) => {
          this.pricesAll.push({ ...price, count: 3 });
          this.pricesFiltered = this.pricesAll.filter((x: any) => {
            if (x?.provider === 'tanzu') {
              return x;
            }
          });
        });

        this.changeDetector.detectChanges();
      }),
      catchError((error) => {
        this.dataError = error;
        this.changeDetector.detectChanges();
        return error;
      }),
    );
  }

  providerChanged(event: any): void {
    const selectedProvider = this.clusterForm.controls['provider'].value;
    if (selectedProvider === ClusterProvider.PrivatSky) {
      this.pricesFiltered = this.pricesAll.filter((x: any) => {
        if (x?.provider === 'tanzu') {
          return x;
        }
      });
    } else if (selectedProvider === ClusterProvider.Azure) {
      this.pricesFiltered = this.pricesAll.filter((x: any) => {
        if (x?.provider === 'azure') {
          return x;
        }
      });
    } else {
      this.pricesFiltered = [];
    }
    this.changeDetector.detectChanges();
  }

  createCluster(): void {
    this.createError = undefined;
    this.clusterCreateModel = this.clusterForm.value as ClusterModel;
    delete this.clusterCreateModel['clusterNameInput'];
    if (!this.clusterForm.valid) {
      this.clusterForm.markAllAsTouched();
      this.changeDetector.detectChanges();
      return;
    }
    this.submitted = true;
    this.clusterCreateModel = this.clusterForm.value as ClusterModel;
    delete this.clusterCreateModel['clusterNameInput'];
    this.changeDetector.detectChanges();
  }

  reset(formReset?: boolean): void {
    this.submitted = false;
    this.clusterCreateModel = undefined;
    this.createError = undefined;
    this.createResponse$ = undefined;
    this.nodePools = undefined;
    if (formReset === true) {
      this.clusterForm.reset();
      this.clusterForm.patchValue({
        provider: this.providerOptions[0]?.value,
        environment: this.environments[0]?.value,
        tags: null,
        capasity: null,
        clusterName: null,
      });
    }
    this.changeDetector.detectChanges();
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
    let sum = 0;
    if (!this.nodePools || this.nodePools.length === 0) {
      return sum;
    }

    this.nodePools.forEach((x: any) => {
      sum = sum + x.price * x.count;
    });

    return sum;
  }

  datacenterChanged(event: any): void {
    this.workspacesFiltered = this.workspacesAll.filter((x: any) => {
      if (x?.name?.indexOf(event?.value?.name) !== -1) {
        return x;
      }
    });

    if (this.workspacesFiltered?.length === 1) {
      this.clusterForm.patchValue({ workspace: this.workspacesFiltered[0] });
    }
    this.clusterNameChanged();
  }

  clusterNameChanged(): void {
    const clusterName = this.clusterForm.get('clusterNameInput').value;
    let workspace = this.clusterForm.get('workspace').value;
    let clusterId: string;

    if (clusterName !== null && clusterName !== '' && workspace !== null && workspace !== '') {
      workspace = workspace.name;
      clusterId = `${clusterName}.${workspace}`;
    } else {
      clusterId = '';
    }

    this.clusterForm.get('clusterName').setValue(clusterId);
    this.clusterForm.get('clusterName').markAsDirty();
    this.clusterForm.get('clusterName').markAsTouched();
    this.changeDetector.detectChanges();
  }
}

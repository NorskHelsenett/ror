import { FormGroup } from '@angular/forms';
import { ChangeDetectionStrategy, ChangeDetectorRef, Component, Input, OnDestroy, OnInit } from '@angular/core';
import { tap, Subscription, Observable, catchError, map } from 'rxjs';
import { ActivatedRoute, Router } from '@angular/router';

import { WorkspacesService } from '../../../../core/services/workspaces.service';
import { ClusterFormService } from '../../services/cluster-form.service';
import { OAuthService } from 'angular-oauth2-oidc';
import { ClusterProvider } from '../../../../clusters/models/clusterProvider';
import { PaginationResult } from '../../../../core/models/paginatedResult';
import { Project } from '../../../../core/models/project';
import { Workspace } from '../../../../core/models/workspace';
import { ProjectService } from '../../../../core/services/project.service';
import { ProvidersService } from '../../../../core/services/providers.service';
import { Provider } from '../../../../core/models/provider';

@Component({
  selector: 'app-location-step',
  templateUrl: './location-step.component.html',
  styleUrls: ['./location-step.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class LocationStepComponent implements OnInit, OnDestroy {
  @Input() clusterForm: FormGroup = this.clusterFormService.clusterForm;

  workspaces: Workspace[] = [];

  azureSubscriptions: any[] = [];
  azureResourceGroups: any[] = [];
  azureRegions: any[] = [];

  providers$: Observable<Provider[]> | undefined;
  providersError: any | undefined;

  projects$: Observable<PaginationResult<Project>> | undefined;
  account: any | undefined;

  private subscriptions = new Subscription();

  constructor(
    private changeDetector: ChangeDetectorRef,
    private router: Router,
    private route: ActivatedRoute,
    private clusterFormService: ClusterFormService,
    private workspacesService: WorkspacesService,
    private projectService: ProjectService,
    private oauthService: OAuthService,
    private providersService: ProvidersService,
  ) {}

  ngOnInit(): void {
    if (this.clusterFormService.clusterForm.pristine) {
      this.router.navigate(['../'], { relativeTo: this.route });
    }

    this.fetchProviders();

    this.account = this.oauthService?.getIdentityClaims();
    this.clusterForm.get('orderBy').setValue(this.account?.email);

    this.fetchProjects();

    this.clusterFormService.clusterForm?.get('provider')?.valueChanges.subscribe((value: any) => {
      this.providerChanged(value);
    });

    this.changeDetector.detectChanges();
  }

  ngOnDestroy(): void {
    this.subscriptions.unsubscribe();
  }

  fetchProviders(): void {
    this.providers$ = this.providersService.get().pipe(
      map((providers: Provider[]) => {
        providers = providers?.sort((a: Provider, b: Provider) => (a?.disabled === true ? 1 : -1));
        providers.sort(function (p, p2) {
          if (p.name < p2.name) {
            return -1;
          }
          if (p.name > p2.name) {
            return 1;
          }
          return 0;
        });
        this.changeDetector.detectChanges();
        return providers;
      }),
    );
  }

  fetchProjects(): void {
    this.projects$ = this.projectService
      .getByFilter({
        limit: 100,
        skip: 0,
      })
      .pipe(
        tap(() => this.changeDetector.detectChanges()),
        catchError((error) => {
          this.changeDetector.detectChanges();
          throw error;
        }),
      );
  }

  providerChanged(event: any): void {
    this.clusterFormService.selectedProvider = this.clusterForm?.controls['provider']?.value;

    if (this.clusterFormService.selectedProvider?.type?.toLocaleLowerCase() === ClusterProvider.PrivatSky.toLocaleLowerCase()) {
      this.fetchWorkspaces(ClusterProvider.PrivatSky);
    } else if (this.clusterFormService.selectedProvider?.type?.toLocaleLowerCase() === ClusterProvider.Azure.toLocaleLowerCase()) {
      // todo fetch azure stuff
    } else if (this.clusterFormService.selectedProvider?.type?.toLocaleLowerCase() === ClusterProvider.Unknown.toLocaleLowerCase()) {
      this.fetchWorkspaces(ClusterProvider.Unknown);
    } else if (this.clusterFormService.selectedProvider?.type?.toLocaleLowerCase() === ClusterProvider.K3d.toLocaleLowerCase()) {
      // todo fetch k3d stuff
    } else if (this.clusterFormService.selectedProvider?.type?.toLocaleLowerCase() === ClusterProvider.Talos.toLocaleLowerCase()) {
      // todo fetch talos stuff
    }
    this.changeDetector.detectChanges();
  }

  fetchWorkspaces(provider: ClusterProvider): void {
    this.subscriptions.add(
      this.workspacesService
        .get()
        .pipe(
          map((workspaces: Workspace[]) => {
            const filtered = workspaces?.filter((x: Workspace) => x?.datacenter?.provider?.toLowerCase() === provider?.toLowerCase());
            this.workspaces = filtered?.sort((a: Workspace, b: Workspace) => (a?.name < b?.name ? -1 : 1));
            this.changeDetector.detectChanges();
            return this.workspaces;
          }),
        )
        .subscribe(),
    );
  }

  updateEvent(event: any): void {
    this.changeDetector.detectChanges();
  }

  get providerConfig(): FormGroup {
    return this.clusterFormService?.clusterForm?.get('providerConfig') as FormGroup;
  }

  validForm(): boolean {
    const providerValid = this.clusterFormService?.clusterForm?.get('provider')?.valid;
    const projectValid = this.clusterFormService?.clusterForm?.get('project')?.valid;

    let providerConfigValid = false;
    if (this.clusterFormService?.selectedProvider?.type?.toLocaleLowerCase() === ClusterProvider.PrivatSky?.toLocaleLowerCase()) {
      providerConfigValid = this.clusterFormService?.clusterForm?.get('providerConfig')?.get('tanzu')?.valid;
    } else if (this.clusterFormService?.selectedProvider?.type?.toLocaleLowerCase() === ClusterProvider.Azure?.toLocaleLowerCase()) {
      providerConfigValid = this.clusterFormService?.clusterForm?.get('providerConfig')?.get('azure')?.valid;
    } else if (this.clusterFormService?.selectedProvider?.type?.toLocaleLowerCase() === ClusterProvider.K3d?.toLocaleLowerCase()) {
      providerConfigValid = true;
    } else if (this.clusterFormService?.selectedProvider?.type?.toLocaleLowerCase() === ClusterProvider.Kind?.toLocaleLowerCase()) {
      providerConfigValid = true;
    } else if (this.clusterFormService?.selectedProvider?.type?.toLocaleLowerCase() === ClusterProvider.Talos?.toLocaleLowerCase()) {
      providerConfigValid = true;
    }

    return providerValid && providerConfigValid && projectValid;
  }
}

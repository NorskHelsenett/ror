import { ChangeDetectionStrategy, ChangeDetectorRef, Component, Input, OnDestroy, OnInit } from '@angular/core';
import { FormGroup } from '@angular/forms';
import { ClusterFormService } from '../../services/cluster-form.service';
import { Observable, Subscription, catchError, share, tap } from 'rxjs';
import { OAuthService } from 'angular-oauth2-oidc';
import { TranslateService } from '@ngx-translate/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AclScopes, AclAccess } from '../../../../core/models/acl-scopes';
import { AclService } from '../../../../core/services/acl.service';

@Component({
  selector: 'app-metadata-step',
  templateUrl: './metadata-step.component.html',
  styleUrls: ['./metadata-step.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class MetadataStepComponent implements OnInit, OnDestroy {
  @Input() clusterForm: FormGroup = this.clusterFormService.clusterForm;

  availableCriticalities: any[];
  selectedCriticality: any;
  availableSensitivities: any[];
  selectedSensitivity: any;

  account: any | undefined;

  accountClaims: any | undefined;
  accessGroups: any[] = [];
  ownergroupEditable = false;
  adminOwner$: Observable<boolean> | undefined;
  aclFetchError: any;

  private subscriptions = new Subscription();

  constructor(
    private changeDetector: ChangeDetectorRef,
    private router: Router,
    private route: ActivatedRoute,
    private clusterFormService: ClusterFormService,
    private oauthService: OAuthService,
    private translateService: TranslateService,
    private aclService: AclService,
  ) {}

  ngOnInit(): void {
    if (this.clusterFormService?.clusterForm?.pristine) {
      this.router.navigate(['../'], { relativeTo: this.route });
      this.changeDetector.detectChanges();
    }

    this.fetchAcl();

    this.setupCriticalityAndSensitivity();
    this.subscriptions.add(this.translateService.onLangChange.subscribe(() => this.setupCriticalityAndSensitivity()));

    let domain = '';
    this.accountClaims = this.oauthService?.getIdentityClaims();
    const emailArray = this.accountClaims?.email?.split('@');
    if (emailArray?.length > 1) {
      domain = emailArray[1];
    }
    this.accessGroups = this.accountClaims?.groups.map((x: any) => {
      return { name: `${x}@${domain}` };
    });
    this.accessGroups = this.accessGroups.filter((x: any) => x?.name !== null);

    this.changeDetector.detectChanges();
  }

  ngOnDestroy(): void {
    if (this.subscriptions) {
      this.subscriptions.unsubscribe();
    }
  }

  fetchAcl(): void {
    this.adminOwner$ = this.aclService.check(AclScopes.ROR, AclScopes.Global, AclAccess.Owner).pipe(
      share(),
      tap((isAdminOwner: boolean) => {
        this.ownergroupEditable = isAdminOwner;
      }),
      catchError((error: any) => {
        this.aclFetchError = error;
        this.changeDetector.detectChanges();
        throw error;
      }),
    );
  }

  setupCriticalityAndSensitivity(): void {
    this.availableCriticalities = [
      {
        name: this.translateService.instant('pages.clusters.details.edit.form.availableCriticalities.1'),
        value: 1,
      },
      {
        name: this.translateService.instant('pages.clusters.details.edit.form.availableCriticalities.2'),
        value: 2,
      },
      {
        name: this.translateService.instant('pages.clusters.details.edit.form.availableCriticalities.3'),
        value: 3,
      },
      {
        name: this.translateService.instant('pages.clusters.details.edit.form.availableCriticalities.4'),
        value: 4,
      },
    ];

    this.availableSensitivities = [
      {
        name: this.translateService.instant('pages.clusters.details.edit.form.availableSensitivities.1'),
        value: 1,
      },
      {
        name: this.translateService.instant('pages.clusters.details.edit.form.availableSensitivities.2'),
        value: 2,
      },
      {
        name: this.translateService.instant('pages.clusters.details.edit.form.availableSensitivities.3'),
        value: 3,
      },
      {
        name: this.translateService.instant('pages.clusters.details.edit.form.availableSensitivities.4'),
        value: 4,
      },
    ];

    let selectedCriticality = this.clusterFormService?.clusterForm?.get('criticality')?.value;
    if (selectedCriticality) {
      const filter = this.availableCriticalities?.filter((x) => x?.value == selectedCriticality?.value);
      if (filter.length === 1) {
        this.clusterFormService?.clusterForm?.get('criticality')?.setValue(filter[0]);
      }
    }

    let selectedSensitivity = this.clusterFormService?.clusterForm?.get('sensitivity')?.value;
    if (selectedSensitivity) {
      const filter = this.availableSensitivities?.filter((x) => x?.value == selectedSensitivity?.value);
      if (filter.length === 1) {
        this.clusterFormService?.clusterForm?.get('sensitivity')?.setValue(filter[0]);
      }
    }

    this.changeDetector.detectChanges();
  }

  validForm(): boolean {
    const ownerGroupValid = this.clusterFormService?.clusterForm?.get('ownergroup')?.valid;
    return ownerGroupValid;
  }

  updateEvent(event: any): void {
    this.changeDetector.detectChanges();
  }
}

import { ChangeDetectionStrategy, ChangeDetectorRef, Component, Input, OnDestroy, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { TranslateService } from '@ngx-translate/core';
import { OAuthService } from 'angular-oauth2-oidc';
import { MessageService } from 'primeng/api';
import { Observable, Subscription, catchError, map } from 'rxjs';
import { AclV2 } from '../../../core/models/aclv2';
import { Cluster } from '../../../core/models/cluster';
import { PaginationResult } from '../../../core/models/paginatedResult';
import { AclService } from '../../../core/services/acl.service';
import { ClustersService } from '../../../core/services/clusters.service';
import { ConfigService } from '../../../core/services/config.service';

@Component({
  selector: 'app-acl-create-update',
  templateUrl: './acl-create-update.component.html',
  styleUrls: ['./acl-create-update.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AclCreateUpdateComponent implements OnInit, OnDestroy {
  @Input() set acl(value: AclV2 | undefined) {
    if (!value) {
      this.permission = undefined;
      return;
    }

    this.permission = value;
    this.permission_orginal = value;
    this.fillForm();
  }

  get acl(): AclV2 {
    return this.permission;
  }

  scopes$: Observable<string[]> | undefined;
  scopes: string[];

  aclForm: FormGroup;
  createError: any;
  updateError: any;
  scopeFetchError: any;
  suggestedGroups: any[] = [];
  suggestedSubjects: any[] = [];
  selectedScope: string;
  userEmail: string;

  private permission: AclV2 | undefined;
  private permission_orginal: AclV2 | undefined;
  private subscriptions = new Subscription();
  private rortextregex = this.configService.config.regex.forms;

  constructor(
    private changeDetector: ChangeDetectorRef,
    private route: ActivatedRoute,
    private router: Router,
    private fb: FormBuilder,
    private aclService: AclService,
    private messageService: MessageService,
    private translateService: TranslateService,
    private oauthService: OAuthService,
    private clustersService: ClustersService,
    private configService: ConfigService,
  ) {}

  ngOnInit(): void {
    this.setupForm();
    this.fetchScopes();
    this.setSuggestedGroups();
    this.subscriptions.add(
      this.translateService.onLangChange.subscribe(() => {
        this.setSuggestedGroups();
        this.changeDetector.detectChanges();
      }),
    );
  }

  ngOnDestroy(): void {
    this.subscriptions.unsubscribe();
  }

  reset(): void {
    this.createError = undefined;
    this.aclForm.reset();
    if (this.permission_orginal) {
      this.aclForm.patchValue(this.permission_orginal);
    }
    this.aclForm.get('subject').disable();
    this.changeDetector.detectChanges();
  }

  create(): void {
    this.createError = undefined;

    let acl: AclV2 = {
      access: {
        create: false,
        delete: false,
        owner: false,
        read: false,
        update: false,
      },
      kubernetes: {
        logon: false,
      },
      created: new Date(),
      group: '',
      issuedBy: '',
      scope: '',
      subject: '',
      version: 2,
      id: '',
    };

    acl = this.aclForm.value as AclV2;
    if (!this.aclForm.valid) {
      this.aclForm.markAllAsTouched();
      this.changeDetector.detectChanges();
      return;
    }

    if (acl?.access?.owner === true) {
      acl = { ...acl, access: { create: true, delete: true, owner: true, read: true, update: true } };
    }

    acl.issuedBy = this.userEmail;

    this.subscriptions.add(
      this.aclService
        .create(acl)
        .pipe(
          catchError((error: any) => {
            this.createError = error;
            this.changeDetector.detectChanges();
            this.messageService.add({
              severity: 'error',
              summary: this.translateService.instant('pages.acl.create.error'),
            });
            throw error;
          }),
        )
        .subscribe((data: AclV2) => {
          this.router.navigate(['../'], { relativeTo: this.route });
          this.messageService.add({
            severity: 'success',
            summary: this.translateService.instant('pages.acl.create.success', { group: data?.group }),
          });
        }),
    );
  }

  update(): void {
    this.updateError = undefined;
    let acl = this.aclForm.value as AclV2;

    if (!this.aclForm.valid) {
      this.aclForm.markAllAsTouched();
      this.changeDetector.detectChanges();
      return;
    }

    if (acl?.access?.owner === true) {
      acl = { ...acl, access: { create: true, delete: true, owner: true, read: true, update: true } };
    }

    acl.issuedBy = this.acl.issuedBy;

    this.subscriptions.add(
      this.aclService
        .update(this.permission?.id, acl)
        .pipe(
          catchError((error: any) => {
            this.createError = error;
            this.changeDetector.detectChanges();
            this.messageService.add({
              severity: 'error',
              summary: this.translateService.instant('pages.acl.update.error'),
            });
            throw error;
          }),
        )
        .subscribe((data: AclV2) => {
          this.router.navigate(['../../'], { relativeTo: this.route });
          this.messageService.add({
            severity: 'success',
            summary: this.translateService.instant('pages.acl.update.success', { group: data?.group }),
          });
        }),
    );
  }

  scopeSelected(scopeEvent: any): void {
    if (!scopeEvent) {
      return;
    }

    this.aclForm.get('subject').enable();

    this.suggestedSubjects = [];
    switch (scopeEvent?.value) {
      case 'cluster': {
        this.clustersToSubject();
        break;
      }
      case 'ror': {
        this.scopesToSubject();
        break;
      }
    }
  }

  ownerOn(event: any): void {
    if (!event) {
      return;
    }

    let isControlDisabled: boolean;

    if (event?.checked) {
      isControlDisabled = true;
      this.aclForm.patchValue({
        access: {
          read: true,
          create: true,
          update: true,
          delete: true,
        },
      });
      this.aclForm.get('access').get('read').disable();
      this.aclForm.get('access').get('create').disable();
      this.aclForm.get('access').get('update').disable();
      this.aclForm.get('access').get('delete').disable();
    } else {
      isControlDisabled = false;
      this.aclForm.patchValue({
        access: {
          read: false,
          create: false,
          update: false,
          delete: false,
        },
      });
      this.aclForm.get('access').get('read').enable();
      this.aclForm.get('access').get('create').enable();
      this.aclForm.get('access').get('update').enable();
      this.aclForm.get('access').get('delete').enable();
    }

    this.changeDetector.detectChanges();
  }

  private setupForm(): void {
    this.aclForm = this.fb.group({
      group: [null, { validators: [Validators.required, Validators.minLength(1), Validators.pattern(this.rortextregex)] }],
      scope: [null, { validators: [Validators.required, Validators.minLength(1), Validators.pattern(this.rortextregex)] }],
      subject: [null, { validators: [Validators.required, Validators.minLength(1), Validators.pattern(this.rortextregex)] }],
      access: this.fb.group({
        read: [false, { validators: [Validators.required] }],
        create: [false, { validators: [Validators.required] }],
        update: [false, { validators: [Validators.required] }],
        delete: [false, { validators: [Validators.required] }],
        owner: [false, { validators: [Validators.required] }],
      }),
      kubernetes: this.fb.group({
        logon: [false, { validators: [Validators.required] }],
      }),
    });
    this.aclForm.get('subject').disable();
  }

  private fillForm(): void {
    this.aclForm.patchValue(this.permission);
    if (this.permission.scope) {
      this.scopeSelected({ value: this.permission.scope });
    }
    if (this.permission?.access?.owner) {
      this.ownerOn({ checked: this.permission?.access?.owner });
    }
    this.changeDetector.detectChanges();
  }

  private fetchScopes(): void {
    this.scopes$ = this.aclService.getScopes().pipe(
      map((data: string[]) => {
        let sorted = data.sort();
        this.scopes = sorted;
        return sorted;
      }),
      catchError((error: any) => {
        this.scopeFetchError = error;
        this.changeDetector.detectChanges();
        throw error;
      }),
    );
  }

  private setSuggestedGroups(): void {
    const claims: any = this.oauthService.getIdentityClaims();
    this.userEmail = claims?.email;
    const emailArray = this.userEmail?.split('@');
    let domain = '';
    if (emailArray?.length > 1) {
      domain = emailArray[1];
    }
    this.suggestedGroups = [];
    let userGroups = {
      label: this.translateService.instant('pages.acl.create.yourGroups'),
      items: [],
    };
    claims?.groups?.forEach((group: string) => {
      userGroups.items.push({
        label: `${group}@${domain}`,
        value: `${group}@${domain}`,
      });
    });
    this.suggestedGroups.push(userGroups);
  }

  private clustersToSubject(): void {
    this.subscriptions.add(
      this.clustersService
        .getByFilter({
          limit: 100,
          skip: 0,
        })
        .subscribe((clusters: PaginationResult<Cluster>) => {
          let clusterList = {
            label: this.translateService.instant('pages.clusters.title'),
            items: [],
          };
          clusters?.data?.forEach((element) => {
            clusterList.items.push({
              label: element?.clusterName,
              value: element?.clusterId,
            });
          });
          this.suggestedSubjects.push(clusterList);
        }),
    );
  }

  private scopesToSubject(): void {
    let scopesList = {
      label: this.translateService.instant('pages.acl.form.scope'),
      items: [],
    };
    scopesList.items.push({
      label: 'globalscope',
      value: 'globalscope',
    });
    this.scopes?.forEach((scope: string) => {
      scopesList.items.push({
        label: scope,
        value: scope,
      });
    });

    this.suggestedSubjects.push(scopesList);
  }
}

import { ChangeDetectionStrategy, ChangeDetectorRef, Component, EventEmitter, Input, OnDestroy, OnInit, Output } from '@angular/core';
import { FormArray, FormBuilder, FormGroup, Validators } from '@angular/forms';
import { TranslateService } from '@ngx-translate/core';
import { MessageService, ConfirmationService, ConfirmEventType } from 'primeng/api';
import { catchError, Observable, of, Subscription, tap } from 'rxjs';
import { ClusterMetadata } from '../../../core/models/clusterMetadata';
import { PaginationResult } from '../../../core/models/paginatedResult';
import { Project, ProjectRole } from '../../../core/models/project';
import { ClustersService } from '../../../core/services/clusters.service';
import { ConfigService } from '../../../core/services/config.service';
import { ProjectService } from '../../../core/services/project.service';

@Component({
  selector: 'app-cluster-details-edit',
  templateUrl: './cluster-details-edit.component.html',
  styleUrls: ['./cluster-details-edit.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ClusterDetailsEditComponent implements OnInit, OnDestroy {
  @Input() cluster: any | undefined;
  @Output() invalidCount = new EventEmitter<number>();
  @Output() updateOk = new EventEmitter<boolean>();

  clusterModelForm: FormGroup;
  environment: any;

  clusterUpdateError: any;
  projects$: Observable<PaginationResult<Project>> = undefined;
  projectsFetchError: any;
  projects: PaginationResult<Project>;
  projectNameSearch: string;
  selectedProject: Project;

  availableCriticalities: { name: string; value: number }[];
  selectedCriticality: { name: string; value: number };
  availableSensitivities: { name: string; value: number }[];
  selectedSensitivity: { name: string; value: number };

  availableRoles: any[];

  private subscriptions = new Subscription();
  private tags: string[] = [];
  private fill = false;
  private rortextregex = this.configService.config.regex.forms;

  constructor(
    private changeDetector: ChangeDetectorRef,
    private fb: FormBuilder,
    private clustersService: ClustersService,
    private projectService: ProjectService,
    private translateService: TranslateService,
    private messageService: MessageService,
    private confirmationService: ConfirmationService,
    private configService: ConfigService,
  ) {}

  ngOnInit(): void {
    this.setAvailableRoles();
    this.setupObservables();
    this.setupAvailableCriticalitiesAndSensitivities();
    this.setupForm();
    this.invalidCount.emit(this.getInvalidCount(this.clusterModelForm));
  }

  ngOnDestroy(): void {
    this.subscriptions.unsubscribe();
  }

  private getInvalidCount(container: FormGroup): number {
    let invalidCount = 0;
    for (let controlKey in container.controls) {
      if (container.controls.hasOwnProperty(controlKey)) {
        if (container.controls[controlKey].status === 'INVALID') {
          invalidCount = invalidCount + 1;
        }
      }
    }
    return invalidCount;
  }

  setupObservables(): void {
    this.subscriptions.add(
      this.translateService.onLangChange
        .pipe(
          tap(() => {
            this.setupAvailableCriticalitiesAndSensitivities();
          }),
        )
        .subscribe(),
    );
    this.projectsFetchError = undefined;
    this.projects$ = this.projectService
      .getByFilter({
        limit: 100,
        skip: 0,
        filters: [
          {
            field: 'name',
            matchMode: 'contains',
            value: this.projectNameSearch,
          },
          {
            field: 'active',
            matchMode: 'equals',
            value: true,
          },
        ],
      })
      .pipe(
        tap((projects: PaginationResult<Project>) => {
          this.projects = projects;
          this.fill = true;
          this.fillForm();
          this.fill = false;
          this.clusterModelForm.markAsPristine();
          this.invalidCount.emit(this.getInvalidCount(this.clusterModelForm));
          this.changeDetector.detectChanges();
        }),
        catchError((error: any) => {
          this.projectsFetchError = error;
          this.changeDetector.detectChanges();
          throw error;
        }),
      );
  }

  setupForm(): void {
    this.clusterModelForm = this.fb.group({
      description: [null, { validators: [Validators.minLength(1), Validators.pattern(this.rortextregex)] }],
      projectId: [null, { validators: [Validators.required, Validators.minLength(1)] }],
      sensitivity: [null, { validators: [Validators.required, Validators.min(1), Validators.max(4)] }],
      criticality: [null, { validators: [Validators.required, Validators.min(1), Validators.max(4)] }],
      billing: this.fb.group({
        workorder: [null, { validators: [Validators.required, Validators.minLength(1), Validators.pattern(this.rortextregex)] }],
      }),
      tags: ['', { validators: [] }],
      roles: this.fb.array([], { validators: [Validators.required, Validators.minLength(2)] }),
    });
  }

  private setAvailableRoles(): void {
    this.availableRoles = [
      {
        name: this.translateService.instant('pages.admin.projects.form.roles.availableRoles.owner'),
        value: 'Owner',
      },
      {
        name: this.translateService.instant('pages.admin.projects.form.roles.availableRoles.responsible'),
        value: 'Responsible',
      },
    ];
  }

  get roles(): FormArray {
    return this.clusterModelForm.get('roles') as FormArray;
  }

  addRole() {
    const roleForm = this.fb.group({
      roleDefinition: [null, { validators: [Validators.required, Validators.minLength(1)] }],
      contactInfo: this.fb.group({
        upn: [null, { validators: [Validators.required, Validators.email] }],
        email: [null, { validators: [Validators.email] }],
        phone: [null, { validators: [Validators.minLength(1)] }],
      }),
    });

    this.roles.push(roleForm);
    this.changeDetector.detectChanges();
  }

  deleteRole(lessonIndex: number) {
    this.roles.removeAt(lessonIndex);
    this.changeDetector.detectChanges();
  }

  setupAvailableCriticalitiesAndSensitivities(): void {
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

    if (this.selectedCriticality) {
      const filter = this.availableCriticalities.filter((x) => x.value == this.selectedCriticality.value);
      if (filter.length === 1) {
        this.selectedCriticality = filter[0];
      }
    }

    if (this.selectedSensitivity) {
      const filter = this.availableSensitivities.filter((x) => x.value == this.selectedSensitivity.value);
      if (filter.length === 1) {
        this.selectedSensitivity = filter[0];
      }
    }
  }

  reset(): void {
    this.clusterUpdateError = undefined;
    this.clusterModelForm.reset();
    this.fillForm();
    this.invalidCount.emit(this.getInvalidCount(this.clusterModelForm));
    this.changeDetector.detectChanges();
  }

  regretChanges(): void {
    this.clusterUpdateError = undefined;
    this.clusterModelForm.reset();
    this.fillForm();
    this.changeDetector.detectChanges();
  }

  fillForm(): void {
    this.clusterModelForm.patchValue(this.cluster?.metadata);

    const projectId = this.cluster?.metadata?.projectId;
    this.projectIdChanged(projectId);

    const criticality = this.cluster?.metadata?.criticality;
    this.criticalityChanged(criticality);

    const sensitivity = this.cluster?.metadata?.sensitivity;
    this.sensitivityChanged(sensitivity);

    let tags: string[] = [];
    if (this.cluster?.metadata?.serviceTags) {
      const keys = Object.keys(this.cluster?.metadata?.serviceTags);
      keys.forEach((key: string) => {
        tags.push(key);
      });
    }

    this.roles.clear();
    const roles: ProjectRole[] = [];
    if (this.cluster?.metadata?.roles?.length > 0) {
      this.cluster?.metadata?.roles?.forEach((role: ProjectRole) => {
        this.addRole();
        roles.push(role);
      });
    } else {
      this.selectedProject?.projectMetadata?.roles?.forEach((role) => {
        this.addRole();
        roles.push(role);
      });
    }

    this.clusterModelForm.patchValue({ tags: tags, roles: roles });

    this.changeDetector.detectChanges();
    this.invalidCount.emit(this.getInvalidCount(this.clusterModelForm));
  }

  updateCluster(): void {
    this.clusterUpdateError = undefined;
    let updateClusterModel = this.clusterModelForm.value as ClusterMetadata;
    if (!updateClusterModel) {
      return;
    }

    const serviceTags = this.createServiceTagArray();
    updateClusterModel.serviceTags = Object.fromEntries(serviceTags);

    this.subscriptions.add(
      this.clustersService
        .updateMetadata(this.cluster.clusterId, updateClusterModel)
        .pipe(
          tap(() => {
            this.updateOk.emit(true);
            this.messageService.add({ severity: 'success', summary: 'Cluster updated' });
            return of(null);
          }),
          catchError((error) => {
            this.clusterUpdateError = error;
            this.updateOk.emit(false);
            this.messageService.add({ severity: 'error', summary: 'Error updating', detail: 'Error updating cluster metadata' });
            this.changeDetector.detectChanges();
            throw error;
          }),
        )
        .subscribe(),
    );
  }

  projectIdChanged(projectId: string): void {
    if (!projectId) {
      this.selectedProject = undefined;
      return;
    }
    this.projects.data.forEach((project) => {
      if (project.id == projectId) {
        this.selectedProject = project;
        if (!this.fill) {
          if (this.cluster?.metadata?.projectId?.length > 0) {
            this.askToCopyProject();
          } else {
            this.copyFromProject();
          }
        }
      }
    });
    this.invalidCount.emit(this.getInvalidCount(this.clusterModelForm));
    this.changeDetector.detectChanges();
  }

  copyFromProject(): void {
    let tags: string[] = [];
    if (this.selectedProject?.projectMetadata?.serviceTags) {
      const keys = Object.keys(this.selectedProject?.projectMetadata?.serviceTags);
      keys.forEach((key: string) => {
        tags.push(key);
      });
    }
    this.roles.clear();
    const roles: ProjectRole[] = [];
    this.selectedProject?.projectMetadata?.roles?.forEach((role) => {
      this.addRole();
      roles.push(role);
    });
    this.clusterModelForm.patchValue({ tags: tags, billing: { workorder: this.selectedProject?.projectMetadata?.billing?.workorder }, roles: roles });
    this.invalidCount.emit(this.getInvalidCount(this.clusterModelForm));
    this.changeDetector.detectChanges();
  }

  askToCopyProject(): void {
    this.confirmationService.confirm({
      header: this.translateService.instant('pages.clusters.details.edit.form.copyProject.title'),
      message: this.translateService.instant('pages.clusters.details.edit.form.copyProject.details', { projectName: this.selectedProject?.name }),
      accept: () => {
        this.copyFromProject();
      },
      reject: (type: ConfirmEventType) => {
        switch (type) {
          case ConfirmEventType.REJECT:
            this.messageService.add({
              severity: 'warn',
              summary: this.translateService.instant('pages.clusters.details.edit.form.copyProject.reject'),
              detail: this.translateService.instant('pages.clusters.details.edit.form.copyProject.reject'),
            });
            break;
          case ConfirmEventType.CANCEL:
            this.messageService.add({
              severity: 'warn',
              summary: this.translateService.instant('pages.clusters.details.edit.form.copyProject.canceled'),
              detail: this.translateService.instant('pages.clusters.details.edit.form.copyProject.reject'),
            });
            break;
        }
      },
    });
  }

  criticalityChanged(criticality: number): void {
    if (!criticality) {
      this.selectedCriticality = undefined;
      return;
    }
    this.availableCriticalities.forEach((c) => {
      if (c.value == criticality) {
        this.selectedCriticality = c;
        return;
      }
    });
    this.invalidCount.emit(this.getInvalidCount(this.clusterModelForm));
    this.changeDetector.detectChanges();
  }

  sensitivityChanged(sensitivity: number): void {
    if (!sensitivity) {
      this.selectedSensitivity = undefined;
      return;
    }
    this.availableSensitivities.forEach((s) => {
      if (s.value == sensitivity) {
        this.selectedSensitivity = s;
        return;
      }
    });
    this.invalidCount.emit(this.getInvalidCount(this.clusterModelForm));
    this.changeDetector.detectChanges();
  }

  clusterNameChanged($event): void {
    this.invalidCount.emit(this.getInvalidCount(this.clusterModelForm));
  }

  private createServiceTagArray(): Map<string, string> {
    let serviceTags: Map<string, string> = new Map();
    const formServiceTags = this.clusterModelForm.get('tags').getRawValue();
    if (!formServiceTags || formServiceTags?.length == 0) {
      return serviceTags;
    }

    formServiceTags.forEach((tag: string) => {
      serviceTags.set(tag, '');
    });

    return serviceTags;
  }
}

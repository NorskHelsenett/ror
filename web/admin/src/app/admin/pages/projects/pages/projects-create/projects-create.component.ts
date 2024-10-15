import { ChangeDetectionStrategy, ChangeDetectorRef, Component, OnDestroy, OnInit } from '@angular/core';
import { FormArray, FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { LangChangeEvent, TranslateService } from '@ngx-translate/core';
import { catchError, Subscription, tap } from 'rxjs';
import { environment } from '../../../../../../environments/environment';
import { Project } from '../../../../../core/models/project';
import { ConfigService } from '../../../../../core/services/config.service';
import { ProjectService } from '../../../../../core/services/project.service';

@Component({
  selector: 'app-projects-create',
  templateUrl: './projects-create.component.html',
  styleUrls: ['./projects-create.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ProjectsCreateComponent implements OnInit, OnDestroy {
  projectId: string = '';
  projectForm: FormGroup;
  projectModel: Project;
  projectCreateError: any;
  projectUpdateError: any;

  project: Project | undefined;
  projectFetchError: any;
  environment: any;

  availableRoles: any[];

  private submitted: boolean = false;

  private subscriptions = new Subscription();
  private rortextregex = this.configService.config.regex.forms;

  constructor(
    private changeDetector: ChangeDetectorRef,
    private fb: FormBuilder,
    private projectService: ProjectService,
    private router: Router,
    private route: ActivatedRoute,
    private translateService: TranslateService,
    private configService: ConfigService,
  ) {}

  ngOnInit(): void {
    this.setAvailableRoles();
    this.setupForm();
    this.environment = environment;

    this.subscriptions.add(
      this.route.params.subscribe((param) => {
        this.projectId = param?.['id'];
        if (this.projectId !== '' && this.projectId !== null && this.projectId !== undefined) {
          this.fetchProject();
          this.changeDetector.detectChanges();
        }
      }),
    );

    this.subscriptions.add(
      this.translateService.onLangChange
        .pipe(
          tap((event: LangChangeEvent) => {
            event.lang;
            this.setAvailableRoles();
          }),
        )
        .subscribe(),
    );

    this.changeDetector.detectChanges();
  }

  ngOnDestroy(): void {
    this.subscriptions.unsubscribe();
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

  setupForm(): void {
    this.projectForm = this.fb.group({
      name: [null, { validators: [Validators.required, Validators.minLength(1), Validators.pattern(this.rortextregex)] }],
      description: [null, { validators: [Validators.minLength(1), Validators.pattern(this.rortextregex)] }],
      active: [true, { validators: [Validators.required] }],
      projectMetadata: this.fb.group({
        billing: this.fb.group({
          workorder: [null, { validators: [Validators.required, Validators.minLength(1)] }],
        }),
        roles: this.fb.array([], { validators: [Validators.required, Validators.minLength(2)] }),
        tags: [[], { validators: [] }],
      }),
    });
  }

  get roles(): FormArray {
    return this.projectForm.get('projectMetadata').get('roles') as FormArray;
  }

  addRole() {
    const roleForm = this.fb.group({
      roleDefinition: [null, { validators: [Validators.required, Validators.minLength(1)] }],
      contactInfo: this.fb.group({
        upn: [null, { validators: [Validators.required, Validators.email] }],
        email: [null, { validators: [Validators.email] }],
        phone: [null, { validators: [Validators.minLength(1), Validators.pattern(this.rortextregex)] }],
      }),
    });

    this.roles.push(roleForm);
    this.changeDetector.detectChanges();
  }

  deleteRole(lessonIndex: number) {
    this.roles.removeAt(lessonIndex);
    this.changeDetector.detectChanges();
  }

  fetchProject(): void {
    this.subscriptions.add(
      this.projectService
        .getById(this.projectId)
        .pipe(
          tap((project: Project) => {
            this.project = project;
            this.fillForm();
          }),
          catchError((error) => {
            this.projectFetchError = error;
            this.changeDetector.detectChanges();
            throw error;
          }),
        )
        .subscribe(),
    );
  }

  fillForm(): void {
    if (this.roles.length > 0) {
      while (this.roles.length !== 0) {
        this.roles.removeAt(0);
      }
    }

    this.project?.projectMetadata?.roles.forEach((role) => {
      this.addRole();
    });

    let tags: string[] = [];
    if (this.project?.projectMetadata?.serviceTags) {
      const keys = Object.keys(this.project?.projectMetadata?.serviceTags);
      keys.forEach((key: string) => {
        tags.push(key);
      });
    }

    this.projectForm.patchValue(this.project);
    this.projectForm.patchValue({ projectMetadata: { tags: tags } });
    this.changeDetector.detectChanges();
  }

  createProject(): void {
    this.projectCreateError = undefined;
    this.projectModel = this.projectForm.value as Project;
    if (!this.projectForm.valid) {
      this.projectForm.markAllAsTouched();
      this.changeDetector.detectChanges();
      return;
    }

    this.submitted = true;
    let project = this.projectForm.value as Project;
    const serviceTags = this.createServiceTagArray();
    project.projectMetadata.serviceTags = Object.fromEntries(serviceTags);
    this.subscriptions.add(
      this.projectService
        .create(project)
        .pipe(
          tap((project: Project) => {
            if (!project) {
              this.projectCreateError = 'Could not create project';
            } else {
              this.router.navigate(['../'], { relativeTo: this.route });
            }
            this.changeDetector.detectChanges();
          }),
          catchError((error) => {
            this.projectCreateError = error;
            this.changeDetector.detectChanges();
            throw error;
          }),
        )
        .subscribe(),
    );
    this.changeDetector.detectChanges();
  }

  updateProject(): void {
    this.projectUpdateError = undefined;
    this.projectModel = this.projectForm.value as Project;
    if (!this.projectForm.valid) {
      this.projectForm.markAllAsTouched();
      this.changeDetector.detectChanges();
      return;
    }
    this.submitted = true;
    let project = this.projectForm.value as Project;
    project.id = this.project?.id;
    const serviceTags = this.createServiceTagArray();

    project.projectMetadata.serviceTags = Object.fromEntries(serviceTags);
    this.subscriptions.add(
      this.projectService
        .update(project)
        .pipe(
          tap((project: Project) => {
            if (!project) {
              this.projectUpdateError = 'Could not update project';
            } else {
              this.router.navigate(['../../'], { relativeTo: this.route });
            }
            this.changeDetector.detectChanges();
          }),
          catchError((error) => {
            this.projectUpdateError = error;
            this.changeDetector.detectChanges();
            throw error;
          }),
        )
        .subscribe(),
    );
    this.changeDetector.detectChanges();
  }

  reset(): void {
    this.projectCreateError = undefined;
    this.projectUpdateError = undefined;
    this.projectModel = null;
    this.submitted = false;
    this.projectForm.reset();
    this.projectForm.patchValue({
      active: true,
    });
    this.changeDetector.detectChanges();
  }

  regretChanges(): void {
    this.fillForm();
    this.changeDetector.detectChanges();
  }

  private createServiceTagArray(): Map<string, string> {
    let serviceTags: Map<string, string> = new Map();
    const formServiceTags = this.projectForm.get('projectMetadata').get('tags').getRawValue();
    if (!formServiceTags || formServiceTags?.length == 0) {
      return serviceTags;
    }

    formServiceTags.forEach((tag: string) => {
      serviceTags.set(tag, '');
    });

    return serviceTags;
  }
}

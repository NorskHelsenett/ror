import { ChangeDetectorRef, Component, OnDestroy, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { mergeMap, Observable, Subscription, tap } from 'rxjs';
import { WorkspacesService } from '../../../core/services/workspaces.service';
import { Workspace } from '../../../core/models/workspace';
import { Project } from '../../../core/models/project';
import { ProjectService } from '../../../core/services/project.service';

@Component({
  selector: 'app-workspace-new',
  templateUrl: './workspace-new.component.html',
  styleUrls: ['./workspace-new.component.scss'],
})
export class WorkspaceNewComponent implements OnInit, OnDestroy {
  form: FormGroup;
  workspaceName: string;
  workspace: Workspace;
  projects: Project[];

  observables$: Observable<any>;

  private subscriptions = new Subscription();

  constructor(
    private formBuilder: FormBuilder,
    private workspaceService: WorkspacesService,
    private projectService: ProjectService,
    private router: Router,
    private route: ActivatedRoute,
    private changeDetector: ChangeDetectorRef,
  ) {}

  ngOnInit(): void {
    this.setupForm();
    this.setupObservables();

    this.workspaceName = this.route.snapshot.params['workspaceName'];
  }

  ngOnDestroy(): void {
    this.subscriptions.unsubscribe();
  }

  setupForm(): void {
    this.form = this.formBuilder.group({
      project: [null, [Validators.required]],
    });
  }

  setupObservables(): void {
    this.observables$ = this.route.params.pipe(
      tap((data: any) => (this.workspaceName = data?.workspaceName)),
      mergeMap((data: any) =>
        this.projectService
          .getByFilter({
            skip: 0,
            limit: 100,
          })
          .pipe(
            tap((projects) => (this.projects = projects?.data)),
            mergeMap((projects) =>
              this.workspaceService.getByName(data?.workspaceName).pipe(
                tap((workspace) => {
                  this.workspace = workspace;
                  if (workspace.projectId !== '' && workspace.projectId !== undefined) {
                    this.form.patchValue({ project: projects?.data?.find((project) => project?.id == workspace?.projectId) });
                  }
                  this.changeDetector.detectChanges();
                }),
              ),
            ),
          ),
      ),
    );
  }

  update(): void {
    this.workspace.projectId = this.form.get('project').value.id;
    this.subscriptions.add(
      this.workspaceService
        .update(this.workspace)
        .pipe(
          tap(() => {
            this.router.navigate(['../'], { relativeTo: this.route });
          }),
        )
        .subscribe(),
    );
  }
}

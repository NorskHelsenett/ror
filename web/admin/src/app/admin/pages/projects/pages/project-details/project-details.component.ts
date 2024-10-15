import { ChangeDetectionStrategy, ChangeDetectorRef, Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { LangChangeEvent, TranslateService } from '@ngx-translate/core';
import { catchError, Observable, share, Subscription, tap } from 'rxjs';
import { AclScopes, AclAccess } from '../../../../../core/models/acl-scopes';
import { ClusterInfo } from '../../../../../core/models/clusterInfo';
import { Project } from '../../../../../core/models/project';
import { AclService } from '../../../../../core/services/acl.service';
import { ConfigService } from '../../../../../core/services/config.service';
import { ProjectService } from '../../../../../core/services/project.service';

@Component({
  selector: 'app-project-details',
  templateUrl: './project-details.component.html',
  styleUrls: ['./project-details.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ProjectDetailsComponent implements OnInit {
  adminRead$: Observable<boolean> | undefined;
  adminReadFetchError: any;
  project$: Observable<Project> = undefined;

  projectId: string;
  tags: string[] = [];
  projectFetchError: any;

  clusterInfos$: Observable<ClusterInfo[]> = undefined;
  clusterInfoFetchError: any;

  rowsPerPage = this.configService.config.rowsPerPage;

  private subscriptions = new Subscription();

  constructor(
    private changeDetector: ChangeDetectorRef,
    private route: ActivatedRoute,
    private projectService: ProjectService,
    private translateService: TranslateService,
    private aclService: AclService,
    private configService: ConfigService,
  ) {}

  ngOnInit(): void {
    this.subscriptions.add(
      this.route.params.subscribe((param) => {
        this.projectId = param?.['id'];
        if (this.projectId !== '' && this.projectId !== null && this.projectId !== undefined) {
          this.fetchProject();
          this.fetchClusterInfos();
          this.changeDetector.detectChanges();
        }
      }),
    );

    this.subscriptions.add(
      this.translateService.onLangChange
        .pipe(
          tap((event: LangChangeEvent) => {
            event.lang;
          }),
        )
        .subscribe(),
    );

    this.fetchAcl();
    this.changeDetector.detectChanges();
  }

  fetchProject(): void {
    this.project$ = this.projectService.getById(this.projectId).pipe(
      share(),
      tap((project: Project) => {
        if (!project) {
          return;
        }
        this.tags = [];
        let tags: string[] = [];
        if (project?.projectMetadata?.serviceTags) {
          const keys = Object.keys(project?.projectMetadata?.serviceTags);
          keys.forEach((key: string) => {
            tags.push(key);
          });
        }
        this.tags = tags;
        this.changeDetector.detectChanges();
      }),
      catchError((err) => {
        this.projectFetchError = err;
        this.changeDetector.detectChanges();
        throw err;
      }),
    );
  }

  fetchClusterInfos(): void {
    this.clusterInfoFetchError = undefined;
    this.clusterInfos$ = this.projectService.clustersByProjectId(this.projectId).pipe(
      catchError((error) => {
        this.clusterInfoFetchError = error;
        this.changeDetector.detectChanges();
        throw error;
      }),
    );
  }

  fetchAcl(): void {
    this.adminReadFetchError = undefined;
    this.adminRead$ = this.aclService.check(AclScopes.ROR, AclScopes.Global, AclAccess.Read).pipe(
      share(),
      catchError((error: any) => {
        this.changeDetector.detectChanges();
        throw error;
      }),
    );
  }
}

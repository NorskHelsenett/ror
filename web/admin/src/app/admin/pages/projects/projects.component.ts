import { TranslateService } from '@ngx-translate/core';
import { ChangeDetectionStrategy, ChangeDetectorRef, Component, OnInit } from '@angular/core';
import { ConfirmationService, MessageService } from 'primeng/api';
import { catchError, finalize, Observable, share, Subscription } from 'rxjs';
import { FilterService } from '../../../core/services/filter.service';
import { AclScopes, AclAccess } from '../../../core/models/acl-scopes';
import { Filter } from '../../../core/models/apiFilter';
import { PaginationResult } from '../../../core/models/paginatedResult';
import { Project, ProjectRole, RoleDefinition } from '../../../core/models/project';
import { AclService } from '../../../core/services/acl.service';
import { ConfigService } from '../../../core/services/config.service';
import { ExportService } from '../../../core/services/export.service';
import { ProjectService } from '../../../core/services/project.service';

@Component({
  selector: 'app-projects',
  templateUrl: './projects.component.html',
  styleUrls: ['./projects.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ProjectsComponent implements OnInit {
  projects$: Observable<PaginationResult<Project>> | undefined;
  projectsError: any;
  columns: any[] = [];
  rows = this.configService.config.rows;
  rowsPerPage = this.configService.config.rowsPerPage;
  totalRecords: number = 0;
  loading: boolean;
  filter: Filter;
  lastFilter: Filter;
  lastLazyLoad: any;
  showExportChoises: boolean;

  adminRead$: Observable<boolean> | undefined;
  adminReadFetchError: any;
  accessTypes: any[];

  private subscriptions: Subscription = new Subscription();

  constructor(
    private changeDetector: ChangeDetectorRef,
    private projectService: ProjectService,
    private filterService: FilterService,
    private aclService: AclService,
    private confirmationService: ConfirmationService,
    private messageService: MessageService,
    private translateService: TranslateService,
    private exportService: ExportService,
    private configService: ConfigService,
  ) {}

  ngOnInit(): void {
    this.columns = [
      {
        field: 'name',
        header: 'name',
      },
      {
        field: 'workorder',
        header: 'billing.workorder',
      },
      {
        field: 'updated',
        header: 'updated',
      },
      {
        field: 'created',
        header: 'created',
      },
      {
        field: 'active',
        header: 'active',
      },
    ];

    this.fetchAcl();
    this.setupTypes();
  }

  ngOnDestroy(): void {
    this.subscriptions.unsubscribe();
  }

  fetchAcl(): void {
    this.adminRead$ = this.aclService?.check(AclScopes.ROR, AclScopes.Global, AclAccess.Read).pipe(
      share(),
      catchError((error: any) => {
        this.adminReadFetchError = error;
        this.changeDetector.detectChanges();
        throw error;
      }),
    );
  }

  fetchProjects(event: any): void {
    if (event) {
      this.filter = this.filterService.mapFilter(event);
    }
    this.loading = true;
    this.lastFilter = this.filter;
    this.projectsError = undefined;
    this.projects$ = this.projectService.getByFilter(this.filter).pipe(
      share(),
      catchError((error: any) => {
        this.projectsError = error;
        throw error;
      }),
      finalize(() => {
        this.loading = false;
        this.changeDetector.detectChanges();
      }),
    );
  }

  getValueFromColumn(project: Project, column: string): any {
    let nestedColumns: string[] = column.split('.');
    let value: any = project[nestedColumns.shift()];
    nestedColumns.forEach((col) => {
      value = value[col];
    });
    return value;
  }

  deleteProject(project: Project): void {
    this.confirmationService.confirm({
      header: this.translateService.instant('pages.admin.projects.delete.title'),
      message: this.translateService.instant('pages.admin.projects.delete.details', { projectName: project?.name }),
      accept: () => {
        this.subscriptions.add(
          this.projectService
            .delete(project)
            .pipe(
              catchError((error) => {
                this.messageService.add({
                  severity: 'error',
                  summary: this.translateService.instant('pages.admin.projects.delete.error.title'),
                  detail: this.translateService.instant('pages.admin.projects.delete.error.details'),
                });

                throw error;
              }),
              finalize(() => {
                this.fetchProjects(this.lastLazyLoad);
                this.changeDetector.detectChanges();
              }),
            )
            .subscribe(() => {
              this.messageService.add({ severity: 'success', summary: this.translateService.instant('pages.admin.projects.delete.success.title') });
              this.fetchProjects(this.lastLazyLoad);
              this.changeDetector.detectChanges();
            }),
        );
      },
    });
  }

  exportToExcel(): void {
    this.exportData('excel');
  }

  exportToCsv(): void {
    this.exportData('csv');
  }

  private exportData(type: string): void {
    this.subscriptions.add(
      this.projectService.getByFilter(this.filter).subscribe((clustersPaginated: PaginationResult<Project>) => {
        const clusters = this.exportProjects(clustersPaginated?.data);
        if (type === 'csv') {
          this.exportService.exportToCsv(clusters, 'ror-projects.csv');
        }
        if (type === 'excel') {
          this.exportService.exportAsExcelFile(clusters, 'ror-projects.xlsx');
        }
      }),
    );
  }

  private exportProjects(projects: any): any[] {
    let exportProject: any[] = [];
    projects.forEach((p) => {
      let tags: string[] = [];
      if (p?.projectMetadata?.serviceTags) {
        const keys = Object.keys(p?.projectMetadata?.serviceTags);
        keys.forEach((key: string) => {
          tags.push(key);
        });
      }

      let project: any = {
        name: p.name,
        active: p.active,
        created: p.created,
        updated: p.updated,
        description: p.description,
        workorder: p.projectMetadata.billing.workorder,
        roles: JSON.stringify(p.projectMetadata.roles),
        tags: tags.join(' '),
      };

      let owner: ProjectRole = p?.projectMetadata?.roles?.find((role: ProjectRole) => role?.roleDefinition === RoleDefinition.Owner);
      project['ownerEmail'] = owner?.contactInfo?.email;
      project['ownerPhone'] = owner?.contactInfo?.phone;

      let responsible: ProjectRole = p?.projectMetadata?.roles?.find((role: ProjectRole) => role?.roleDefinition === RoleDefinition.Responsible);
      project['responsibleEmail'] = responsible?.contactInfo?.email;
      project['responsiblePhone'] = responsible?.contactInfo?.phone;

      exportProject.push(project);
    });

    return exportProject;
  }

  setupTypes(): void {
    this.accessTypes = [
      {
        name: this.translateService.instant('shared.trueOrFalse.true'),
        value: true,
      },
      {
        name: this.translateService.instant('shared.trueOrFalse.false'),
        value: false,
      },
    ];
  }
}

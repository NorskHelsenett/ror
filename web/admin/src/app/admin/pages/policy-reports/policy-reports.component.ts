import { ChangeDetectorRef, Component, OnInit } from '@angular/core';
import { TranslateService } from '@ngx-translate/core';
import { MessageService } from 'primeng/api';
import { Observable, catchError } from 'rxjs';
import { PolicyReportGlobal, PolicyReportGlobalQueryType } from '../../../core/models/policyReport';
import { PolicyReportsService } from '../../../core/services/policy-reports.service';

@Component({
  selector: 'app-policy-reports',
  templateUrl: './policy-reports.component.html',
  styleUrls: ['./policy-reports.component.scss'],
})
export class PolicyReportsComponent implements OnInit {
  policyReportsPerCluster$: Observable<PolicyReportGlobal[]>;
  error: any;
  cols: any[];

  constructor(
    private changeDetector: ChangeDetectorRef,
    private translateService: TranslateService,
    private messageService: MessageService,
    private policyReportsService: PolicyReportsService,
  ) {}

  ngOnInit(): void {
    this.setupColumns();
    this.getPolicyReportsPerCluster();
  }

  getPolicyReportsPerCluster(): void {
    this.policyReportsPerCluster$ = this.policyReportsService
      .getPolicyReportsGlobal(PolicyReportGlobalQueryType.PolicyReportGlobalQueryTypeCluster, '')
      .pipe(
        catchError((error) => {
          this.error = error;
          this.messageService.add({
            severity: 'error',
            summary: this.translateService.instant('pages.admin.policyreports.error'),
          });
          this.changeDetector.detectChanges();
          throw error;
        }),
      );
  }

  setupColumns(): void {
    this.cols = [
      {
        field: 'cluster',
        header: 'cluster',
      },
      {
        field: 'fail',
        header: 'fail',
      },
      {
        field: 'pass',
        header: 'pass',
      },
    ];
  }

  formatExport(policyReports: PolicyReportGlobal[]): any[] {
    const exportObjects: any[] = [];
    policyReports?.forEach((report) => {
      exportObjects.push({
        cluster: report?.cluster,
        policy: report?.policy,
        failed: report?.fail,
        passed: report?.pass,
      });
    });
    return exportObjects;
  }
}

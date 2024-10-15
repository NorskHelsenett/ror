import { ChangeDetectorRef, Component, Input, OnInit } from '@angular/core';
import { TranslateService } from '@ngx-translate/core';
import { MessageService } from 'primeng/api';
import { Observable, catchError } from 'rxjs';
import { PolicyReportGlobal, PolicyReportGlobalQueryType } from '../../../../core/models/policyReport';
import { PolicyReportsService } from '../../../../core/services/policy-reports.service';

@Component({
  selector: 'app-per-policy-table',
  templateUrl: './per-policy-table.component.html',
  styleUrls: ['./per-policy-table.component.scss'],
})
export class PerPolicyTableComponent implements OnInit {
  @Input() clusterID: string;

  policyReports$: Observable<PolicyReportGlobal[]>;
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
    this.getPolicyReports();
  }

  getPolicyReports(): void {
    this.policyReports$ = this.policyReportsService
      .getPolicyReportsGlobal(PolicyReportGlobalQueryType.PolicyReportGlobalQueryTypePolicy, this.clusterID)
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
        field: 'policy',
        header: 'policy',
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
}

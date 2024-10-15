import { ChangeDetectionStrategy, ChangeDetectorRef, Component, OnDestroy, OnInit, QueryList, ViewChildren } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Table } from 'primeng/table';
import { catchError, map, Observable, Subscription, tap } from 'rxjs';
import { PolicyReportView } from '../../../core/models/policyReport';
import { ClustersService } from '../../../core/services/clusters.service';
import { ConfigService } from '../../../core/services/config.service';
import { ExportService } from '../../../core/services/export.service';

@Component({
  selector: 'app-cluster-policy-report',
  templateUrl: './cluster-policy-report.component.html',
  styleUrls: ['./cluster-policy-report.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ClusterPolicyReportComponent implements OnInit, OnDestroy {
  @ViewChildren('resultTable')
  tables: QueryList<Table>;

  policyreports$: Observable<PolicyReportView> | undefined;
  policyreports: PolicyReportView;
  policyreportsError: any;
  clusterId: string;
  showExportChoices: boolean;

  rowsPerPage = this.configService.config.rowsPerPage;

  resultFilter: string[] = ['fail', 'error', 'pass', 'warn', 'skip'];
  resultFilterValue: string[];

  private subscriptions = new Subscription();

  constructor(
    private changeDetector: ChangeDetectorRef,
    private route: ActivatedRoute,
    private clustersService: ClustersService,
    private configService: ConfigService,
    private exportService: ExportService,
  ) {}

  ngOnInit() {
    this.subscriptions.add(
      this.route.params
        .pipe(
          tap((data: any) => {
            this.clusterId = data?.id;
            this.getPolicyReports();
          }),
        )
        .subscribe(() => {
          this.changeDetector.detectChanges();
        }),
    );
    this.triggerFilter();
    this.changeDetector.detectChanges();

    this.resultFilterValue = ['fail'];
  }

  ngOnDestroy() {
    this.subscriptions.unsubscribe();
  }

  getPolicyReports() {
    this.policyreports = undefined;
    this.policyreportsError = undefined;
    this.policyreports$ = this.clustersService.getPolicyreports(this.clusterId).pipe(
      map((data: PolicyReportView) => {
        this.policyreports = data;
        this.changeDetector.detectChanges();
        return this.policyreports;
      }),
      catchError((error) => {
        this.changeDetector.detectChanges();
        this.policyreportsError = error;
        throw error;
      }),
    );
  }

  triggerFilter(): void {
    this.tables.forEach((table: Table) => {
      table.filter(['fail', 'error'], 'result', 'in');
    });
  }

  exportToCsv(reports: PolicyReportView): void {
    this.exportService.exportToCsv(this.reportsToArray(reports), 'policy-reports.csv');
  }

  exportToExcel(reports: PolicyReportView): void {
    this.exportService.exportAsExcelFile(this.reportsToArray(reports), 'policy-reports');
  }

  reportsToArray(reports: PolicyReportView): any[] {
    const reportArray: any[] = [];
    reports?.namespaces?.forEach((namespace) => {
      namespace?.policies?.forEach((policy: any) => {
        policy?.reports?.forEach((report: any) => {
          const entry: any = {};
          entry.cluster = reports?.clusterid;
          entry.namespace = namespace?.name;
          entry.name = policy?.name;
          entry.category = report?.category;
          entry.apiVersion = report?.apiversion;
          entry.kind = report?.kind;
          entry.resourceName = report?.name;
          entry.result = report?.result;
          entry.message = report?.message;
          reportArray?.push(entry);
        });
      });
    });
    return reportArray?.length > 0 ? reportArray : [''];
  }
}

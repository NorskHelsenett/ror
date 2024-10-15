import { Component, OnInit } from '@angular/core';
import { ComplianceReportsService } from '../../../core/services/compliance-reports.service';
import { ComplianceReport, ComplianceReportGlobal, ComplianceReportReport, ComplianceReportSummary } from '../../../core/models/complianceReport';
import { Observable, catchError, concatMap, groupBy, map, mergeAll, mergeMap, of, tap, toArray } from 'rxjs';

@Component({
  selector: 'app-compliance-reports',
  templateUrl: './compliance-reports.component.html',
  styleUrl: './compliance-reports.component.scss',
})
export class ComplianceReportsComponent implements OnInit {
  complianceReports$: Observable<ComplianceReportGlobal[]>;
  loading: boolean = false;
  error: any = null;
  columns: any[] = [];

  constructor(private complianceReportsService: ComplianceReportsService) {}

  ngOnInit(): void {
    this.setupColumns();
    this.getComplianceReports();
  }

  getComplianceReports(): void {
    this.loading = true;
    this.error = null;
    this.complianceReports$ = this.complianceReportsService.getComplianceReportsGlobal().pipe(
      map((reports) => {
        const complianceReports: ComplianceReportGlobal[] = [];
        reports.forEach((report) => {
          let index: number = complianceReports?.findIndex((r) => r?.clusterid === report?.clusterid);
          if (index === -1) {
            if (report?.metadata?.name === 'cis') {
              complianceReports.push({ clusterid: report?.clusterid, cis: report?.summary });
            } else {
              complianceReports.push({ clusterid: report?.clusterid, nsa: report?.summary });
            }
          } else {
            if (report?.metadata?.name === 'cis') {
              complianceReports[index].cis = report?.summary;
            } else {
              complianceReports[index].nsa = report?.summary;
            }
          }
        });
        this.loading = false;
        this.error = null;
        return complianceReports;
      }),
      catchError((error) => {
        this.loading = false;
        this.error = error;
        return of(error);
      }),
    );
  }

  setupColumns(): void {
    this.columns = [
      {
        field: 'clusterid',
        header: 'cluster',
      },
      {
        field: 'cis.passcount',
        header: 'passcount',
      },
      {
        field: 'cis.failcount',
        header: 'failcount',
      },
      {
        field: 'nsa.passcount',
        header: 'passcount',
      },
      {
        field: 'nsa.failcount',
        header: 'failcount',
      },
    ];
  }

  getValue(field: string, complianceReport: ComplianceReport): string {
    return field?.split('.')?.reduce((o, i) => o[i], complianceReport);
  }

  formatExport(complianceReports: ComplianceReportGlobal[]): any[] {
    const exportObjects: any[] = [];
    complianceReports?.forEach((report) => {
      exportObjects.push({
        cluster: report?.clusterid,
        cisPassed: report?.cis?.passcount,
        cisFailed: report?.cis?.failcount,
        nsaPassed: report?.nsa?.passcount,
        nsaFailed: report?.nsa?.failcount,
      });
    });
    return exportObjects;
  }
}

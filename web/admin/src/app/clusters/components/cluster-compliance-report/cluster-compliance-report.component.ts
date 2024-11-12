import { Component, Input, OnInit } from '@angular/core';
import { Observable, catchError, of } from 'rxjs';
import { ComplianceReport } from '../../../core/models/complianceReport';
import { ComplianceReportsService } from '../../../core/services/compliance-reports.service';
import { ExportService } from '../../../core/services/export.service';

@Component({
  selector: 'app-cluster-compliance-report',
  templateUrl: './cluster-compliance-report.component.html',
  styleUrl: './cluster-compliance-report.component.scss',
})
export class ClusterComplianceReportComponent implements OnInit {
  @Input() clusterId: string;

  complianceReports$: Observable<ComplianceReport[]>;
  mainColumns: any[];
  subColumns: any[];
  showExportChoices: boolean;

  constructor(
    private complianceReportService: ComplianceReportsService,
    private exportService: ExportService,
  ) {}

  ngOnInit(): void {
    this.setupColumns();
    this.getComplianceReports();
  }

  getComplianceReports(): void {
    this.complianceReports$ = this.complianceReportService.getComplianceReports(this.clusterId).pipe(
      catchError((error) => {
        throw of(error);
      }),
    );
  }

  setupColumns(): void {
    this.mainColumns = [
      {
        field: 'metadata.name',
        header: 'name',
      },
      {
        field: 'metadata.title',
        header: 'title',
      },
      {
        field: 'summary.passcount',
        header: 'passcount',
      },
      {
        field: 'summary.failcount',
        header: 'failcount',
      },
    ];
    this.subColumns = [
      {
        field: 'id',
        header: 'id',
      },
      {
        field: 'name',
        header: 'name',
      },
      {
        field: 'severity',
        header: 'severity',
      },
      {
        field: 'totalfail',
        header: 'totalfail',
      },
    ];
  }

  getValue(object: Object, key: string): void {
    const keySplit = key?.split('.', 2);
    if (keySplit?.length > 1) {
      return this?.getValue(object[keySplit[0]], keySplit[1]);
    }
    return object[key];
  }

  exportToCsv(reports: ComplianceReport[]): void {
    this.exportService.exportToCsv(this.formatExport(reports), 'compliance-reports.csv');
  }

  exportToExcel(reports: ComplianceReport[]): void {
    this.exportService.exportAsExcelFile(this.formatExport(reports), 'compliance-reports');
  }

  formatExport(reports: ComplianceReport[]): any[] {
    const formattedExport: any[] = [];
    reports?.forEach((report) => {
      const entry: any = {};
      entry.date = new Date()?.toLocaleString('no');
      entry.name = report?.metadata?.name;
      entry.title = report?.metadata?.title;
      entry.passed = report?.summary?.passcount;
      entry.failed = report?.summary?.failcount;
      formattedExport.push(entry);
    });
    return formattedExport?.length > 0 ? formattedExport : [''];
  }
}

export interface ComplianceReport {
  clusterid: string;
  metadata: ComplianceReportMetadata;
  summary: ComplianceReportSummary;
  reports: ComplianceReportReport[];
}

export interface ComplianceReportMetadata {
  name: string;
  title: string;
}

export interface ComplianceReportSummary {
  failcount: number;
  passcount: number;
}

export interface ComplianceReportReport {
  id: string;
  name: string;
  severity: ComplianceReportSeverity;
  totalfail: number;
}

enum ComplianceReportSeverity {
  CRITICAL = 'CRITICAL',
  HIGH = 'HIGH',
  MEDIUM = 'MEDIUM',
  LOW = 'LOW',
  UNKNOWN = 'UNKOWN',
}

export interface ComplianceReportGlobal {
  clusterid: string;
  cis?: ComplianceReportSummary;
  nsa?: ComplianceReportSummary;
}

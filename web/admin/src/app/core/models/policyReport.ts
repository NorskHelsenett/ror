export interface PolicyReportView extends PolicyReportSummary {
  clusterid: string;
  namespaces: PolicyReportNamespace[];
}

export interface PolicyReportNamespace extends PolicyReportSummary {
  name: string;
  policies: PolicyReport[];
}

export interface PolicyReportSummary {
  failed: number;
  passed: number;
  error: number;
  warning: number;
  skipped: number;
  total: number;
}

export interface PolicyReport {
  uid: string;
  name: string;
  apiversion: string;
  kind: string;
  result: string;
  category: string;
  message: string;
}

export enum PolicyReportGlobalQueryType {
  PolicyReportGlobalQueryTypeUnknown = 'Unknown',
  PolicyReportGlobalQueryTypeCluster = 'Cluster',
  PolicyReportGlobalQueryTypePolicy = 'Policy',
}

export interface PolicyReportGlobal {
  cluster: string;
  policy: string;
  fail: number;
  pass: number;
}

export interface PolicyReportGlobalQuery {
  type: PolicyReportGlobalQueryType;
  internal: boolean;
}

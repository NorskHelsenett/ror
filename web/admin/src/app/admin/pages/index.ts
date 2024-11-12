import { AdminDatacentersComponent } from './datacenters/admin-datacenters.component';
import { AdminPricesComponent } from './prices/admin-prices.component';
import { AdminAuditlogsComponent } from './auditlogs/admin-auditlogs.component';
import { ProjectsComponent } from './projects/projects.component';
import { PolicyReportsComponent } from './policy-reports/policy-reports.component';
import { PerPolicyTableComponent } from './policy-reports/per-policy-table/per-policy-table.component';
import { VulnerabilityReportsComponent } from './vulnerability-reports/vulnerability-reports.component';
import { adminVulnerabilityReportsPages } from './vulnerability-reports/pages';
import { ComplianceReportsComponent } from './compliance-reports/compliance-reports.component';

export * from './datacenters/admin-datacenters.component';
export * from './prices/admin-prices.component';
export * from './auditlogs/admin-auditlogs.component';
export * from './projects/projects.component';
export * from './policy-reports/policy-reports.component';
export * from './policy-reports/per-policy-table/per-policy-table.component';
export * from './vulnerability-reports/vulnerability-reports.component';
export * from './vulnerability-reports/pages';
export * from './compliance-reports/compliance-reports.component';

export const adminPages = [
  AdminDatacentersComponent,
  AdminPricesComponent,
  AdminAuditlogsComponent,
  ProjectsComponent,
  PolicyReportsComponent,
  PerPolicyTableComponent,
  VulnerabilityReportsComponent,
  ...adminVulnerabilityReportsPages,
  ComplianceReportsComponent,
];

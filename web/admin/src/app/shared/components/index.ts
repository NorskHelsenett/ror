import { ClusterEnvironmentComponent } from './cluster-environment/cluster-environment.component';
import { ClusterStatusComponent } from './cluster-status/cluster-status.component';
import { TrueFalseComponent } from './true-false/true-false.component';
import { StatusComponent } from './status/status.component';
import { SpinnerComponent } from './spinner/spinner.component';
import { ExportComponent } from './export-button/export-button.component';

export * from './cluster-environment/cluster-environment.component';
export * from './cluster-status/cluster-status.component';
export * from './true-false/true-false.component';
export * from './status/status.component';
export * from './spinner/spinner.component';
export * from './export-button/export-button.component';

export const sharedComponents = [
  ClusterEnvironmentComponent,
  ClusterStatusComponent,
  TrueFalseComponent,
  StatusComponent,
  SpinnerComponent,
  ExportComponent,
];

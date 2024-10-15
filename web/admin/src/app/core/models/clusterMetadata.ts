import { Billing } from './billing';
import { Project, ProjectRole } from './project';

export interface ClusterMetadata {
  name: string;
  criticality: number;
  sensitivity: number;
  projectId: string;
  project: Project;
  billing: Billing;
  serviceTags: any;
  roles: ProjectRole[];
}

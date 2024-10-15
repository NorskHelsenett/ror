import { Billing } from './billing';

export interface Project {
  id: string;
  name: string;
  description: string;
  active: boolean;
  created: Date;
  updated: Date;
  projectMetadata: ProjectMetadata;
}

export interface ProjectMetadata {
  roles: ProjectRole[];
  billing: Billing;
  serviceTags: any;
}

export interface ProjectRole {
  roleDefinition: RoleDefinition;
  contactInfo: ProjectContactInfo;
}

export interface ProjectContactInfo {
  upn: string;
  email: string;
  phone: string;
}

export enum RoleDefinition {
  Unknown = 'Unknown',
  Owner = 'Owner',
  Responsible = 'Responsible',
}

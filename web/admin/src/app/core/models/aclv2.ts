export interface AclV2 {
  id: string;
  access: Access;
  group: string;
  kubernetes: KubernetesAccess;
  scope: string;
  subject: string;
  version: number;
  created: Date;
  issuedBy: string;
}

export interface Access {
  read: boolean;
  create: boolean;
  update: boolean;
  delete: boolean;
  owner: boolean;
}

export interface KubernetesAccess {
  logon: boolean;
}

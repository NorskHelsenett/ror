export enum AclScopes {
  Unknown = '',

  Global = 'globalscope',
  Acl = 'acl',

  ROR = 'ror',
  DataCenter = 'datacenter',
  Cluster = 'cluster',
  Project = 'project',
  Workspace = 'workspace',
  Price = 'price',
}

export enum AclAccess {
  Unknown = '',
  Read = 'read',
  Create = 'create',
  Update = 'update',
  Delete = 'delete',
  Owner = 'owner',
}

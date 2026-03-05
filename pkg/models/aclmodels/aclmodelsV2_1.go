package aclmodels

type AclV3ListItemAccess map[string]bool

// system:[subsystem]:verb/access, e.g. ror:read, kubernetes:logon, ror:metadata:write

// ror:read: true
// ror:write: true
// ror:owner: true
// kubernetes:logon: true
// ror:metadata:write true

//VALKEY:

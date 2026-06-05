# Authentication & ACL Flow

> Visual diagram: [auth-acl-flow.excalidraw](auth-acl-flow.excalidraw)

## Overview

All HTTP requests go through the same authentication pipeline. The ACL
check diverges into V2 (current) and V3 (new) paths at the controller layer.

## Shared Pipeline (all requests)

```
HTTP Request (Bearer JWT)
  │
  ▼
OAuth Middleware (oauthmiddleware)
  • Validates JWT via oidchelper.MultiIssuerValidator
  • Extracts groups + email from claims
  • Builds Identity{Type, User{Groups, Email}}
  • Stores in Gin context via context.WithValue()
  │
  ▼
Gin Context
  • rorcontext.GetIdentityFromRorContext(ctx) → Identity
  • identity.User.Groups → []string (for users)
  • identity.IsCluster() / identity.GetId() (for clusters)
  │
  ▼
Controller Handler
  • Determines scope + subject for the check
  • Checks identity type (cluster vs user vs service)
```

### Cluster Identity Special Case

Cluster identities bypass the ACL store entirely. They get hardcoded access
to their own resources:

| Field  | Value |
| ------ | ----- |
| Read   | ✓     |
| Create | ✓     |
| Update | ✓     |
| Delete | ✗     |
| Owner  | ✗     |

Condition: `identity.IsCluster() && scope == cluster && subject == identity.GetId()`

## V2 Path (Current — ror-api)

```
Controller
  │
  ▼
aclservice.CheckAccessByContextAclQuery(ctx, query)
  │
  ▼
aclrepository.CheckAcl2ByIdentityQuery(ctx, query)
  • Extracts identity from context
  • Builds MongoDB aggregation pipeline:
    $match {group: {$in: groups}, scope, subject, version: 2}
  • Executes per-request aggregation
  │
  ▼
compileAccess()
  • Boolean OR across all matching ACL entries
  → AclV2ListItemAccess {Read, Create, Update, Delete, Owner}
  │
  ▼
Access Decision (403 or continue)
```

### V2 Ownerref Scoping (resource list queries)

```
aclservice.GetOwnerrefByContextAccess(ctx, AccessTypeRead)
  │
  ▼
aclrepository.GetOwnerrefsQueryAcl2ByIdentityAccess()
  • Loads ALL ACL entries for user's groups (version:2)
  • compileOwnerrefs() builds bson.M $match:
    - scope=ror, subject=globalscope → {} (unrestricted)
    - scope=ror, subject=cluster → {rormeta.ownerref.scope: "cluster"}
    - specific scope+subject → {rormeta.ownerref: {$in: [...]}}
  │
  ▼
Appended as pipeline stage to resource query
```

## V3 Path (New — pkg/acl)

```
Controller
  │
  ▼
acl.Resolver.ResolveAccess(ctx, groups, scope, subject)
  │
  ▼
acl.Store.GetByGroups(ctx, groups)
  • Single batch load for all groups
  • MongoDB Find: {version: {$in: [2, 3]}, group: {$in: groups}}
  • V2 entries auto-converted to V3 via aclmodels.V2ToV3()
  • Optional: CachedStore with Redis MGET layer
  │
  ▼
In-memory resolution
  • matchesScopeSubject() for each entry
  • Set union of matching AccessTypeV3 capabilities
  → []AccessTypeV3 (e.g. ["ror:read", "kubernetes:logon"])
  │
  ▼
Access Decision (check slices.Contains)
```

### V3 Ownerref Scoping

```
acl.Resolver.ResolveOwnerrefs(ctx, groups, requiredAccess)
  • Returns nil (unrestricted) or []Ownerref
  │
  ▼
ScopeExpander.ExpandScope() [optional]
  • BFS walk: Project → Workspace → Cluster
  • CachedScopeExpander for in-memory TTL cache
  │
  ▼
aclstore.OwnerrefsToFilter(refs) → bson.M
  • nil → {} (unrestricted)
  • [] → DenyAllFilter (matches nothing)
  • scope=ror → scope-level grants
  • specific → $in query
  │
  ▼
Appended as pipeline stage to resource query
```

## Package Map

| Package      | Location                         | Purpose                                                        |
| ------------ | -------------------------------- | -------------------------------------------------------------- |
| `acl`        | `pkg/acl/`                       | Resolver, Store interface, ScopeExpander, Ownerref type        |
| `aclstore`   | `pkg/acl/aclstore/`              | MongoStore, CachedStore, MongoScopeExpander, OwnerrefsToFilter |
| `aclmodels`  | `pkg/models/aclmodels/`          | V2 + V3 types, V2↔V3 converters                                |
| `aclscope`   | `pkg/models/aclmodels/aclscope/` | Scope + Subject types, shared between V2 and V3                |
| `rorcontext` | `pkg/context/rorcontext/`        | GetIdentityFromRorContext()                                    |
| `identity`   | `pkg/models/identity/`           | Identity type, Groups, IsCluster(), GetId()                    |

## V2 ↔ V3 Conversion

Driven by `v3` struct tags on `AclV2ListItemAccess`:

| V2 Boolean         | V3 Capability      |
| ------------------ | ------------------ |
| `Read`             | `ror:read`         |
| `Create`           | `ror:create`       |
| `Update`           | `ror:update`       |
| `Delete`           | `ror:delete`       |
| `Owner`            | `ror:owner`        |
| `Kubernetes.Logon` | `kubernetes:logon` |

V3-only capabilities (e.g. `kubernetes:admin`, `resource:Deployment:read`)
have no V2 equivalent and are silently dropped in V3→V2 conversion.

## Store Interface

```go
type Store interface {
    GetByGroups(ctx, groups) → map[string][]AclV3ListItem   // V3 view
    GetV2ByGroups(ctx, groups) → map[string][]AclV2ListItem // V2 view
}
```

Both methods query ALL versions (V2+V3) and convert to the requested format.

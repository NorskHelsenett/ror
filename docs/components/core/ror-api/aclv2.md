# ACLv2

## Scope/subject

Acl v2 implements an accessmodel in two layers, scope og subject.

### Scope

Scope represent the extent of the right. eg. if you want to grant access to a cluster the scope is **cluster**.

There is a global scope called **ror** that aplies to the whole system

All scopes are represented by the enum type [Acl2Scope](https://docs.ror.sky.test.nhn.no/code/internal/acl/models/#Acl2Scope).

### Subject

Subject represent id of the resource the access applies to. Eg. if the scope is **cluster** the subject is a clusterid.

If the scope is **ror** the subject is a group of subject eg **cluster** granting access to all clusters. The largest spanning subject is **Acl2RorSubjectGlobal**

Valid subjects under the scope **ror** is defined in the const **Acl2RorSubject**

Scopes are represented by the type [Acl2Subject](https://docs.ror.sky.test.nhn.no/code/internal/acl/models/#Acl2Subject) that represents a string.

### Validation

Scopes can be validated with the method _(s Acl2Scope) IsValid() bool_

Subjects can be validated against its coresponding scope with the method \*(s Acl2Subject) HasValidScope(scope Acl2Scope) bool

## Query

The query to the acl v2 engine should be represented by the type [AclV2QueryAccessScopeSubject](https://docs.ror.sky.test.nhn.no/code/internal/acl/models/#type-aclv2queryaccessscopesubject)

It **must** be created with the factory `NewAclV2QueryAccessScopeSubject(scope any, subject any) AclV2QueryAccessScopeSubject` which implements type casting and validation of the query.

## Access

Access is defined by the type [ AclV2ListItemAccess](https://docs.ror.sky.test.nhn.no/code/internal/acl/models/#type-aclv2listitemaccess) as five boolean values representing Read, Create, Update, Delete, Owner.

The returnes accessobject can be queried with the methods representing the level of acces we want to check eg:

```go
accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectGlobal)
accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
if !accessObject.Read {
    c.JSON(http.StatusForbidden, "403: No access")
    return
}
```

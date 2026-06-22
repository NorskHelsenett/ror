package aclmodels

import (
	"reflect"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
)

// V2ToV3 converts a V2 ACL entry to a V3 representation.
// The V3 access list is derived from the v3 struct tags on AclV2ListItemAccess
// and AclV2ListItemKubernetes. Only boolean fields that are true are included.
// V3 capabilities that have no V2 equivalent (e.g. "kubernetes:admin") cannot
// be represented and are absent from the result.
func V2ToV3(v2 AclV2ListItem) AclV3ListItem {
	var access []AccessTypeV3

	// Extract from AclV2ListItemAccess using v3 struct tags
	accessVal := reflect.ValueOf(v2.Access)
	accessType := accessVal.Type()
	for i := range accessType.NumField() {
		field := accessType.Field(i)
		tag := field.Tag.Get("v3")
		if tag == "" {
			continue
		}
		if accessVal.Field(i).Bool() {
			access = append(access, AccessTypeV3(tag))
		}
	}

	// Extract from AclV2ListItemKubernetes using v3 struct tags
	k8sVal := reflect.ValueOf(v2.Kubernetes)
	k8sType := k8sVal.Type()
	for i := range k8sType.NumField() {
		field := k8sType.Field(i)
		tag := field.Tag.Get("v3")
		if tag == "" {
			continue
		}
		if k8sVal.Field(i).Bool() {
			access = append(access, AccessTypeV3(tag))
		}
	}

	return AclV3ListItem{
		Id:       v2.Id,
		Version:  3,
		Group:    v2.Group,
		Scope:    aclscope.Scope(v2.Scope),
		Subject:  aclscope.Subject(v2.Subject),
		Access:   access,
		Created:  v2.Created,
		IssuedBy: v2.IssuedBy,
	}
}

// v3TagToV2Field maps v3 capability strings to the corresponding V2 boolean field setter.
var v3TagToV2Field = buildV3TagMap()

func buildV3TagMap() map[AccessTypeV3]string {
	m := make(map[AccessTypeV3]string)
	t := reflect.TypeFor[AclV2ListItemAccess]()
	for f := range t.Fields() {
		f := f
		if tag := f.Tag.Get("v3"); tag != "" {
			m[AccessTypeV3(tag)] = "access." + f.Name
		}
	}
	t = reflect.TypeFor[AclV2ListItemKubernetes]()
	for f := range t.Fields() {
		f := f
		if tag := f.Tag.Get("v3"); tag != "" {
			m[AccessTypeV3(tag)] = "kubernetes." + f.Name
		}
	}
	return m
}

// V3ToV2 converts a V3 ACL entry to a V2 representation.
// Only V3 access types that have a corresponding v3 struct tag on the V2 struct
// are mapped. V3-only capabilities (e.g. "kubernetes:admin", "resource:Deployment:read")
// are silently dropped since V2 has no way to represent them.
func V3ToV2(v3 AclV3ListItem) AclV2ListItem {
	v2 := AclV2ListItem{
		Id:       v3.Id,
		Version:  2,
		Group:    v3.Group,
		Scope:    Acl2Scope(v3.Scope),
		Subject:  Acl2Subject(v3.Subject),
		Created:  v3.Created,
		IssuedBy: v3.IssuedBy,
	}

	for _, a := range v3.Access {
		path, ok := v3TagToV2Field[a]
		if !ok {
			continue // V3-only capability, no V2 equivalent
		}
		switch path {
		case "access.Read":
			v2.Access.Read = true
		case "access.Create":
			v2.Access.Create = true
		case "access.Update":
			v2.Access.Update = true
		case "access.Delete":
			v2.Access.Delete = true
		case "access.Owner":
			v2.Access.Owner = true
		case "access.KubernetesLogon":
			v2.Access.KubernetesLogon = true
		case "kubernetes.Logon":
			v2.Kubernetes.Logon = true
		}
	}

	return v2
}

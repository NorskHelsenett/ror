// Package aclprincipal builds the ROR-owned group names that identify
// principals (clusters, services, service accounts) in ACL grants.
//
// Groups are principal names in an email/DNS hierarchy, read right-to-left:
//
//	<identity>@<qualifiers...>.<type>.ror.system
//
//	<cluster-uid>@cluster.ror.system
//	<service-id>@service.ror.system
//	<sa>@<namespace>.<cluster-uid>.sa.ror.system
//
// The package has no dependencies on other ror packages so both ror-api (which
// resolves grants today) and ror-auth (which will mint them into tokens) can
// share one source of truth for principal naming.
package aclprincipal

import "strings"

const (
	// Domain is the reserved suffix for ROR-minted principal groups. Groups in
	// this domain are authoritative for authorization and must never be accepted
	// from an external identity provider.
	Domain = "ror.system"

	// ClusterDomain qualifies cluster principals, keyed by cluster uid.
	ClusterDomain = "cluster." + Domain
	// ServiceDomain qualifies service principals, keyed by service id.
	ServiceDomain = "service." + Domain
	// ServiceAccountDomain qualifies Kubernetes ServiceAccount principals,
	// further qualified by cluster uid and namespace.
	ServiceAccountDomain = "sa." + Domain

	// Wildcard is the local part of an aggregate ("all of this type") group. It
	// is deliberately a character that cannot occur in a Kubernetes object name
	// or a cluster uid, so no principal can be named such that it inherits an
	// aggregate group.
	Wildcard = "*"
)

// Cluster returns the group identifying a single cluster by its uid.
func Cluster(uid string) string {
	return uid + "@" + ClusterDomain
}

// AllClusters returns the aggregate group every cluster is a member of.
func AllClusters() string {
	return Wildcard + "@" + ClusterDomain
}

// Service returns the group identifying a single service by its id.
func Service(id string) string {
	return id + "@" + ServiceDomain
}

// LegacyService returns the pre-hierarchy service group name. It is emitted
// alongside Service during the rename migration so grants keyed on the old name
// keep resolving, and is removed once no legacy entries remain.
func LegacyService(id string) string {
	return "service-" + id + "@" + Domain
}

// AllServices returns the aggregate group every service is a member of.
func AllServices() string {
	return Wildcard + "@" + ServiceDomain
}

// ServiceAccount returns the group identifying a single Kubernetes
// ServiceAccount within a namespace on a cluster.
func ServiceAccount(clusterUID, namespace, name string) string {
	return name + "@" + namespace + "." + clusterUID + "." + ServiceAccountDomain
}

// AllServiceAccountsInNamespace returns the aggregate group for every
// ServiceAccount in a namespace on a cluster.
func AllServiceAccountsInNamespace(clusterUID, namespace string) string {
	return Wildcard + "@" + namespace + "." + clusterUID + "." + ServiceAccountDomain
}

// AllServiceAccountsInCluster returns the aggregate group for every
// ServiceAccount on a cluster.
func AllServiceAccountsInCluster(clusterUID string) string {
	return Wildcard + "@" + clusterUID + "." + ServiceAccountDomain
}

// AllServiceAccounts returns the aggregate group every ServiceAccount is a
// member of.
func AllServiceAccounts() string {
	return Wildcard + "@" + ServiceAccountDomain
}

// ClusterGroups returns the full group membership of a cluster principal:
// its own group plus the all-clusters aggregate.
func ClusterGroups(uid string) []string {
	return []string{Cluster(uid), AllClusters()}
}

// ServiceGroups returns the full group membership of a service principal. The
// legacy name is included so grants are resolved during the rename migration.
func ServiceGroups(id string) []string {
	return []string{Service(id), LegacyService(id), AllServices()}
}

// ServiceAccountGroups returns the full group membership of a ServiceAccount
// principal: the exact account plus the namespace, cluster and fleet aggregates,
// so a grant can be attached at any level of the hierarchy.
func ServiceAccountGroups(clusterUID, namespace, name string) []string {
	return []string{
		ServiceAccount(clusterUID, namespace, name),
		AllServiceAccountsInNamespace(clusterUID, namespace),
		AllServiceAccountsInCluster(clusterUID),
		AllServiceAccounts(),
	}
}

// IsReserved reports whether a group name belongs to the ROR-owned principal
// domain. Such groups are minted by ROR from a verified identity and grant
// authorization directly, so accepting one from an external source would let a
// caller impersonate a cluster, service or ServiceAccount.
func IsReserved(group string) bool {
	at := strings.LastIndex(group, "@")
	if at < 0 {
		return false
	}
	domain := strings.ToLower(group[at+1:])
	return domain == Domain || strings.HasSuffix(domain, "."+Domain)
}

// SanitizeExternalGroups drops every reserved group from an externally supplied
// list (e.g. the groups claim of an identity-provider token).
func SanitizeExternalGroups(groups []string) []string {
	out := make([]string, 0, len(groups))
	for _, g := range groups {
		if IsReserved(g) {
			continue
		}
		out = append(out, g)
	}
	return out
}

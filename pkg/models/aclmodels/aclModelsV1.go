// aclmodels contains models for acl v1 and v2
package aclmodels

import identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"

// Used to query the v1 acl model
type AclV1QueryUserCluster struct {
	User      identitymodels.User
	ClusterId string
}

// Used to verify access using the v1 acl model
type AclV1DBResult struct {
	ClusterId string `bson:"clusterid"`
}

// Full acl v1 model
type AclV1ListItem struct {
	Cluster string `bson:"cluster"`
	Group   string `bson:"group"`
}

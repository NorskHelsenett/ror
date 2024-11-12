package identitymocks

import identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"

func ValiduserWithGroups(groups []string) identitymodels.Identity {
	return identitymodels.Identity{
		Type: identitymodels.IdentityTypeUser,
		User: &identitymodels.User{
			Email:           "valid.user@ror.dev",
			IsEmailVerified: true,
			Name:            "Valid User",
			Groups:          groups,
			Audience:        "",
			Issuer:          "",
			ExpirationTime:  0,
		},
	}
}

var IdentityUserValid identitymodels.Identity = identitymodels.Identity{
	Type: identitymodels.IdentityTypeUser,
	User: &identitymodels.User{
		Email:           "valid.user@ror.dev",
		IsEmailVerified: true,
		Name:            "Valid User",
		Groups: []string{
			"test1@ror.dev",
			"test1-admin@ror.dev",
		},
		Audience:       "",
		Issuer:         "",
		ExpirationTime: 0,
	},
}

var IdentityClusterValid identitymodels.Identity = identitymodels.Identity{
	Type: identitymodels.IdentityTypeCluster,
	ClusterIdentity: &identitymodels.ServiceIdentity{
		Id: "test-cluster-43232",
	},
	ServiceIdentity: &identitymodels.ServiceIdentity{
		Id: "",
	},
}

var IdentityServiceValid identitymodels.Identity = identitymodels.Identity{
	Type: identitymodels.IdentityTypeCluster,
	ServiceIdentity: &identitymodels.ServiceIdentity{
		Id: "serivce-test@ror.system",
	},
}

package modelhelper

import (
	"errors"

	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"go.mongodb.org/mongo-driver/bson"
)

func DefaultClusterFilter() apicontracts.Filter {
	return apicontracts.Filter{
		Skip:  0,
		Limit: 20,
		Sort: []apicontracts.SortMetadata{
			{
				SortField: "clusterid",
				SortOrder: 1,
			},
		},
	}
}

func ExtractGroupsFromUser(user *identitymodels.User) (bson.A, error) {
	if user == nil {
		return nil, errors.New("user is nil")
	}

	if len(user.Groups) == 0 {
		return nil, errors.New("User.Groups is empty")
	}

	filterGroups := bson.A{}
	for i := 0; i < len(user.Groups); i++ {
		filterGroups = append(filterGroups, user.Groups[i])
	}

	return filterGroups, nil
}

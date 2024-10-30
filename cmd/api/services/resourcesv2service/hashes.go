package resourcesv2service

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"
)

func GetHashList(ctx context.Context, owner rortypes.RorResourceOwnerReference) (apiresourcecontracts.HashList, error) {

	databaseHelpers := NewResourceMongoDB(mongodb.GetMongodbConnection())
	hashList, err := databaseHelpers.GetHashList(ctx, owner)
	if err != nil {
		return apiresourcecontracts.HashList{}, err
	}
	return hashList, nil
}

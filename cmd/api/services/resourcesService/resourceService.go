// the resource service package provides services to get and manipulate resources
package resourcesservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/api/apiconnections"
	"github.com/NorskHelsenett/ror/internal/mongodbrepo/repositories/resourcesmongodbrepo"

	"github.com/NorskHelsenett/ror/pkg/messagebuscontracts"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CheckResourceExist checks whether a resource with the provided `uid` exists in the MongoDB database.
// The function uses the `resourcesmongodbrepo.ResourceExistbyUid` function to query the database and returns
// `true` if a matching resource is found, otherwise it returns `false`.
func CheckResourceExist(ctx context.Context, uid string) bool {
	return resourcesmongodbrepo.ResourceExistbyUid(uid, ctx)
}

// GetResources retrieves resources of type `T` from Mongo DB based on the provided `ResourceQuery`.
// The function queries the `resourcesmongodbrepo` using `GetResourcesByQuery[T]` method and returns a slice of the retrieved resources if successful.
// The function returns an error if the resource retrieval process fails.
func GetResources[T apiresourcecontracts.Resourcetypes](ctx context.Context, query apiresourcecontracts.ResourceQuery) ([]T, error) {
	return resourcesmongodbrepo.GetResourcesByQuery[T](ctx, query)
}

// PatchResource updates a resource in the MongoDB database based on the provided `ResourceQuery` and `ResourceUpdateModel`.
// The function returns an error if the resource update process fails.
// The resourceUpdate parameter is a `bson.M` type which should be flattened to the following format:
//
//	bson.M{
//	    "metadata.name": "test",
//	}
//
// This function is inteded used by internal functions and does not perform any validation on the provided parameters.
func PatchResource(ctx context.Context, uid string, resourceUpdate bson.M) (mongo.UpdateResult, error) {
	query := bson.M{
		"$set": resourceUpdate,
	}

	return resourcesmongodbrepo.PatchResource(ctx, uid, query)
}

// Get one resource by query (owner/apiVersion/Kind/uid)
func GetResource[T apiresourcecontracts.Resourcetypes](ctx context.Context, query apiresourcecontracts.ResourceQuery) (T, error) {
	var emptyresult T

	if query.Uid == "" {
		err := fmt.Errorf("uid is empty")
		return emptyresult, err
	}
	result, err := resourcesmongodbrepo.GetResourcesByQuery[T](ctx, query)
	if err != nil {
		return emptyresult, err
	}
	if len(result) != 1 {
		return emptyresult, errors.New("could not find exactly one resource")
	}
	return result[0], nil
}

// wrapper to allow create to update if hashlist is bugged or download failed
func ResourceNewCreateService(ctx context.Context, resourceUpdate apiresourcecontracts.ResourceUpdateModel) error {
	if resourcesmongodbrepo.ResourceExistbyUid(resourceUpdate.Uid, ctx) {
		err := ResourceUpdateService(ctx, resourceUpdate)
		if err != nil {
			return err
		}
	} else {
		err := ResourceCreateService(ctx, resourceUpdate)
		if err != nil {
			return err
		}
	}
	return nil
}

// Function deletes a resource
func ResourceDeleteService(ctx context.Context, resourceUpdate apiresourcecontracts.ResourceUpdateModel) error {
	err := resourcesmongodbrepo.DeleteResourceByUid(resourceUpdate, ctx)
	if err != nil {
		rlog.Errorc(ctx, "could not update resource", err)
		return err
	}

	if err := apiconnections.RabbitMQConnection.SendMessage(ctx,
		resourceUpdate,
		messagebuscontracts.Route_ResourceDeleted,
		map[string]interface{}{"apiVersion": resourceUpdate.ApiVersion, "kind": resourceUpdate.Kind}); err != nil {
		return err
	}

	//return switchboard.PublishResourceToSwitchboard(ctx, messages.RulesetRuleTypeDeleted, resourceUpdate)
	return nil
}

// returns the list of hashes owned by the ownerref
func ResourceGetHashlist(ctx context.Context, owner apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.HashList, error) {
	return resourcesmongodbrepo.GetHashList(ctx, owner)
}

func sendToMessageBus(ctx context.Context, resource any, action apiresourcecontracts.ResourceAction) error {
	b, err := json.Marshal(resource)
	if err != nil {
		return errors.New("could not cast resource to byte[]")
	}

	var payload apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceNamespace]
	err = json.Unmarshal(b, &payload)
	if err != nil {
		rlog.Error("Could not convert to json", err)
		return errors.New("could not cast resource to ResourceNamespace")
	}
	payload.Version = apiresourcecontracts.ResourceVersionV1

	switch action {
	case apiresourcecontracts.K8sActionAdd:
		_ = apiconnections.RabbitMQConnection.SendMessage(ctx,
			payload,
			messagebuscontracts.Route_ResourceCreated,
			map[string]interface{}{"apiVersion": payload.ApiVersion, "kind": payload.Kind})
	case apiresourcecontracts.K8sActionUpdate:
		_ = apiconnections.RabbitMQConnection.SendMessage(ctx,
			payload,
			messagebuscontracts.Route_ResourceUpdated,
			map[string]interface{}{"apiVersion": payload.ApiVersion, "kind": payload.Kind})
	}
	return nil
}

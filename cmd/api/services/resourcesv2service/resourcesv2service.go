package resourcesv2service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/api/apiconnections"
	aclservice "github.com/NorskHelsenett/ror/internal/acl/services"
	"net/http"
	"time"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"
	"github.com/NorskHelsenett/ror/pkg/messagebuscontracts"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/NorskHelsenett/ror/pkg/rorresources"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"
)

func HandleResourceUpdate(ctx context.Context, resource *rorresources.Resource) rorresources.ResourceUpdateResults {
	switch resource.GetRorMeta().Action {
	case rortypes.K8sActionAdd:
		return NewOrUpdateResource(ctx, resource)
	case rortypes.K8sActionUpdate:
		return NewOrUpdateResource(ctx, resource)
	case rortypes.K8sActionDelete:
		err := DeleteResource(ctx, resource)
		if err != nil {
			rlog.Error("Could not delete resource", err)
			return rorresources.ResourceUpdateResults{
				Results: map[string]rorresources.ResourceUpdateResult{
					resource.GetUID(): {
						Status:  http.StatusInternalServerError,
						Message: "500: Could not delete resource",
					},
				},
			}
		}
		return rorresources.ResourceUpdateResults{
			Results: map[string]rorresources.ResourceUpdateResult{
				resource.GetUID(): {
					Status:  http.StatusAccepted,
					Message: "202: Resource deleted",
				},
			},
		}
	default:
		return rorresources.ResourceUpdateResults{
			Results: map[string]rorresources.ResourceUpdateResult{
				resource.GetUID(): {
					Status:  http.StatusBadRequest,
					Message: "400: Unknown action",
				},
			},
		}
	}
}

func NewOrUpdateResource(ctx context.Context, resource *rorresources.Resource) rorresources.ResourceUpdateResults {
	ownerref := resource.GetRorMeta().Ownerref

	// Access check
	// Scope: input.Owner.Scope
	// Subject: input.Owner.Subject
	// Access: create
	accessObject := aclservice.CheckAccessByRorOwnerref(ctx, ownerref)
	if !accessObject.Create {
		return rorresources.ResourceUpdateResults{
			Results: map[string]rorresources.ResourceUpdateResult{
				resource.GetUID(): {
					Status:  http.StatusForbidden,
					Message: "403: No access",
				},
			},
		}
	}

	err := resource.ApplyInputFilter()
	if err != nil {
		return rorresources.ResourceUpdateResults{
			Results: map[string]rorresources.ResourceUpdateResult{
				resource.GetUID(): {
					Status:  http.StatusBadRequest,
					Message: "400: Could not apply filter to resource",
				},
			},
		}
	}
	cache := GetResourceCache()
	cache.Set(ctx, resource)

	databaseHelpers := NewResourceMongoDB(mongodb.GetMongodbConnection())
	err = databaseHelpers.Set(ctx, resource)
	if err != nil {
		rlog.Errorc(ctx, "Failed to set resource", err)
		return rorresources.ResourceUpdateResults{
			Results: map[string]rorresources.ResourceUpdateResult{
				resource.GetUID(): {
					Status:  http.StatusInternalServerError,
					Message: "500: Could not create resource",
				},
			},
		}
	}

	sendToMessageBus(ctx, resource, resource.RorMeta.Action)

	//rlog.Info("Resource created", rlog.Any("resource", resource.GetAPIVersion()), rlog.Any("kind", resource.GetKind()), rlog.Any("name", resource.GetName()))
	return rorresources.ResourceUpdateResults{
		Results: map[string]rorresources.ResourceUpdateResult{
			resource.GetUID(): {
				Status:  http.StatusAccepted,
				Message: "202: Resource created",
			},
		},
	}
}

func GetResourceByUID(ctx context.Context, uid string) *rorresources.ResourceSet {
	var returnrs *rorresources.ResourceSet
	start := time.Now()
	cache := GetResourceCache()
	resource := cache.Get(ctx, uid)
	if resource != nil {
		returnrs = rorresources.NewResourceSet()
		returnrs.Resources = append(returnrs.Resources, resource)
		rlog.Debug("Resource found in cache", rlog.String("uid", uid), rlog.Any("duration", time.Since(start)))
	} else {
		databaseHelpers := NewResourceMongoDB(mongodb.GetMongodbConnection())
		var err error
		returnrs, err = databaseHelpers.Get(ctx, rorresources.NewResourceQuery().WithUID(uid))
		if err != nil {
			rlog.Error("Could not get resource by uid", err, rlog.String("uid", uid), rlog.Any("error", err))
			return nil
		}
		if returnrs == nil {
			return nil
		}
		cache.Set(ctx, returnrs.Resources[0])

		rlog.Debug("Resource found in database", rlog.String("uid", uid), rlog.Any("duration", time.Since(start)))
	}

	// Access check
	// Scope: input.Owner.Scope
	// Subject: input.Owner.Subject
	// Access: read
	for _, resource := range returnrs.Resources {
		accessModel := aclservice.CheckAccessByRorOwnerref(ctx, resource.GetRorMeta().Ownerref)
		if !accessModel.Read {
			return nil
		}
	}

	return returnrs
}

func DeleteResource(ctx context.Context, resource *rorresources.Resource) error {
	// Access check
	// Scope: input.Owner.Scope
	// Subject: input.Owner.Subject
	// Access: delete

	accessModel := aclservice.CheckAccessByRorOwnerref(ctx, resource.GetRorMeta().Ownerref)
	if !accessModel.Update {
		err := fmt.Errorf("403: No access to uid %s", resource.GetUID())
		return err
	}
	cache := GetResourceCache()
	cache.Remove(ctx, resource.GetUID())
	databaseHelpers := NewResourceMongoDB(mongodb.GetMongodbConnection())
	return databaseHelpers.Del(ctx, resource)
}

func GetResourceByQuery(ctx context.Context, query *rorresources.ResourceQuery) *rorresources.ResourceSet {
	start := time.Now()
	databaseHelpers := NewResourceMongoDB(mongodb.GetMongodbConnection())
	rs, err := databaseHelpers.Get(ctx, query)
	if err != nil {
		rlog.Error("Could not get resource by query", err, rlog.Any("error", err))
		return nil
	}
	if rs == nil {
		return nil
	}

	// Access check
	// Scope: input.Owner.Scope
	// Subject: input.Owner.Subject
	// Access: read

	returnrs := rorresources.NewResourceSet()
	var checkedOwnerRef = make(map[string]int, 0)
	for _, resource := range rs.Resources {
		if checked, ok := checkedOwnerRef[resource.GetRorMeta().Ownerref.String()]; ok {
			if checked == 1 {
				returnrs.Add(resource)
			}
			continue
		}
		accessModel := aclservice.CheckAccessByRorOwnerref(ctx, resource.GetRorMeta().Ownerref)
		if accessModel.Read {
			checkedOwnerRef[resource.GetRorMeta().Ownerref.String()] = 1
			returnrs.Add(resource)
			continue
		} else {
			checkedOwnerRef[resource.GetRorMeta().Ownerref.String()] = -1
		}
	}
	rlog.Debug("Resource found in database", rlog.Any("number", len(rs.Resources)), rlog.Any("duration", time.Since(start)))
	return returnrs
}

func sendToMessageBus(ctx context.Context, resource *rorresources.Resource, action rortypes.ResourceAction) error {
	b, err := json.Marshal(resource)
	if err != nil {
		return errors.New("could not cast resource to byte[]")
	}

	payload := apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceNamespace]{}
	payload.ApiVersion = resource.GetAPIVersion()
	payload.Kind = resource.GetKind()
	payload.Uid = resource.GetUID()
	payload.Hash = resource.GetRorHash()
	payload.Internal = resource.GetRorMeta().Internal
	payload.Owner.Scope = resource.GetRorMeta().Ownerref.Scope
	payload.Owner.Subject = string(resource.GetRorMeta().Ownerref.Subject)
	payload.Version = apiresourcecontracts.ResourceVersionV2
	err = json.Unmarshal(b, &payload)
	if err != nil {
		rlog.Error("Could not convert to json", err)
		return errors.New("could not cast resource to ResourceNamespace")
	}

	switch action {
	case rortypes.K8sActionAdd:
		_ = apiconnections.RabbitMQConnection.SendMessage(ctx,
			payload,
			messagebuscontracts.Route_ResourceCreated,
			map[string]interface{}{"apiVersion": payload.ApiVersion, "kind": payload.Kind})
	case rortypes.K8sActionUpdate:
		_ = apiconnections.RabbitMQConnection.SendMessage(ctx,
			payload,
			messagebuscontracts.Route_ResourceUpdated,
			map[string]interface{}{"apiVersion": payload.ApiVersion, "kind": payload.Kind})
	}
	return nil
}

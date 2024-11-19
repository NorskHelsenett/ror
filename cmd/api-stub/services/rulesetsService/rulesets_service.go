package rulesetsService

import (
	"context"
	"errors"
	"fmt"

	resourcesservice "github.com/NorskHelsenett/ror/cmd/api-stub/services/resourcesService"
	"github.com/NorskHelsenett/ror/internal/mongodbrepo/repositories/rulesetsRepo"

	aclmodels "github.com/NorskHelsenett/ror/pkg/models/acl"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/apicontracts/messages"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

func Create(ctx context.Context, clusterId string) (*messages.RulesetModel, error) {
	set := new(messages.RulesetModel)

	set.Identity.Type = messages.RulesetIdentityTypeCluster
	set.Identity.Id = clusterId

	if err := rulesetsRepo.Create(ctx, set); err != nil {
		return nil, err
	}

	return set, nil
}

func CreateInternal(ctx context.Context) (*messages.RulesetModel, error) {
	set := new(messages.RulesetModel)

	set.Identity.Type = messages.RulesetIdentityTypeInternal
	set.Identity.Id = "internal-primary"

	if err := rulesetsRepo.Create(ctx, set); err != nil {
		return nil, err
	}

	return set, nil
}

func FindInternal(ctx context.Context) (*messages.RulesetModel, error) {
	set, err := rulesetsRepo.FindInternal(ctx)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}

		return CreateInternal(ctx)
	}

	return set, err
}

func FindAll(ctx context.Context) ([]*messages.RulesetModel, error) {
	return rulesetsRepo.FindAll(ctx)
}

func FindCluster(ctx context.Context, clusterId string) (*messages.RulesetModel, error) {
	set, err := rulesetsRepo.FindCluster(ctx, clusterId)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}

		return Create(ctx, clusterId)
	}

	return set, err
}

func Find(ctx context.Context, setId string) (*messages.RulesetModel, error) {
	return rulesetsRepo.FindById(ctx, setId)
}

func DeleteResource(ctx context.Context, setId string, resourceId string) error {
	set, err := rulesetsRepo.FindById(ctx, setId)
	if err != nil {
		return err
	}

	if err := rulesetsRepo.PullResource(ctx, set, resourceId); err != nil {
		return nil
	}

	return nil
}

func AddResource(ctx context.Context, setId string, input *messages.RulesetResourceInput) (*messages.RulesetResourceModel, error) {
	model := new(messages.RulesetResourceModel)

	set, err := rulesetsRepo.FindById(ctx, setId)
	if err != nil {
		return nil, err
	}

	switch set.Identity.Type {
	case messages.RulesetIdentityTypeCluster:
		{
			rlog.Info("adding resource for cluster")

			if input.Uid != "*" {
				resource, err := resourcesservice.GetResource[apiresourcecontracts.ResourceIngress](ctx, apiresourcecontracts.ResourceQuery{
					Owner: apiresourcecontracts.ResourceOwnerReference{
						Scope:   aclmodels.Acl2ScopeCluster,
						Subject: set.Identity.Id,
					},
					Uid:        input.Uid,
					ApiVersion: input.ApiVersion,
					Kind:       input.Kind,
				})

				if err != nil {
					return nil, err
				}
				rlog.Info("found resource", rlog.Any("resource", resource))

				model.Ref.ApiVersion = resource.ApiVersion
				model.Ref.Uid = resource.Metadata.Uid
				model.Ref.Kind = resource.Kind
				model.Ref.Name = resource.Metadata.Name
				model.Ref.Namespace = resource.Metadata.Namespace
			} else {
				model.Ref.ApiVersion = input.ApiVersion
				model.Ref.Uid = "*"
				model.Ref.Name = fmt.Sprintf("%s - *", input.Kind)
				model.Ref.Kind = input.Kind
				model.Ref.Namespace = input.Namespace
			}
		}
	case messages.RulesetIdentityTypeInternal:
		{
			model.Ref.Kind = input.Kind
			model.Ref.ApiVersion = input.ApiVersion
			model.Ref.Uid = input.Uid
			model.Ref.Kind = "RorEntity"
			model.Ref.Name = input.Uid
			model.Ref.Namespace = "ror"
		}
	}

	model.Id = uuid.NewString()

	if err := rulesetsRepo.AddResource(ctx, set, model); err != nil {
		return nil, err
	}

	return model, nil
}

func AddResourceRule(ctx context.Context, setId string, resourceId string, input *messages.RulesetRuleInput) (*messages.RulesetRuleModel, error) {
	model := new(messages.RulesetRuleModel)

	model.Id = uuid.NewString()
	model.Lifetime = input.Lifetime
	model.Type = input.Type
	model.Service = input.Service

	model.Slack = input.Slack

	set, err := rulesetsRepo.FindById(ctx, setId)
	if err != nil {
		return nil, err
	}
	rlog.Infoc(ctx, "found set", rlog.String("id", set.ID))

	resource, err := set.FindResourceById(resourceId)
	if err != nil {
		return nil, err
	}
	rlog.Infoc(ctx, "found resource", rlog.String("id", resource.Id))

	if err := rulesetsRepo.AddResourceRule(ctx, set, resource, model); err != nil {
		return nil, err
	}
	rlog.Infoc(ctx, "added event", rlog.String("id", model.Id))

	return model, nil
}

func DeleteResourceRule(ctx context.Context, setId string, resourceId string, ruleId string) error {
	set, err := rulesetsRepo.FindById(ctx, setId)
	if err != nil {
		return err
	}

	resource, err := set.FindResourceById(resourceId)
	if err != nil {
		return err
	}

	if err := rulesetsRepo.PullResourceRule(ctx, set, resource, ruleId); err != nil {
		return nil
	}

	return nil
}

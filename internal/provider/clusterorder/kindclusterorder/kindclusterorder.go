package kindclusterorder

import (
	"context"
	"encoding/json"
	"github.com/NorskHelsenett/ror/internal/provider/clusterorder/utils"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/rlog"
)

type ClusterOrderKind struct {
	order apiresourcecontracts.ResourceClusterOrder
	saved bool
}

// NewClusterOrderKind creates a new instance of ClusterOrderKind based on the provided order.
// It takes a context.Context and an apiresourcecontracts.ResourceClusterOrder as input parameters.
// If the order's Metadata.Uid is empty, it creates a new ClusterOrderResource using the order's Spec.
// If there is an error during the creation of the ClusterOrderResource, it returns the error.
// Otherwise, it sets the order field of the ClusterOrderKind to the provided order and sets the saved field to true.
// Finally, it returns the created ClusterOrderKind and nil error.
func NewClusterOrderKind(ctx context.Context, order apiresourcecontracts.ResourceClusterOrder) (ClusterOrderKind, error) {
	var ct ClusterOrderKind
	var err error
	if order.Metadata.Uid == "" {
		ct.order, err = utils.NewClusterOrderResource(ctx, order.Spec)
		if err != nil {
			rlog.Errorc(ctx, "error creating cluster order", err)
			return ct, err
		}
	} else {
		ct.order = order
		ct.saved = true
	}

	return ct, nil
}

func (c ClusterOrderKind) Validate(ctx context.Context) error {
	err := utils.ValidateOrder(ctx, c.order.Spec)
	if err != nil {
		rlog.Error("error validating order", err)
		return err
	}

	var providerConfig apiresourcecontracts.ResourceProviderConfigKind
	jsonString, _ := json.Marshal(c.order.Spec.ProviderConfig)
	err = json.Unmarshal(jsonString, &providerConfig)
	if err != nil {
		rlog.Errorc(ctx, "could not cast to kindProviderConfig", err)
		return err
	}

	return nil
}

func (c ClusterOrderKind) GetProviderConfig() any {
	var providerConfig apiresourcecontracts.ResourceProviderConfigKind
	jsonString, _ := json.Marshal(c.order.Spec.ProviderConfig)
	err := json.Unmarshal(jsonString, &providerConfig)
	if err != nil {
		rlog.Error("could not cast to kindProviderConfig", err)
		return nil
	}
	return providerConfig
}

func (c *ClusterOrderKind) UpdateStatus(ctx context.Context, status apiresourcecontracts.ResourceClusterOrderStatus) error {
	if !c.saved {
		rlog.Debug("Order not saved, appending to clusterorderspec")
		if status.Status != "" {
			c.order.Status.Status = status.Status
		}
		if status.Phase != "" {
			c.order.Status.Phase = status.Phase
		}
		return nil
	}
	rlog.Debug("Updating clusterorder status", rlog.Any("status", status))
	return utils.UpdateStatus(ctx, c.order.Metadata.Uid, status)
}

func (c *ClusterOrderKind) Save(ctx context.Context) error {
	ret, err := utils.NewResourceUpdate(ctx, c.order)
	if err != nil {
		rlog.Errorc(ctx, "error saving cluster order", err)
		return err
	}
	err = utils.CreateResource(ctx, ret)
	if err == nil {
		c.saved = true
	}
	return err
}

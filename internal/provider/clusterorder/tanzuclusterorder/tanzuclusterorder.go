package tanzuclusterorder

import (
	"context"
	"encoding/json"
	"errors"
	workspacesservice "github.com/NorskHelsenett/ror/cmd/api/services/workspacesService"
	"github.com/NorskHelsenett/ror/internal/provider/clusterorder/utils"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

type ClusterOrderTanzu struct {
	order apiresourcecontracts.ResourceClusterOrder
	saved bool
}

func NewClusterOrderTanzu(ctx context.Context, order apiresourcecontracts.ResourceClusterOrder) (ClusterOrderTanzu, error) {

	var ct ClusterOrderTanzu
	var err error
	if order.Metadata.Uid == "" {
		ct.order, err = utils.NewClusterOrderResource(ctx, order.Spec)
		if err != nil {
			return ct, err
		}
	} else {
		ct.order = order
		ct.saved = true
	}

	return ct, nil
}

func (c ClusterOrderTanzu) Validate(ctx context.Context) error {

	err := utils.ValidateOrder(ctx, c.order.Spec)
	if err != nil {
		rlog.Error("error validating order", err)
		return err
	}

	var providerConfig apiresourcecontracts.ResourceProviderConfigTanzu
	jsonString, _ := json.Marshal(c.order.Spec.ProviderConfig)
	err = json.Unmarshal(jsonString, &providerConfig)
	if err != nil {
		rlog.Error("could not cast to tanzuProviderConfig", err)
		return err
	}

	workspace, err := workspacesservice.GetById(ctx, providerConfig.NamespaceId)
	if err != nil {
		rlog.Errorc(ctx, "error getting workspace", err)
		return err
	}
	if workspace == nil {
		rlog.Errorc(ctx, "workspace not found", errors.New("workspace not found, namespaceId: "+providerConfig.NamespaceId))
		return errors.New("workspace not found, namespaceId: " + providerConfig.NamespaceId)
	}

	if providerConfig.DatacenterId != workspace.DatacenterID {
		return errors.New("datacenterId does not match workspace datacenterId (id )")
	}

	//TODO: Validate nodepools (machineclasses...)
	//TODO: Implement kubernetes version validation

	return nil
}

func (c ClusterOrderTanzu) GetProviderConfig() any {
	var providerConfig apiresourcecontracts.ResourceProviderConfigTanzu
	jsonString, _ := json.Marshal(c.order.Spec.ProviderConfig)
	err := json.Unmarshal(jsonString, &providerConfig)
	if err != nil {
		rlog.Error("could not cast to tanzuProviderConfig", err)
		return nil
	}
	return providerConfig
}

func (c *ClusterOrderTanzu) UpdateStatus(ctx context.Context, status apiresourcecontracts.ResourceClusterOrderStatus) error {
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

func (c *ClusterOrderTanzu) Save(ctx context.Context) error {
	ret, err := utils.NewResourceUpdate(ctx, c.order)
	if err != nil {
		rlog.Errorc(ctx, "error creating resource", err)
		return err
	}
	err = utils.CreateResource(ctx, ret)
	if err == nil {
		c.saved = true
	}
	return err
}

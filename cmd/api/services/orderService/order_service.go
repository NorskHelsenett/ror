package orderservice

import (
	"context"
	"github.com/NorskHelsenett/ror/internal/provider/clusterorder"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"
)

func OrderCluster(ctx context.Context, orderspec apiresourcecontracts.ResourceClusterOrderSpec) error {
	order, err := clusterorder.NewClusterOrder(ctx, orderspec)
	if err != nil {
		rlog.Errorc(ctx, "error creating cluster order", err)
		return err
	}
	err = order.Validate(ctx)
	if err != nil {
		rlog.Errorc(ctx, "error validating cluster order", err)
		return err
	}
	err = order.UpdateStatus(ctx, apiresourcecontracts.ResourceClusterOrderStatus{
		Phase:  apiresourcecontracts.ResourceClusterOrderStatusPhaseCreating,
		Status: "Accepted",
	})
	if err != nil {
		rlog.Error("could not update status", err)
	}

	err = order.Save(ctx)
	if err != nil {
		rlog.Errorc(ctx, "error saving cluster order", err)
		return err
	}

	return nil
}

// The package clusterorder provides the interface and the factory method to create a new clusterorder based on the provider
// TODO: This might be a implementation in the upcoming ResourceFramework
package clusterorder

import (
	"context"
	"errors"

	clustersservice "github.com/NorskHelsenett/ror/cmd/api/services/clustersService"
	"github.com/NorskHelsenett/ror/internal/provider/clusterorder/kindclusterorder"
	"github.com/NorskHelsenett/ror/internal/provider/clusterorder/talosclusterorder"
	"github.com/NorskHelsenett/ror/internal/provider/clusterorder/tanzuclusterorder"

	"github.com/NorskHelsenett/ror/pkg/context/rorcontext"
	providersconsts "github.com/NorskHelsenett/ror/pkg/models/providers"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// ClusterOrder is an interface that defines the methods that a clusterprovider must implement
type ClusterOrder interface {
	Validate(ctx context.Context) error
	GetProviderConfig() any
	Save(ctx context.Context) error
	UpdateStatus(ctx context.Context, status apiresourcecontracts.ResourceClusterOrderStatus) error
}

// NewClusterOrder returns a new clusterorder based on the provider
func NewClusterOrder(ctx context.Context, orderspec apiresourcecontracts.ResourceClusterOrderSpec) (ClusterOrder, error) {

	identity := rorcontext.GetIdentityFromRorContext(ctx)
	if identity.GetId() == "" {
		return nil, errors.New("user not found")
	}

	orderspec.OrderBy = identity.GetId()

	if orderspec.OrderType != apiresourcecontracts.ResourceActionTypeCreate {
		cluster, err := clustersservice.GetByClusterId(ctx, orderspec.Cluster)
		if err != nil {
			return nil, err
		}
		if cluster == nil {
			return nil, errors.New("cluster not found")
		}
		orderspec.Provider = cluster.Workspace.Datacenter.Provider
	} else {
		if orderspec.Cluster == "" {
			return nil, errors.New("clustername can not be empty")
		}
	}

	if providersconsts.ProviderType(orderspec.Provider) == "" {
		return nil, errors.New("provider not supported")
	}

	order := apiresourcecontracts.ResourceClusterOrder{
		Spec: orderspec,
	}
	return newClusterOrder(ctx, order)
}

// NewClusterOrderFromResource returns the interface ClusterOrder based on the provider specified in the apiresourcecontracts.ResourceClusterOrder
func newClusterOrder(ctx context.Context, order apiresourcecontracts.ResourceClusterOrder) (ClusterOrder, error) {
	switch order.Spec.Provider {
	case providersconsts.ProviderTypeTanzu:
		tanzuOrder, err := tanzuclusterorder.NewClusterOrderTanzu(ctx, order)
		return &tanzuOrder, err
	case providersconsts.ProviderTypeKind:
		kindOrder, err := kindclusterorder.NewClusterOrderKind(ctx, order)
		return &kindOrder, err
	case providersconsts.ProviderTypeTalos:
		kindOrder, err := talosclusterorder.NewClusterOrderTalos(ctx, order)
		return &kindOrder, err
	default:
		return nil, errors.New("provider not supported")
	}
}

// NewClusterOrderFromResource is a wrapper for the newClusterOrder function.
// Its intended use is to be called from the outside of the package after the clusterorder is persisted.
func NewClusterOrderFromResource(ctx context.Context, resource apiresourcecontracts.ResourceClusterOrder) (ClusterOrder, error) {
	return newClusterOrder(ctx, resource)
}

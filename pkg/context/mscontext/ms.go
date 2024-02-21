// The package provides ror-context for services
package mscontext

import (
	"context"
	"time"

	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"
)

// Function provides a ror context for a given servicename, the context must be used to authenticate against existing services
func GetRorContextFromServiceContext(c *context.Context, servicename string) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(*c, 10*time.Second)
	identity := identitymodels.Identity{
		Type:            identitymodels.IdentityTypeService,
		ServiceIdentity: &identitymodels.ServiceIdentity{Id: servicename},
	}
	ctx = context.WithValue(ctx, identitymodels.ContexIdentity, identity)
	return ctx, cancel
}

// GetRorContextFromServiceContextWithoutCancel  Function provides a ror context for a given servicename, the context must be used to authenticate against existing services.
// The functionality is the same except that it forces a cancel on us.
func GetRorContextFromServiceContextWithoutCancel(c context.Context, servicename string) context.Context {
	identity := identitymodels.Identity{
		Type:            identitymodels.IdentityTypeService,
		ServiceIdentity: &identitymodels.ServiceIdentity{Id: servicename},
	}
	ctx := context.WithValue(c, identitymodels.ContexIdentity, identity)
	return ctx
}

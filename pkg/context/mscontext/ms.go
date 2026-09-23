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
	return context.WithValue(ctx, identitymodels.ContexIdentity, serviceIdentity(servicename)), cancel
}

// GetRorContextFromServiceContextWithoutCancel  Function provides a ror context for a given servicename, the context must be used to authenticate against existing services.
// The functionality is the same except that it forces a cancel on us.
func GetRorContextFromServiceContextWithoutCancel(c context.Context, servicename string) context.Context {
	return context.WithValue(c, identitymodels.ContexIdentity, serviceIdentity(servicename))
}

// serviceIdentity builds the service identity placed in the context. A name that
// cannot form a valid identity yields the zero value, which fails closed on every
// authorization check.
func serviceIdentity(servicename string) identitymodels.Identity {
	identity, err := identitymodels.NewServiceIdentity(identitymodels.AuthInfo{}, servicename)
	if err != nil {
		return identitymodels.Identity{}
	}
	return identity
}

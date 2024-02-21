// Package contextservices The package provides functions to get and work with ror contexts
package rorcontext

import (
	"context"
	"fmt"

	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"

	"github.com/NorskHelsenett/ror/pkg/rlog"
)

// GetIdentityFromRorContext Function returns the identity from the ror context.
func GetIdentityFromRorContext(ctx context.Context) identitymodels.Identity {
	identity, ok := ctx.Value(identitymodels.ContexIdentity).(identitymodels.Identity)
	if !ok {
		rlog.Error("failed to get identity from RorContext", fmt.Errorf("error getting identity from context"))
		panic("Faild to get identity")
	}
	return identity
}

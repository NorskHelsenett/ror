// Package contextservices The package provides functions to get and work with ror contexts
package rorcontext

import (
	"context"
	"fmt"

	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"

	"github.com/NorskHelsenett/ror/pkg/rlog"
)

type RorContext = context.Context

// MustGetIdentityFromRorContext Function returns the identity from the ror context.
func MustGetIdentityFromRorContext(ctx RorContext) identitymodels.Identity {
	identity, ok := ctx.Value(identitymodels.ContexIdentity).(identitymodels.Identity)
	if !ok {
		rlog.Error("failed to get identity from RorContext", fmt.Errorf("error getting identity from context"))
		panic("Faild to get identity")
	}
	return identity
}

// MustGetIdentityFromRorContext Function returns the identity from the ror context.
func GetIdentityFromRorContext(ctx RorContext) (identitymodels.Identity, error) {
	identity, ok := ctx.Value(identitymodels.ContexIdentity).(identitymodels.Identity)
	if !ok {
		err := fmt.Errorf("error getting identity from context")
		rlog.Error("failed to get identity from RorContext", err)
		return identitymodels.Identity{}, err
	}
	return identity, nil
}

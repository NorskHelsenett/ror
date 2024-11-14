// Package contextservices The package provides functions to get and work with ror contexts
package gincontext

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"

	"github.com/NorskHelsenett/ror/pkg/helpers/rorerror"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/gin-gonic/gin"
)

// GetRorContextFromGinContext Function creates ror context from gin context, identity is added to the context
func GetRorContextFromGinContext(c *gin.Context) (context.Context, context.CancelFunc) {
	var timeout time.Duration
	timeoutstring, exists := c.Get("timeout")
	if !exists {
		msg := fmt.Sprintf("timeout not set in gin context for path %s", c.Request.URL.Path)
		rlog.Warn(msg)
		timeout = 10 * time.Second
	} else {
		rlog.Debug("timeout set in gin context", rlog.String("timeout", timeoutstring.(time.Duration).String()))
		timeout = timeoutstring.(time.Duration)
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
	identity, err := getIdentityFromGinContext(c)
	if err != nil {
		rlog.Error("could not get user from gin context: %v", err)
		c.JSON(http.StatusUnauthorized, rorerror.RorError{
			Status:  http.StatusUnauthorized,
			Message: "Could not fetch user",
		})
		return nil, cancel
	}
	ctx = context.WithValue(ctx, identitymodels.ContexIdentity, *identity)
	return ctx, cancel
}

// GetUserFromGinContext Function extracts the user from the gin context
//
// !!! Should only be used in audit middleware !!!
func GetUserFromGinContext(c *gin.Context) (*identitymodels.User, error) {
	userObject, exists := c.Get("user")
	if !exists {
		return nil, errors.New("user not set in gin context")
	}

	if userObject == nil {
		return nil, errors.New("user object is nil")
	}

	user := userObject.(identitymodels.User)

	return &user, nil
}

// Function extracts the identity from gin context
func getIdentityFromGinContext(c *gin.Context) (*identitymodels.Identity, error) {
	identityObj, ok := c.Get("identity")
	if !ok {
		return nil, errors.New("identity not set in gin context")
	}

	if identityObj == nil {
		return nil, errors.New("identity object is nil")
	}

	identity := identityObj.(identitymodels.Identity)
	return &identity, nil
}

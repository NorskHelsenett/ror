package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func DexMiddleware(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		c.String(http.StatusForbidden, "No Authorization header provided ")
		c.Abort()
		return
	}

	skipVerificationCheck := viper.GetBool(configconsts.OIDC_SKIP_ISSUER_VERIFY)

	oicdProvider := viper.GetString(configconsts.OIDC_PROVIDER)

	var provider *oidc.Provider
	var err error

	if !skipVerificationCheck {
		provider, err = oidc.NewProvider(c, oicdProvider)
	} else {
		issuerURL := oicdProvider
		ctx := oidc.InsecureIssuerURLContext(c, issuerURL)
		// Provider will be discovered with the discoveryBaseURL, but use issuerURL
		// for future issuer validation.
		provider, err = oidc.NewProvider(ctx, oicdProvider)
	}

	if err != nil {
		rlog.Error(fmt.Sprintf("Could not get provider, %s", oicdProvider), err)
		c.String(http.StatusForbidden, "Could not get provider, %s", oicdProvider)
		c.Abort()
		return
	}

	token := strings.TrimPrefix(auth, "Bearer ")
	if token == auth {
		c.String(http.StatusForbidden, "Could not find bearer token in Authorization header")
		c.Abort()
		return
	}

	idTokenVerifier := provider.Verifier(&oidc.Config{
		ClientID:                   viper.GetString(configconsts.OIDC_CLIENT_ID),
		SkipIssuerCheck:            skipVerificationCheck,
		InsecureSkipSignatureCheck: skipVerificationCheck,
	})

	flowKind := c.Request.Header.Get("Flow")
	if len(flowKind) > 0 && flowKind == "device" {
		idTokenVerifier = provider.Verifier(&oidc.Config{
			ClientID:                   viper.GetString(configconsts.OIDC_DEVICE_CLIENT_ID),
			SkipIssuerCheck:            skipVerificationCheck,
			InsecureSkipSignatureCheck: skipVerificationCheck,
		})
	}

	// Parse and verify ID Token payload.
	idToken, err := idTokenVerifier.Verify(c, token)
	if err != nil {
		// handle error
		c.String(http.StatusForbidden, "Could not verify token.")
		c.Abort()
		return
	}

	// Extract custom user.
	user := identitymodels.User{Groups: []string{"NotAuthorized"}}
	if err := idToken.Claims(&user); err != nil {
		c.String(http.StatusUnauthorized, "Not authorized")
		c.Abort()
		return
	}

	groupsWithDomain, err := ExtractGroups(&user)
	if err != nil || len(groupsWithDomain) == 0 {
		c.String(http.StatusUnauthorized, "Not authorized, missing groups")
		c.Abort()
		return
	}

	user.Groups = groupsWithDomain

	exptime := time.Unix(int64(user.ExpirationTime), 0)

	c.Set("user", user)
	c.Set("identity", identitymodels.Identity{
		Auth: identitymodels.AuthInfo{
			AuthProvider:   identitymodels.IdentityProviderOidc,
			AuthProviderID: user.Email,
			ExpirationTime: exptime,
		},
		Type: identitymodels.IdentityTypeUser,
		User: &user,
	})
	// Pass on to the next-in-chain
	c.Next()
}

func Contains[T comparable](array []T, element T) bool {
	for _, v := range array {
		if v == element {
			return true
		}
	}
	return false
}

// Function extracts groups from user object
func ExtractGroups(user *identitymodels.User) ([]string, error) {
	if user == nil {
		msg := "user is nil"
		rlog.Debug(msg)
		return make([]string, 0), errors.New(msg)
	}

	emailArray := strings.Split(user.Email, "@")
	if len(emailArray) > 2 {
		msg := "could not extract domain from email"
		rlog.Debug(msg)
		return make([]string, 0), errors.New(msg)
	}

	domain := emailArray[1]
	groups := make([]string, 0)
	for i := 0; i < len(user.Groups); i++ {
		g := fmt.Sprintf("%s@%s", user.Groups[i], domain)
		groups = append(groups, g)
	}

	return groups, nil
}

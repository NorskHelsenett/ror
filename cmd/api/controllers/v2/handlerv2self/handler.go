package handlerv2self

import (
	"net/http"

	"github.com/NorskHelsenett/ror/pkg/context/gincontext"
	"github.com/NorskHelsenett/ror/pkg/context/rorcontext"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/v2/apicontractsv2self"

	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

func init() {
	rlog.Debug("init user controller")
	validate = validator.New()
}

// @Summary	Get self
// @Schemes
// @Description	Get user details
// @Tags			users
// @Accept			application/json
// @Produce		application/json
// @Success		200	{object}	apicontractsv2self.SelfData
// @Failure		403	{string}	Forbidden
// @Failure		401	{string}	Unauthorized
// @Failure		500	{string}	Failure	message
// @Router			/v2/self [get]
// @Security		ApiKey || AccessToken
func GetSelf() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, _ := gincontext.GetRorContextFromGinContext(c)

		identity := rorcontext.GetIdentityFromRorContext(ctx)

		result := apicontractsv2self.SelfData{
			Auth: identity.GetAuthInfo(),
			Type: identity.Type,
		}
		if identity.IsUser() {
			result.User = apicontractsv2self.SelfUser{
				Name:   identity.User.Name,
				Email:  identity.User.Email,
				Groups: identity.User.Groups,
			}
		}
		if identity.IsCluster() {
			result.User = apicontractsv2self.SelfUser{
				Name: identity.ClusterIdentity.Id,
			}
		}
		if identity.IsService() {
			result.User = apicontractsv2self.SelfUser{
				Name: identity.ServiceIdentity.Id,
			}
		}
		c.JSON(http.StatusOK, result)
	}
}

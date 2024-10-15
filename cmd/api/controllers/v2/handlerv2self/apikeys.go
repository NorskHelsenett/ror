package handlerv2self

import (
	"fmt"
	apikeysservice "github.com/NorskHelsenett/ror/cmd/api/services/apikeysService"
	"net/http"

	"github.com/NorskHelsenett/ror/pkg/context/gincontext"
	"github.com/NorskHelsenett/ror/pkg/context/rorcontext"

	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/v2/apicontractsv2self"
	"github.com/NorskHelsenett/ror/pkg/helpers/rorerror"

	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/gin-gonic/gin"
)

// @Summary	Create api key
// @Schemes
// @Description	Create a api key
// @Tags			users
// @Accept			application/json
// @Produce		application/json
// @Success		200						{object}	apicontractsv2self.CreateOrRenewApikeyResponse
// @Failure		403						{object}	rorerror.RorError
// @Failure		401						{object}	rorerror.RorError
// @Failure		500						{object}	rorerror.RorError
// @Router			/v2/self/apikeys	[post]
// @Param			project					body	apicontractsv2self.CreateOrRenewApikeyRequest	true	"Api key"
// @Security		ApiKey || AccessToken
func CreateOrRenewApikey() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input apicontractsv2self.CreateOrRenewApikeyRequest
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		identity := rorcontext.GetIdentityFromRorContext(ctx)
		if identity.Auth.AuthProvider == identitymodels.IdentityProviderApiKey {
			rlog.Error("cannot create apikey with apikey", fmt.Errorf("cannot create apikey with apikey"))
			c.AbortWithError(http.StatusBadRequest, fmt.Errorf("cannot create apikey with apikey"))
		}

		if err := c.BindJSON(&input); err != nil {
			rlog.Errorc(ctx, "could not bind JSON", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Required fields are missing",
			})
			return
		}

		if err := validate.Struct(&input); err != nil {
			rlog.Errorc(ctx, "could not validate object", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Could not validate project object",
			})
			return
		}

		apikeyresponse, err := apikeysservice.CreateOrRenew(ctx, &input)
		if err != nil {
			rlog.Errorc(ctx, "could not create api key", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Unable to create api key, perhaps it already exist?",
			})
			return
		}

		c.JSON(http.StatusOK, apikeyresponse)
	}
}

// @Summary	Delete api key for user
// @Schemes
// @Description	Delete a api key by id for user
// @Tags			user
// @Accept			application/json
// @Produce		application/json
// @Success		200									{bool}		bool
// @Failure		403									{object}	rorerror.RorError
// @Failure		401									{object}	rorerror.RorError
// @Failure		500									{object}	rorerror.RorError
// @Router			/v2/self/apikeys/{apikeyId}	[delete]
// @Param			apikeyId							path	string	true	"apikeyId"
// @Security		ApiKey || AccessToken
func DeleteApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		apikeyId := c.Param("id")
		if apikeyId == "" || len(apikeyId) == 0 {
			rlog.Errorc(ctx, "invalid id", fmt.Errorf("id is zero length"))
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Invalid id",
			})
			return
		}

		identity := rorcontext.GetIdentityFromRorContext(ctx)

		if !identity.IsUser() {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Invalid identity",
			})
			return
		}

		// todo fix delete for user
		result, err := apikeysservice.DeleteForUser(ctx, apikeyId, &identity)
		if err != nil {
			rlog.Errorc(ctx, "could not delete object", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Could not delete object",
			})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

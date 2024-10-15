package users

import (
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/api/customvalidators"
	apikeysservice "github.com/NorskHelsenett/ror/cmd/api/services/apikeysService"
	"net/http"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/context/gincontext"
	"github.com/NorskHelsenett/ror/pkg/context/rorcontext"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/helpers/rorerror"

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
	customvalidators.Setup(validate)
}

// @Summary	Get user
// @Schemes
// @Description	Get user details
// @Tags			users
// @Accept			application/json
// @Produce		application/json
// @Success		200	{object}	apicontracts.User
// @Failure		403	{string}	Forbidden
// @Failure		401	{string}	Unauthorized
// @Failure		500	{string}	Failure	message
// @Router			/v1/users/self [get]
// @Security		ApiKey || AccessToken
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, _ := gincontext.GetRorContextFromGinContext(c)

		identity := rorcontext.GetIdentityFromRorContext(ctx)
		if !identity.IsUser() {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Invalid identity",
			})
			return
		}

		if identity.User == nil {
			c.JSON(http.StatusUnauthorized, nil)
			return
		}

		result := apicontracts.User{
			Name:   identity.User.Name,
			Email:  identity.User.Email,
			Groups: identity.User.Groups,
		}

		c.JSON(http.StatusOK, result)
	}
}

// @Summary	Get apikeys by filter
// @Schemes
// @Description	Get apikeys by filter
// @Tags			users
// @Accept			application/json
// @Produce		application/json
// @Success		200								{object}	apicontracts.PaginatedResult[apicontracts.ApiKey]
// @Failure		403								{object}	rorerror.RorError
// @Failure		401								{object}	rorerror.RorError
// @Failure		500								{object}	rorerror.RorError
// @Router			/v1/users/self/apikeys/filter	[post]
// @Param			filter							body	apicontracts.Filter	true	"Filter"
// @Security		ApiKey || AccessToken
func GetApiKeysByFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		var filter apicontracts.Filter

		//validate the request body
		if err := c.BindJSON(&filter); err != nil {
			rlog.Errorc(ctx, "missing parameter", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Missing parameter",
			})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&filter); validationErr != nil {
			rlog.Errorc(ctx, "failed to validate required fields", validationErr)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: validationErr.Error(),
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

		// importing apicontracts for swagger
		var _ apicontracts.PaginatedResult[apicontracts.Cluster]
		filter.Filters = append(filter.Filters, apicontracts.FilterMetadata{
			Field:     "identifier",
			MatchMode: apicontracts.MatchModeEquals,
			Value:     identity.User.Email,
		})
		paginatedResult, err := apikeysservice.GetByFilter(ctx, &filter)
		if err != nil {
			rlog.Errorc(ctx, "could not get apikeys", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		if paginatedResult == nil {
			empty := apicontracts.PaginatedResult[apicontracts.Cluster]{}
			c.JSON(http.StatusOK, empty)
			return
		}

		c.JSON(http.StatusOK, paginatedResult)
	}
}

// @Summary	Create api key
// @Schemes
// @Description	Create a api key
// @Tags			users
// @Accept			application/json
// @Produce		application/json
// @Success		200						{string}	api	key
// @Failure		403						{object}	rorerror.RorError
// @Failure		401						{object}	rorerror.RorError
// @Failure		500						{object}	rorerror.RorError
// @Router			/v1/users/self/apikeys	[post]
// @Param			project					body	apicontracts.ApiKey	true	"Api key"
// @Security		ApiKey || AccessToken
func CreateApikey() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		identity := rorcontext.GetIdentityFromRorContext(ctx)

		if !identity.IsUser() {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Invalid identity",
			})
			return
		}

		var input apicontracts.ApiKey
		if err := c.BindJSON(&input); err != nil {
			rlog.Errorc(ctx, "could not bind JSON", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Required fields are missing",
			})
			return
		}

		// Ensure that you cant create a api key for another user
		input.Identifier = identity.GetId()

		if err := validate.Struct(&input); err != nil {
			rlog.Errorc(ctx, "could not validate object", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Could not validate project object",
			})
			return
		}

		apikeyText, err := apikeysservice.Create(ctx, &input, &identity)
		if err != nil {
			if strings.Contains(err.Error(), "too many apikeys") {
				c.JSON(http.StatusBadRequest, rorerror.RorError{
					Status:  http.StatusForbidden,
					Message: "Too many apikeys, limit of 100 reached.",
				})
				return
			}

			rlog.Errorc(ctx, "could not create api key", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Unable to create api key, perhaps it already exist?",
			})
			return
		}

		c.JSON(http.StatusOK, apikeyText)
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
// @Router			/v1/users/self/apikeys/{apikeyId}	[delete]
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

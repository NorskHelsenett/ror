// TODO: Describe package
package prices

import (
	"fmt"
	"net/http"
	"strings"

	pricesService "github.com/NorskHelsenett/ror/cmd/api/services/pricesService"
	aclservice "github.com/NorskHelsenett/ror/internal/acl/services"

	"github.com/NorskHelsenett/ror/pkg/context/gincontext"

	aclmodels "github.com/NorskHelsenett/ror/pkg/models/aclmodels"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	rorerror "github.com/NorskHelsenett/ror/pkg/helpers/rorerror"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

func init() {
	rlog.Debug("init cluster controller")
	validate = validator.New()
}

// TODO: Describe function
//
// // @Summary Create a price
// // @Schemes
// // @Description Create a price
// // @Tags prices
// // @Accept application/json
// // @Produce application/json
// // @Param        price  body      apicontracts.Cluster  true  "Add a price"
// // @Success 200 {array} apicontracts.Price
// // @Failure 403  {string}  Forbidden
// // @Failure 401  {string}  Unauthorized
// // @Failure 500  {string}  Failure message
// // @Router	/v1/prices [post]
// // @Security		ApiKey || AccessToken
func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		// Access check
		// Scope: ror
		// Subject: price
		// Access: create
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectPrice)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Create {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		var price apicontracts.Price
		//validate the request body
		if err := c.BindJSON(&price); err != nil {
			rlog.Errorc(ctx, "could not bind JSON", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Could not validate price object",
			})
			return
		}

		//use the validator library to validate required fields
		if err := validate.Struct(&price); err != nil {
			rlog.Errorc(ctx, "could not validate object", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Required fields are missing",
			})
			return
		}

		createdPrice, err := pricesService.Create(ctx, &price)
		if err != nil {
			rlog.Errorc(ctx, "could not create price", err)
			if strings.Contains(err.Error(), "exists") {
				c.JSON(http.StatusBadRequest, rorerror.RorError{
					Status:  http.StatusBadRequest,
					Message: "Already exists",
				})
				return
			}

			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Required fields are missing",
			})
			return
		}

		c.Set("newObject", createdPrice)
		c.JSON(http.StatusOK, createdPrice)
	}
}

// TODO: Describe function
//
//	@Summary	Get prices by provider
//	@Schemes
//	@Description	Get prices by provider
//	@Tags			prices
//	@Accept			application/json
//	@Produce		application/json
//	@Param			providerName	path		string	true	"providerName"
//	@Success		200				{array}		apicontracts.Price
//	@Failure		403				{string}	Forbidden
//	@Failure		401				{string}	Unauthorized
//	@Failure		500				{string}	Failure	message
//	@Router			/v1/prices/provider/{providerName} [get]
//	@Security		ApiKey || AccessToken
func GetByProvider() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		providerName := c.Param("providerName")
		defer cancel()

		if providerName == "" || len(providerName) == 0 {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Invalid provider name",
			})
			return
		}

		prices, err := pricesService.GetByProperty(ctx, "provider", strings.ToLower(providerName))
		if err != nil {
			c.JSON(http.StatusInternalServerError, rorerror.RorError{
				Status:  http.StatusInternalServerError,
				Message: "could not get prices",
			})
		}

		c.JSON(http.StatusOK, prices)
	}
}

// TODO: Describe function
//
//	@Summary	Update a price
//	@Schemes
//	@Description	Update a price by id
//	@Tags			prices
//	@Accept			application/json
//	@Produce		application/json
//	@Param			id		path		string				true	"id"
//	@Param			price	body		apicontracts.Price	true	"Update price"
//	@Success		200		{object}	apicontracts.Price
//	@Failure		403		{string}	Forbidden
//	@Failure		401		{string}	Unauthorized
//	@Failure		500		{string}	Failure	message
//	@Router			/v1/prices/:id [put]
//	@Security		ApiKey || AccessToken
func Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		priceId := c.Param("priceId")
		if priceId == "" || len(priceId) == 0 {
			rlog.Errorc(ctx, "invalid price id", fmt.Errorf("id is zero length"))
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Invalid price id",
			})
			return
		}

		// Access check
		// Scope: ror
		// Subject: price
		// Access: update
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectPrice)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Update {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		var priceInput apicontracts.Price
		//validate the request body
		if err := c.BindJSON(&priceInput); err != nil {
			rlog.Errorc(ctx, "object is not valid", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Object is not valid",
			})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&priceInput); validationErr != nil {
			rlog.Errorc(ctx, "could not validate reqired fields", validationErr)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Required fields missing",
			})
			return
		}

		updatedprice, originalprice, err := pricesService.Update(ctx, priceId, &priceInput)
		if err != nil {
			rlog.Errorc(ctx, "could not update price", err)
			c.JSON(http.StatusInternalServerError, rorerror.RorError{
				Status:  http.StatusInternalServerError,
				Message: "Could not update price",
			})
			return
		}

		if updatedprice == nil {
			rlog.Errorc(ctx, "Could not update price", fmt.Errorf("object does not exist"))
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Could not update price, does it exist?!",
			})
			return
		}

		c.Set("newObject", updatedprice)
		c.Set("oldObject", originalprice)
		c.JSON(http.StatusOK, updatedprice)
	}
}

// TODO: Describe function
//
//	@Summary	Delete a price
//	@Schemes
//	@Description	Delete a price by id
//	@Tags			prices
//	@Accept			application/json
//	@Produce		application/json
//	@Param			id	path		string	true	"id"
//	@Success		200	{bool}		true
//	@Failure		403	{string}	Forbidden
//	@Failure		401	{string}	Unauthorized
//	@Failure		500	{string}	Failure	message
//	@Router			/v1/prices/:id [delete]
//	@Security		ApiKey || AccessToken
func Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		priceId := c.Param("priceId")
		if priceId == "" || len(priceId) == 0 {
			rlog.Errorc(ctx, "invalid price id", fmt.Errorf("id is zero length"))
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Invalid price id",
			})
			return
		}
		// Access check
		// Scope: ror
		// Subject: price
		// Access: delete
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectPrice)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Delete {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		result, deletedPrice, err := pricesService.Delete(ctx, priceId)
		if err != nil {
			rlog.Errorc(ctx, "could not delete price", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Could not delete price",
			})
			return
		}

		c.Set("oldObject", deletedPrice)
		c.JSON(http.StatusOK, result)
	}
}

// TODO: Describe function
//
//	@Summary	Get prices
//	@Schemes
//	@Description	Get all prices
//	@Tags			prices
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200			{array}		apicontracts.Price
//	@Failure		403			{string}	Forbidden
//	@Failure		401			{string}	Unauthorized
//	@Failure		500			{string}	Failure	message
//	@Router			/v1/prices	[get]
//	@Security		ApiKey || AccessToken
func GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		prices, err := pricesService.GetAll(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, rorerror.RorError{
				Status:  http.StatusInternalServerError,
				Message: "Could not find prices ...",
			})
		}

		c.JSON(http.StatusOK, prices)
	}
}

// TODO: Describe function
func GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		priceId := c.Param("priceId")
		if priceId == "" || len(priceId) == 0 {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "invalid price id",
			})
			return
		}

		// Access check
		// Scope: ror
		// Subject: price
		// Access: read
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectPrice)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Read {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		price, err := pricesService.GetById(ctx, priceId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, rorerror.RorError{
				Status:  http.StatusInternalServerError,
				Message: "could not get price",
			})
			return
		}

		c.JSON(http.StatusOK, price)
	}
}

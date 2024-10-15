// TODO: Describe package
package datacenters

import (
	datacentersservice "github.com/NorskHelsenett/ror/cmd/api/services/datacentersService"
	aclservice "github.com/NorskHelsenett/ror/internal/acl/services"
	"net/http"

	"github.com/NorskHelsenett/ror/pkg/context/gincontext"
	"github.com/NorskHelsenett/ror/pkg/context/rorcontext"

	aclmodels "github.com/NorskHelsenett/ror/pkg/models/acl"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/helpers/rorerror"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

func init() {
	validate = validator.New()
}

// TODO: Describe function
//
//	@Summary	Get datacenters
//	@Schemes
//	@Description	Get datacenters
//	@Tags			datacenters
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200				{array}		apicontracts.Datacenter
//	@Failure		403				{string}	Forbidden
//	@Failure		401				{string}	Unauthorized
//	@Failure		500				{string}	Failure	message
//	@Router			/v1/datacenters	[get]
//	@Security		ApiKey || AccessToken
func GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		_, err := gincontext.GetUserFromGinContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, rorerror.RorError{
				Status:  http.StatusUnauthorized,
				Message: "Could not fetch user",
			})
			return
		}

		// importing apicontracts for swagger
		var _ apicontracts.Datacenter

		datacenters, err := datacentersservice.GetAllByUser(ctx)
		if err != nil {
			c.JSON(http.StatusUnauthorized, rorerror.RorError{
				Status:  http.StatusUnauthorized,
				Message: "Could not fetch datacenters",
			})
			return
		}

		c.JSON(http.StatusOK, datacenters)
	}
}

// TODO: Describe function
//
//	@Summary	Get datacenter by name
//	@Schemes
//	@Description	Get datacenter by name
//	@Tags			datacenters
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200									{object}	apicontracts.Datacenter
//	@Failure		403									{string}	Forbidden
//	@Failure		401									{string}	Unauthorized
//	@Failure		500									{string}	Failure	message
//	@Router			/v1/datacenters/{datacenterName}	[get]
//	@Param			datacenterName						path	string	true	"datacenterName"
//	@Security		ApiKey || AccessToken
func GetByName() gin.HandlerFunc {
	// todo scheduled for deletion
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		datacenterName := c.Param("datacenterName")
		defer cancel()

		// importing apicontracts for swagger
		var _ apicontracts.Datacenter

		datacenter, err := datacentersservice.GetByName(ctx, datacenterName)
		if err != nil {
			c.JSON(http.StatusUnauthorized, rorerror.RorError{
				Status:  http.StatusUnauthorized,
				Message: "Could not fetch datacenter",
			})
			return
		}

		if datacenter == nil {
			c.JSON(http.StatusNotFound, nil)
			return
		}

		c.JSON(http.StatusOK, datacenter)
	}
}

// @Summary	Get datacenter by id
// @Schemes
// @Description	Get datacenter by id
// @Tags			datacenters
// @Accept			application/json
// @Produce		application/json
// @Success		200									{object}	apicontracts.Datacenter
// @Failure		403									{string}	Forbidden
// @Failure		401									{string}	Unauthorized
// @Failure		500									{string}	Failure	message
// @Router			/v1/datacenters/id/{id}	[get]
// @Param			id						path	string	true	"id"
// @Security		ApiKey || AccessToken
func GetById() gin.HandlerFunc {
	// todo scheduled for deletion
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		datacenterId := c.Param("id")
		defer cancel()

		// importing apicontracts for swagger
		var _ apicontracts.Datacenter

		datacenter, err := datacentersservice.GetById(ctx, datacenterId)
		if err != nil {
			c.JSON(http.StatusUnauthorized, rorerror.RorError{
				Status:  http.StatusUnauthorized,
				Message: "Could not fetch datacenter",
			})
			return
		}

		if datacenter == nil {
			c.JSON(http.StatusNotFound, nil)
			return
		}

		c.JSON(http.StatusOK, datacenter)
	}
}

// TODO: Describe function
//
//	@Summary	Create datacenter
//	@Schemes
//	@Description	Create datacenter
//	@Tags			datacenters
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200				{array}		apicontracts.Datacenter
//	@Failure		403				{string}	Forbidden
//	@Failure		401				{string}	Unauthorized
//	@Failure		500				{string}	Failure					message
//	@Param			datacenter		body		apicontracts.Datacenter	true	"Datacenter"
//	@Router			/v1/datacenters	[post]
//	@Security		ApiKey || AccessToken
func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		identity := rorcontext.GetIdentityFromRorContext(ctx)

		// Access check
		// Scope: ror
		// Subject: datacenter
		// Access: create
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectDatacenter)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Create {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		var datacenterInput apicontracts.DatacenterModel

		//validate the request body
		if err := c.BindJSON(&datacenterInput); err != nil {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Missing body",
			})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&datacenterInput); validationErr != nil {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: validationErr.Error(),
			})
			return
		}

		datacenter, err := datacentersservice.Create(ctx, &datacenterInput, identity.User)
		if err != nil {
			c.JSON(http.StatusInternalServerError, rorerror.RorError{
				Status:  http.StatusInternalServerError,
				Message: "Could not create datacenter",
			})
			return
		}

		if datacenter == nil {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Could not create datacenter, does it already exists?! ",
			})
			return
		}

		c.JSON(http.StatusOK, datacenter)
	}
}

// TODO: Describe function
//
//	@Summary	Update a datacenter
//	@Schemes
//	@Description	Update a datacenter by id
//	@Tags			datacenters
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200								{array}		apicontracts.Datacenter
//	@Failure		403								{string}	Forbidden
//	@Failure		401								{string}	Unauthorized
//	@Failure		500								{string}	Failure					message
//	@Param			datacenterId					path		string					true	"datacenterId"
//	@Param			datacenter						body		apicontracts.Datacenter	true	"Datacenter"
//	@Router			/v1/datacenters/{datacenterId}	[put]
//	@Security		ApiKey || AccessToken
func Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		datacenterId := c.Param("datacenterId")
		defer cancel()

		identity := rorcontext.GetIdentityFromRorContext(ctx)

		// Access check
		// Scope: ror
		// Subject: datacenter
		// Access: update
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectDatacenter)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Update {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		var datacenterInput apicontracts.DatacenterModel

		//validate the request body
		if err := c.BindJSON(&datacenterInput); err != nil {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Missing body",
			})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&datacenterInput); validationErr != nil {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: validationErr.Error(),
			})
			return
		}

		datacenter, err := datacentersservice.Update(ctx, datacenterId, &datacenterInput, identity.User)
		if err != nil {
			c.JSON(http.StatusInternalServerError, rorerror.RorError{
				Status:  http.StatusInternalServerError,
				Message: "Could not update datacenter",
			})
			return
		}

		if datacenter == nil {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Could not update datacenter, does it exists?!",
			})
			return
		}

		c.JSON(http.StatusOK, datacenter)
	}
}

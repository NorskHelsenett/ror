// TODO: Describe package
package operatorconfigs

import (
	"fmt"
	"net/http"
	"strings"

	operatorconfigservice "github.com/NorskHelsenett/ror/cmd/api/services/operatorConfigService"
	aclservice "github.com/NorskHelsenett/ror/internal/acl/services"

	"github.com/NorskHelsenett/ror/pkg/context/gincontext"

	aclmodels "github.com/NorskHelsenett/ror/pkg/models/aclmodels"

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
	rlog.Debug("init operator config controller")
	validate = validator.New()
}

// TODO: Describe function
//
//	@Summary	Get a operator config
//	@Schemes
//	@Description	Get a operator config by id
//	@Tags			operatorconfigs
//	@Accept			application/json
//	@Produce		application/json
//	@Param			id				path		string						true	"id"
//	@Param			operatorconfig	body		apicontracts.OperatorConfig	true	"Get a operator config"
//	@Success		200				{object}	apicontracts.OperatorConfig
//	@Failure		403				{string}	Forbidden
//	@Failure		401				{string}	Unauthorized
//	@Failure		500				{string}	Failure	message
//	@Router			/v1/operatorconfigs/:id [get]
//	@Security		ApiKey || AccessToken
func GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()
		// Access check
		// Scope: ror
		// Subject: global
		// Access: read
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectGlobal)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Read {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		id := c.Param("id")
		if id == "" || len(id) == 0 {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "invalid id",
			})
			return
		}

		result, err := operatorconfigservice.GetById(ctx, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, rorerror.RorError{
				Status:  http.StatusInternalServerError,
				Message: "could not get operator config",
			})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

// TODO: Describe function
//
//	@Summary	Get all operator configs
//	@Schemes
//	@Description	Get all operator configs
//	@Tags			operatorconfigs
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200					{array}		apicontracts.OperatorConfig
//	@Failure		403					{string}	Forbidden
//	@Failure		401					{string}	Unauthorized
//	@Failure		500					{string}	Failure	message
//	@Router			/v1/operatorconfigs	[get]
//	@Security		ApiKey || AccessToken
func GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		// Access check
		// Scope: ror
		// Subject: global
		// Access: read
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectGlobal)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Read {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		elements, err := operatorconfigservice.GetAll(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, rorerror.RorError{
				Status:  http.StatusInternalServerError,
				Message: "Could not find operator configs ...",
			})
		}

		c.JSON(http.StatusOK, elements)
	}
}

// TODO: Describe function
//
//	@Summary	Create a operator config
//	@Schemes
//	@Description	Create a operator config
//	@Tags			operatorconfigs
//	@Accept			application/json
//	@Produce		application/json
//	@Param			operatorconfig	body		apicontracts.OperatorConfig	true	"Add a operator config"
//	@Success		200				{array}		apicontracts.OperatorConfig
//	@Failure		403				{string}	Forbidden
//	@Failure		401				{string}	Unauthorized
//	@Failure		500				{string}	Failure	message
//	@Router			/v1/operatorconfigs [post]
//	@Security		ApiKey || AccessToken
func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()
		// Access check
		// Scope: ror
		// Subject: global
		// Access: create
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectGlobal)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Create {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		var config apicontracts.OperatorConfig
		//validate the request body
		if err := c.BindJSON(&config); err != nil {
			rlog.Errorc(ctx, "could not bind JSON", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Could not validate operator config input",
			})
			return
		}

		//use the validator library to validate required fields
		if err := validate.Struct(&config); err != nil {
			rlog.Errorc(ctx, "could not validate input", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: fmt.Sprintf("Required fields are missing: %s", err),
			})
			return
		}

		created, err := operatorconfigservice.Create(ctx, &config)
		if err != nil {
			rlog.Errorc(ctx, "could not create operator config", err)
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

		c.Set("newObject", created)

		c.JSON(http.StatusOK, created)
	}
}

// TODO: Describe function
//
//	@Summary	Update a operator config
//	@Schemes
//	@Description	Update a operator config by id
//	@Tags			operatorconfigs
//	@Accept			application/json
//	@Produce		application/json
//	@Param			id				path		string						true	"id"
//	@Param			operatorconfig	body		apicontracts.OperatorConfig	true	"Update operator config"
//	@Success		200				{object}	apicontracts.OperatorConfig
//	@Failure		403				{string}	Forbidden
//	@Failure		401				{string}	Unauthorized
//	@Failure		500				{string}	Failure	message
//	@Router			/v1/operatorconfigs/:id [put]
//	@Security		ApiKey || AccessToken
func Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		id := c.Param("id")
		if id == "" || len(id) == 0 {
			rlog.Errorc(ctx, "invalid operator config id", fmt.Errorf("id is zero length"))
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Invalid operator config id",
			})
			return
		}
		// Access check
		// Scope: ror
		// Subject: global
		// Access: update
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectGlobal)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Update {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		var input apicontracts.OperatorConfig

		//validate the request body
		if err := c.BindJSON(&input); err != nil {
			rlog.Errorc(ctx, "input is not valid", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Input is not valid",
			})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&input); validationErr != nil {
			rlog.Errorc(ctx, "could not validate required fields", validationErr)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Required fields missing",
			})
			return
		}

		updated, original, err := operatorconfigservice.Update(ctx, id, &input)
		if err != nil {
			rlog.Errorc(ctx, "could not update operator config", err)
			c.JSON(http.StatusInternalServerError, rorerror.RorError{
				Status:  http.StatusInternalServerError,
				Message: "Could not update operator config",
			})
			return
		}

		if updated == nil {
			rlog.Errorc(ctx, "Could not update operator config", fmt.Errorf("object does not exist"))
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Could not update operator config, does it exist?!",
			})
			return
		}

		c.Set("newObject", updated)
		c.Set("oldObject", original)

		c.JSON(http.StatusOK, updated)
	}
}

// TODO: Describe function
//
//	@Summary	Delete a operator config
//	@Schemes
//	@Description	Delete a operator config by id
//	@Tags			operatorconfigs
//	@Accept			application/json
//	@Produce		application/json
//	@Param			id	path		string	true	"id"
//	@Success		200	{bool}		true
//	@Failure		403	{string}	Forbidden
//	@Failure		401	{string}	Unauthorized
//	@Failure		500	{string}	Failure	message
//	@Router			/v1/operatorconfigs/:id [delete]
//	@Security		ApiKey || AccessToken
func Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		id := c.Param("id")
		if id == "" || len(id) == 0 {
			rlog.Errorc(ctx, "invalid id", fmt.Errorf("id is zero length"))
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Invalid id",
			})
			return
		}
		// Access check
		// Scope: ror
		// Subject: global
		// Access: delete
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectGlobal)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Delete {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		result, err := operatorconfigservice.Delete(ctx, id)
		if err != nil {
			rlog.Errorc(ctx, "could not delete operator config", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Could not delete operator config",
			})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

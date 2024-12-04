// TODO: Describe package
package desired_version

import (
	"net/http"

	desiredversionservice "github.com/NorskHelsenett/ror/cmd/api/services/desiredversionService"
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

// TODO: Describe function
//
//	@Summary	Get a desired version by its name
//	@Schemes
//	@Description	Get a desired version by its name
//	@Tags			desired_version
//	@Accept			application/json
//	@Produce		application/json
//	@Param			key	path		string	true	"key"
//	@Success		200	{object}	apicontracts.DesiredVersion
//	@Failure		500	{string}	Failure	message
//	@Router			/v1/desired_versions/{key} [get]
func GetByKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		key := c.Param("key")
		if key == "" {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "invalid desired version key",
			})
			return
		}

		desiredversion, err := desiredversionservice.GetByKey(ctx, key)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "failed")
			return
		}
		c.JSON(http.StatusOK, desiredversion)
	}
}

// TODO: Describe function
//
//	@Summary	Get all desired versions
//	@Schemes
//	@Description	Get all desired versions
//	@Tags			desired_version
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200	{array}		[]apicontracts.DesiredVersion
//	@Failure		500	{string}	Failure	message
//	@Router			/v1/desired_versions [get]
func GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		desiredversions, err := desiredversionservice.GetAll(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "failed")
		}
		c.JSON(http.StatusOK, desiredversions)
	}
}

// TODO: Describe function
//
//	@Summary	Create a desired version
//	@Schemes
//	@Description	Create a desired version
//	@Tags			desired_version
//	@Accept			application/json
//	@Produce		application/json
//	@Param			version	body		apicontracts.DesiredVersion	true	"Add a desired version"
//	@Success		200		{string}	Ok
//	@Failure		403		{string}	Forbidden
//	@Failure		401		{string}	Unauthorized
//	@Failure		500		{string}	Failure	message
//	@Router			/v1/desired_versions [post]
//	@Security		ApiKey || AccessToken
func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()
		// Access check
		// Scope: ror
		// Subject: cluster
		// Access: create
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectCluster)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)

		//accessObject := aclservice.CheckAccessByContextScopeSubject(ctx, aclmodels.Acl2ScopeRor, aclmodels.Acl2Subject(aclmodels.Acl2RorSubjectAcl))
		if !accessObject.Create {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		var desiredversion apicontracts.DesiredVersion
		//validate the request body
		if err := c.BindJSON(&desiredversion); err != nil {
			rlog.Errorc(ctx, "could not bind JSON", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Could not validate desired version object",
			})
			return
		}

		err := validate.Struct(&desiredversion)
		if err != nil {
			rlog.Errorc(ctx, "could not validate the request body", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Required fields missing",
			})
			return
		}

		creation, err := desiredversionservice.Create(ctx, desiredversion)
		if err != nil {
			rlog.Errorc(ctx, "could not create desired version", err)
			c.JSON(http.StatusInternalServerError, "failed to create desired version")
			return
		}

		c.JSON(http.StatusOK, creation)
	}
}

// TODO: Describe function
//
//	@Summary	Update a desired version by it's key
//	@Schemes
//	@Description	Update a desired version by it's key
//	@Tags			desired_version
//	@Accept			application/json
//	@Produce		application/json
//	@Param			key		path		string						true	"key"
//	@Param			version	body		apicontracts.DesiredVersion	true	"Update the desired version"
//	@Success		200		{string}	Ok
//	@Failure		403		{string}	Forbidden
//	@Failure		401		{string}	Unauthorized
//	@Failure		500		{string}	Failure	message
//	@Router			/v1/desired_versions/{key} [put]
//	@Security		ApiKey || AccessToken
func Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()
		// Access check
		// Scope: ror
		// Subject: cluster
		// Access: update
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectCluster)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Update {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		var desiredversion apicontracts.DesiredVersion
		if err := c.BindJSON(&desiredversion); err != nil {
			rlog.Errorc(ctx, "could not bind JSON", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Could not validate desired version object",
			})
			return
		}

		err := validate.Struct(&desiredversion)
		if err != nil {
			rlog.Errorc(ctx, "could not validate the request body", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Required fields missing",
			})
			return
		}

		key := c.Param("key")
		if key == "" {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "invalid desired version key",
			})
			return
		}

		update, err := desiredversionservice.UpdateByKey(ctx, key, desiredversion)
		if err != nil {
			rlog.Errorc(ctx, "could not update desired version", err)
			c.JSON(http.StatusInternalServerError, "failed to update desired version")
			return
		}

		c.JSON(http.StatusOK, update)
	}
}

// TODO: Describe function
//
//	@Summary	Delete a desired version by it's key
//	@Schemes
//	@Description	Delete a desired version by it's key
//	@Tags			desired_version
//	@Accept			application/json
//	@Produce		application/json
//	@Param			key	path		string	true	"key"
//	@Success		200	{string}	Ok
//	@Failure		403	{string}	Forbidden
//	@Failure		401	{string}	Unauthorized
//	@Failure		500	{string}	Failure	message
//	@Router			/v1/desired_versions/{key} [delete]
//	@Security		ApiKey || AccessToken
func Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		// Access check
		// Scope: ror
		// Subject: cluster
		// Access: delete
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectCluster)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Delete {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		key := c.Param("key")
		if key == "" {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "invalid desired version key",
			})
			return
		}

		_, err := desiredversionservice.DeleteByKey(ctx, key)
		if err != nil {
			rlog.Errorc(ctx, "could not delete desired version", err)
			c.JSON(http.StatusInternalServerError, "failed to delete desired version")
			return
		}

		c.JSON(http.StatusOK, "deleted")
	}
}

// The acl controller package provides controller functions for the /acl endpoints in the api V1.
package acl

import (
	"github.com/NorskHelsenett/ror/cmd/api/apiconnections"
	"github.com/NorskHelsenett/ror/cmd/api/customvalidators"
	aclservice "github.com/NorskHelsenett/ror/internal/acl/services"
	"net/http"

	"github.com/NorskHelsenett/ror/pkg/context/gincontext"
	"github.com/NorskHelsenett/ror/pkg/context/rorcontext"

	"github.com/NorskHelsenett/ror/pkg/messagebuscontracts"

	"github.com/NorskHelsenett/ror/pkg/helpers/rorerror"
	aclmodels "github.com/NorskHelsenett/ror/pkg/models/acl"
	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

func init() {
	rlog.Debug("init cluster controller")
	validate = validator.New()
	customvalidators.Setup(validate)
}

// GetScopes provides a array of aclmodels.Acl2Scope
//
//	@Summary	Get acl scopes
//	@Schemes
//	@Description	Get acl scopes
//	@Tags			acl
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200				{array}		aclmodels.Acl2Scope
//	@Failure		403				{object}	rorerror.RorError
//	@Failure		401				{object}	rorerror.RorError
//	@Failure		500				{object}	rorerror.RorError
//	@Router			/v1/acl/scopes	[get]
//	@Security		ApiKey || AccessToken
func GetScopes() gin.HandlerFunc {
	return func(c *gin.Context) {

		results := aclmodels.GetScopes()
		c.JSON(http.StatusOK, results)
	}
}

// TODO: Describe
//
//	@Summary	Check acl
//	@Schemes
//	@Description	Check acl by scope, subject and access method
//	@Tags			acl
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200
//	@Failure		403
//	@Failure		401									{object}	rorerror.RorError
//	@Failure		500									{object}	rorerror.RorError
//	@Param			scope								path		string	false	"Scope"
//	@Param			subject								path		string	false	"Subject"
//	@Param			access								path		string	false	"read,write,update or delete"
//	@Router			/v1/acl/{scope}/{subject}/{access}	[head]
//	@Security		ApiKey || AccessToken
func CheckAcl() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		scope := c.Param("scope")
		if scope == "" || len(scope) == 0 {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "invalid scope",
			})
			return
		}

		subject := c.Param("subject")
		if subject == "" || len(subject) == 0 {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "invalid subject",
			})
			return
		}

		access := c.Param("access")
		if access == "" || len(access) == 0 {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "invalid id",
			})
			return
		}

		// Check access
		// Scope: c.Param("scope")
		// Subject: c.Param("subject")
		// Access: c.Param("access")
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(scope, subject)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		switch access {
		case "read":
			if accessObject.Read {
				c.Status(http.StatusOK)
				return
			}
		case "create":
			if accessObject.Create {
				c.Status(http.StatusOK)
				return
			}
		case "update":
			if accessObject.Update {
				c.Status(http.StatusOK)
				return
			}
		case "delete":
			if accessObject.Delete {
				c.Status(http.StatusOK)
				return
			}
		case "owner":
			if accessObject.Owner {
				c.Status(http.StatusOK)
				return
			}
		default:
			c.Status(http.StatusForbidden)
			return
		}

		c.Status(http.StatusForbidden)
	}
}

// TODO: Describe
//
//	@Summary	Get acl by id
//	@Schemes
//	@Description	Get acl by id
//	@Tags			acl
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200				{object}	apicontracts.PaginatedResult[aclmodels.AclV2ListItem]
//	@Failure		403				{object}	rorerror.RorError
//	@Failure		401				{object}	rorerror.RorError
//	@Failure		500				{object}	rorerror.RorError
//	@Router			/v1/acl/{aclId}	[get]
//	@Param			id				path	string	true	"id"
//	@Security		ApiKey || AccessToken
func GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		// Check access
		// Scope: Ror
		// Subject: Acl
		// Access: Read
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2Subject(aclmodels.Acl2RorSubjectAcl))
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Read {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		aclId := c.Param("id")
		if aclId == "" || len(aclId) == 0 {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "invalid id",
			})
			return
		}

		var _ aclmodels.AclV2ListItem
		object, err := aclservice.GetById(ctx, aclId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, rorerror.RorError{
				Status:  http.StatusInternalServerError,
				Message: "could not get object",
			})
			return
		}

		c.JSON(http.StatusOK, object)
	}
}

// TODO: Describe
//
//	@Summary	Get acl by filter
//	@Schemes
//	@Description	Get acl by filter
//	@Tags			acl
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200				{object}	apicontracts.PaginatedResult[aclmodels.AclV2ListItem]
//	@Failure		403				{object}	rorerror.RorError
//	@Failure		401				{object}	rorerror.RorError
//	@Failure		500				{object}	rorerror.RorError
//	@Router			/v1/acl/filter	[post]
//	@Param			filter			body	apicontracts.Filter	true	"Filter"
//	@Security		ApiKey || AccessToken
func GetByFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		var filter apicontracts.Filter

		// Check access
		// Scope: Ror
		// Subject: Acl
		// Access: Read
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2Subject(aclmodels.Acl2RorSubjectAcl))
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)

		if !accessObject.Read {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		//validate the request body
		if err := c.BindJSON(&filter); err != nil {
			rlog.Errorc(ctx, "Missing parameter", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Missing parameter",
			})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&filter); validationErr != nil {
			rlog.Errorc(ctx, validationErr.Error(), validationErr)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: validationErr.Error(),
			})
			return
		}

		// importing apicontracts for swagger
		var _ apicontracts.PaginatedResult[aclmodels.AclV2ListItem]
		paginatedResult, err := aclservice.GetByFilter(ctx, &filter)
		if err != nil {
			rlog.Errorc(ctx, err.Error(), err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		if paginatedResult == nil {
			empty := apicontracts.PaginatedResult[aclmodels.AclV2ListItem]{}
			c.JSON(http.StatusOK, empty)
			return
		}

		c.JSON(http.StatusOK, paginatedResult)
	}
}

// TODO: Describe
//
//	@Summary	Create acl
//	@Schemes
//	@Description	Create acl
//	@Tags			acl
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200		{object}	aclmodels.AclV2ListItem
//	@Failure		403		{object}	rorerror.RorError
//	@Failure		401		{object}	rorerror.RorError
//	@Failure		500		{object}	rorerror.RorError
//	@Router			/v1/acl	[post]
//	@Param			acl		body	aclmodels.AclV2ListItem	true	"Acl"
//	@Security		ApiKey || AccessToken
func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		identity := rorcontext.GetIdentityFromRorContext(ctx)

		// Check access
		// Scope: Ror
		// Subject: Acl
		// Access: Create
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2Subject(aclmodels.Acl2RorSubjectAcl))
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Create {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		var aclModel aclmodels.AclV2ListItem
		if err := c.BindJSON(&aclModel); err != nil {
			rlog.Errorc(ctx, "could not bind JSON", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Required fields are missing",
			})
			return
		}

		if err := validate.Struct(&aclModel); err != nil {
			rlog.Errorc(ctx, "could not validate object", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Could not validate object",
			})
			return
		}

		created, err := aclservice.Create(ctx, &aclModel, &identity)
		if err != nil {
			rlog.Errorc(ctx, "could not create", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Unable to create",
			})
			return
		}

		_ = apiconnections.RabbitMQConnection.SendMessage(ctx, messagebuscontracts.AclUpdateEvent{Action: "Create"}, messagebuscontracts.Route_Acl_Update, nil)
		c.JSON(http.StatusOK, created)
	}
}

// TODO: Describe
//
//	@Summary	Update acl
//	@Schemes
//	@Description	Update acl
//	@Tags			acl
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200				{object}	aclmodels.AclV2ListItem
//	@Failure		403				{object}	rorerror.RorError
//	@Failure		401				{object}	rorerror.RorError
//	@Failure		500				{object}	rorerror.RorError
//	@Router			/v1/acl/{aclId}	[put]
//	@Param			aclId			path	string					true	"aclId"
//	@Param			acl				body	aclmodels.AclV2ListItem	true	"Acl"
//	@Security		ApiKey || AccessToken
func Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		identity := rorcontext.GetIdentityFromRorContext(ctx)

		// Check access
		// Scope: Ror
		// Subject: Acl
		// Access: Update
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2Subject(aclmodels.Acl2RorSubjectAcl))
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Update {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		aclId := c.Param("id")
		if aclId == "" || len(aclId) == 0 {
			rlog.Errorc(ctx, "invalid id", nil)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Invalid id",
			})
			return
		}

		var aclModel aclmodels.AclV2ListItem
		if err := c.BindJSON(&aclModel); err != nil {
			rlog.Errorc(ctx, "could not bind JSON", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Required fields are missing",
			})
			return
		}

		if err := validate.Struct(&aclModel); err != nil {
			rlog.Errorc(ctx, "could not validate object", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Could not validate object",
			})
			return
		}

		created, err := aclservice.Update(ctx, aclId, &aclModel, &identity)
		if err != nil {
			rlog.Errorc(ctx, "could not update", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Unable to update",
			})
			return
		}

		payload := messagebuscontracts.AclUpdateEvent{Action: "Update"}
		_ = apiconnections.RabbitMQConnection.SendMessage(ctx, payload, messagebuscontracts.Route_Acl_Update, nil)

		c.JSON(http.StatusOK, created)
	}
}

// TODO: Describe
//
//	@Summary		Delete acl
//	@Description	Delete a acl by id
//	@Tags			acl
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200				{bool}		bool
//	@Failure		403				{object}	rorerror.RorError
//	@Failure		401				{object}	rorerror.RorError
//	@Failure		500				{object}	rorerror.RorError
//	@Router			/v1/acl/{aclId}	[delete]
//	@Param			aclId			path	string	true	"aclId"
//	@Security		ApiKey || AccessToken
func Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		identity := rorcontext.GetIdentityFromRorContext(ctx)
		aclId := c.Param("id")
		if aclId == "" || len(aclId) == 0 {
			rlog.Errorc(ctx, "invalid id", nil)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Invalid id",
			})
			return
		}

		// Check access
		// Scope: Ror
		// Subject: Acl
		// Access: Delete
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2Subject(aclmodels.Acl2RorSubjectAcl))
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Delete {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		result, _, err := aclservice.Delete(ctx, aclId, &identity)
		if err != nil {
			rlog.Errorc(ctx, "could not delete object", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Could not delete object",
			})
			return
		}

		payload := messagebuscontracts.AclUpdateEvent{Action: "Delete"}
		_ = apiconnections.RabbitMQConnection.SendMessage(ctx, payload, messagebuscontracts.Route_Acl_Update, nil)

		c.JSON(http.StatusOK, result)
	}
}

// TODO: Describe
//
//	@Summary	Migrate acl
//	@Schemes
//	@Description	Migrate acl
//	@Tags			acl
//	@Accept			application/json
//	@Produce		application/json
//	@Param			id	path		string	true	"id"
//	@Success		200	{string}	Status
//	@Failure		403	{string}	Forbidden
//	@Failure		401	{string}	Unauthorized
//	@Failure		500	{string}	Failure	message
//	@Router			/v1/acl/migrate [get]
//	@Security		ApiKey || AccessToken
func MigrateAcls() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		// Check access
		// Scope: Ror
		// Subject: Acl
		// Access: Update
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2Subject(aclmodels.Acl2RorSubjectAcl))
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Update {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		err := aclservice.MigrateAcl1toAcl2(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"success": false, "message": "error migrating ACL"})
			return
		}

		payload := messagebuscontracts.AclUpdateEvent{Action: "Migrate"}
		_ = apiconnections.RabbitMQConnection.SendMessage(ctx, payload, messagebuscontracts.Route_Acl_Update, nil)

		c.JSON(http.StatusOK, map[string]interface{}{"success": true, "message": "ACL migrated"})
	}
}

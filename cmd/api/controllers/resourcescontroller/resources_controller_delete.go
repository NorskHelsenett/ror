package resourcescontroller

import (
	"net/http"

	"github.com/NorskHelsenett/ror/cmd/api/responses"
	resourcesservice "github.com/NorskHelsenett/ror/cmd/api/services/resourcesService"
	aclservice "github.com/NorskHelsenett/ror/internal/acl/services"

	"github.com/NorskHelsenett/ror/pkg/context/gincontext"

	aclmodels "github.com/NorskHelsenett/ror/pkg/models/aclmodels"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"

	"github.com/gin-gonic/gin"
)

// Delete a cluster resource of given group/version/kind by uid.
//
//	@Summary	Delete a resource by uid
//	@Schemes
//	@Description	Delete a resources
//	@Tags			resources
//	@Accept			application/json
//	@Produce		application/json
//	@Param			uid	path		string	true	"UID"
//	@Success		200	{bool}		bool
//	@Failure		403	{string}	Forbidden
//	@Failure		401	{string}	Unauthorized
//	@Failure		500	{string}	Failure	message
//	@Router			/v1/resources/uid/{uid} [delete]
//	@Security		ApiKey || AccessToken
func DeleteResource() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()
		var input apiresourcecontracts.ResourceUpdateModel

		//validate the request body
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.Cluster{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		//use the validator library to validate required fields
		if validationErr := validate.Struct(&input); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.Cluster{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		// Validate that the correct uid is provided
		if input.Uid != c.Param("uid") {
			c.JSON(http.StatusNotImplemented, "501: Wrong uid")
			return
		}

		scope := aclmodels.Acl2Scope(input.Owner.Scope)
		subject := input.Owner.Subject

		if subject == "" || scope == "" {
			c.JSON(http.StatusBadRequest, responses.Cluster{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "owner scope and subject must be set"}})
			return
		}

		// Access check
		// Scope: input.Owner.Scope
		// Subject: input.Owner.Subject
		// Access: update
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(scope, subject)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Update {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		err := resourcesservice.ResourceDeleteService(ctx, input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Cluster{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, nil)

	}
}

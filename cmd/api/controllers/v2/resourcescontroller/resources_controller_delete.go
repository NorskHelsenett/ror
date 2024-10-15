package resourcescontroller

import (
	"github.com/NorskHelsenett/ror/cmd/api/responses"
	"github.com/NorskHelsenett/ror/cmd/api/services/resourcesv2service"
	aclservice "github.com/NorskHelsenett/ror/internal/acl/services"
	"net/http"

	"github.com/NorskHelsenett/ror/pkg/context/gincontext"
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
//	@Router			/v2/resources/uid/{uid} [delete]
//	@Security		ApiKey || AccessToken
func DeleteResource() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		resources := resourcesv2service.GetResourceByUID(ctx, c.Param("uid"))

		if resources == nil {
			c.JSON(http.StatusNotFound, "404: Resource not found")
			return
		}

		// Validate that the correct uid is provided
		if len(resources.Resources) != 1 {
			c.JSON(http.StatusNotImplemented, "501: Wrong number of resources found")
			return
		}

		resource := resources.Resources[0]

		if c.Param("uid") != resource.GetUID() {
			c.JSON(http.StatusNotImplemented, "501: Wrong resource found")
			return
		}

		// Access check
		// Scope: input.Owner.Scope
		// Subject: input.Owner.Subject
		// Access: update

		accessModel := aclservice.CheckAccessByRorOwnerref(ctx, resource.GetRorMeta().Ownerref)
		if !accessModel.Update {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		err := resourcesv2service.DeleteResource(ctx, resource)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Cluster{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, nil)

	}
}

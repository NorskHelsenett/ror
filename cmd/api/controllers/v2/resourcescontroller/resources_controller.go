// resourcecontroller implements all controllers for resources
package resourcescontroller

import (
	"net/http"

	"github.com/NorskHelsenett/ror/cmd/api/customvalidators"
	"github.com/NorskHelsenett/ror/cmd/api/responses"
	resourcesservice "github.com/NorskHelsenett/ror/cmd/api/services/resourcesService"
	"github.com/NorskHelsenett/ror/cmd/api/services/resourcesv2service"

	aclservice "github.com/NorskHelsenett/ror/internal/acl/services"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/context/gincontext"
	aclmodels "github.com/NorskHelsenett/ror/pkg/models/acl"
	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

// Init is called to initialize the resources controller
func init() {
	rlog.Debug("init resources controller")
	validate = validator.New()
	customvalidators.Setup(validate)
}

// Check if a cluster resource of given uid exists.
//
//	@Summary	Check cluster resource by uid
//	@Schemes
//	@Description	Get a list of cluster resources
//	@Tags			resources
//	@Accept			application/json
//	@Produce		application/json
//	@Param			ownerScope		query	aclmodels.Acl2Scope	true	"The kind of the owner, currently only support 'Cluster'"
//	@Param			ownerSubject	query	string							true	"The name og the owner"
//	@Param			uid				path	string							true	"UID"
//	@Success		204
//	@Failure		404	{string}	NotFound
//	@Failure		401	{string}	Unauthorized
//	@Failure		500	{string}	Failure	message
//	@Router			/v2/resources/uid/{uid} [head]
//	@Security		ApiKey || AccessToken
func ExistsResources() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		if c.Param("uid") == "" {
			c.JSON(http.StatusBadRequest, "empty uid")
			return
		}

		resources := resourcesv2service.GetResourceByUID(ctx, c.Param("uid"))
		if resources == nil {
			c.Status(http.StatusNotFound)
			return
		}

		// Validate that the correct uid is provided
		if len(resources.Resources) != 1 {
			c.JSON(http.StatusNotImplemented, "501: Wrong number of resources found")
			return
		}

		resource := resources.Resources[0]

		if c.Param("uid") != resource.GetUID() {
			c.JSON(http.StatusBadRequest, "400: Wrong resource found")
			return
		}

		// Access check
		// Scope: input.Owner.Scope
		// Subject: input.Owner.Subject
		// Access: Read
		accessModel := aclservice.CheckAccessByRorOwnerref(ctx, resource.GetRorMeta().Ownerref)
		if !accessModel.Read {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		if resources.Len() == 1 {
			c.Status(http.StatusNoContent)
			return
		} else {
			c.Status(http.StatusNotFound)
			return
		}
	}
}

// Get a list of hashes of saved resources for given cluster.
// Parameter clusterid must match authorized clusterid
//
//	@Summary	Get resource hash list
//	@Schemes
//	@Description	Get a resource list
//	@Tags			resources
//	@Accept			application/json
//	@Produce		application/json
//	@Param			ownerScope		query		aclmodels.Acl2Scope	true	"The kind of the owner, currently only support 'Cluster'"
//	@Param			ownerSubject	query		string							true	"The name og the owner"
//	@Success		200				{array}		apiresourcecontracts.HashList
//	@Failure		403				{string}	Forbidden
//	@Failure		401				{string}	Unauthorized
//	@Failure		500				{string}	Failure	message
//	@Router			/v2/resources/hashes [get]
//	@Security		ApiKey || AccessToken
func GetResourceHashList() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()
		resourceOwner := apiresourcecontracts.ResourceOwnerReference{
			Scope:   aclmodels.Acl2Scope(c.Query("ownerScope")),
			Subject: c.Query("ownerSubject"),
		}

		// Access check
		// Scope: c.Query("ownerScope")
		// Subject: c.Query("ownerSubject")
		// Access: update
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(resourceOwner.Scope, resourceOwner.Subject)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Update {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		hashList, err := resourcesservice.ResourceGetHashlist(ctx, resourceOwner)
		if err != nil {
			rlog.Error("Error getting resource hash list:", err)
			c.JSON(http.StatusInternalServerError, responses.Cluster{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if len(hashList.Items) == 0 {
			c.JSON(http.StatusNotFound, nil)
			return
		}

		c.JSON(http.StatusOK, hashList)
	}
}

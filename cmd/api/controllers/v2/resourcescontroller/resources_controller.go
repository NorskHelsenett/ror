// resourcecontroller implements all controllers for resources
package resourcescontroller

import (
	"net/http"

	"github.com/NorskHelsenett/ror/cmd/api/customvalidators"
	"github.com/NorskHelsenett/ror/cmd/api/responses"
	resourcesservice "github.com/NorskHelsenett/ror/cmd/api/services/resourcesService"
	resourcesv2service "github.com/NorskHelsenett/ror/cmd/api/services/resourcesv2service"

	aclservice "github.com/NorskHelsenett/ror/internal/acl/services"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/context/gincontext"
	"github.com/NorskHelsenett/ror/pkg/context/rorcontext"
	aclmodels "github.com/NorskHelsenett/ror/pkg/models/acl"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"

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

		resourceOwner := apiresourcecontracts.ResourceOwnerReference{
			Scope:   aclmodels.Acl2Scope(c.Query("ownerScope")),
			Subject: c.Query("ownerSubject"),
		}

		if c.Param("uid") == "" {
			c.JSON(http.StatusBadRequest, "")
			return
		}

		// Access check
		// Scope: c.Query("ownerScope")
		// Subject: c.Query("ownerSubject")
		// Access: update
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(resourceOwner.Scope, resourceOwner.Subject)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Update {
			c.JSON(http.StatusForbidden, "")
			return
		}

		if resourcesservice.CheckResourceExist(ctx, c.Param("uid")) {
			c.Status(http.StatusNoContent)
			return
		} else {
			c.JSON(http.StatusNotFound, "")
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
		identity := rorcontext.GetIdentityFromRorContext(ctx)
		var resourceOwner rortypes.RorResourceOwnerReference
		if identity.IsCluster() {

			resourceOwner = rortypes.RorResourceOwnerReference{
				Scope:   aclmodels.Acl2ScopeCluster,
				Subject: aclmodels.Acl2Subject(identity.GetId()),
			}

		} else {
			resourceOwner = rortypes.RorResourceOwnerReference{
				Scope:   aclmodels.Acl2Scope(c.Query("ownerScope")),
				Subject: aclmodels.Acl2Subject(c.Query("ownerSubject")),
			}
		}
		if ok, err := resourceOwner.Validate(); !ok {
			c.JSON(http.StatusBadRequest, responses.Cluster{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		// Access check
		// Scope: c.Query("ownerScope")
		// Subject: c.Query("ownerSubject")
		// Access: update
		accessObject := aclservice.CheckAccessByRorOwnerref(ctx, resourceOwner)
		if !accessObject.Update {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		hashList, err := resourcesv2service.GetHashList(ctx, resourceOwner)
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

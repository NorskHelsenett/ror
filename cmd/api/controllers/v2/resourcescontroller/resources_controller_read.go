// THIS FILE IS GENERATED, DO NOT EDIT
// ref: build/generator/main.go

package resourcescontroller

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"github.com/NorskHelsenett/ror/cmd/api/services/resourcesv2service"

	"github.com/NorskHelsenett/ror/pkg/context/gincontext"
	"github.com/NorskHelsenett/ror/pkg/rorresources"
	"github.com/NorskHelsenett/ror/pkg/helpers/rorerror"

	"github.com/gin-gonic/gin"
)

// Get a list of cluster resources og given group/version/kind.
//
// @Summary	Get resources
// @Schemes
// @Description	Get a list of resources
// @Tags			resources
// @Accept			application/json
// @Produce		application/json
// @Param			ownerScope	query		aclmodels.Acl2Scope	true	"The kind of the owner, currently only support 'Cluster'"
// @Param			ownerSubject	query		string	true	"The name og the owner"
// @Param			apiversion	query		string	true	"ApiVersion"
// @Param			kind	query		string	true	"Kind"
// @Success		200		{array}		apiresourcecontracts.ResourceNode
// @Failure		403		{string}	Forbidden
// @Failure		401		{string}	Unauthorized
// @Failure		500		{string}	Failure	message
// @Router			/v2/resources [get]
// @Security		ApiKey || AccessToken
func GetResources() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		q := c.Query("query")

		rsQuery := rorresources.NewResourceQuery()
		base64Query, err := base64.StdEncoding.DecodeString(q)
		if err != nil {
			c.JSON(http.StatusBadRequest, "400: Invalid base64 query")
			return
		}
		err = json.Unmarshal(base64Query, rsQuery)
		if err != nil {
			c.JSON(http.StatusBadRequest, "400: Invalid query")
			return
		}

		if validationErr := validate.Struct(rsQuery); validationErr != nil {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: validationErr.Error(),
			})
			return
		}

		rsSet := resourcesv2service.GetResourceByQuery(ctx, rsQuery)

		c.JSON(http.StatusOK, rsSet)
	}
}

// Get a cluster resources og given group/version/kind/uid.
//
// @Summary	Get resource
// @Schemes
// @Description	Get a resource by uid
// @Tags			resources
// @Accept			application/json
// @Produce		application/json
// @Param			uid	path		string	true	"The uid of the resource"
// @Param			ownerScope	query		aclmodels.Acl2Scope	true	"The kind of the owner, currently only support 'Cluster'"
// @Param			ownerSubject	query		string	true	"The name og the owner"
// @Param			apiversion	query		string	true	"ApiVersion"
// @Param			kind	query		string	true	"Kind"
// @Success		200		{array}		apiresourcecontracts.ResourceNode
// @Failure		403		{string}	Forbidden
// @Failure		401		{string}	Unauthorized
// @Failure		500		{string}	Failure	message
// @Router			/v2/resource/uid/{uid} [get]
// @Security		ApiKey || AccessToken
func GetResource() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()
	
		if c.Param("uid") == "" {
			c.JSON(http.StatusBadRequest, "400: Missing uid")
			return
		}

		resources := resourcesv2service.GetResourceByUID(ctx, c.Param("uid")) 
		if resources == nil {
			c.JSON(http.StatusNotFound, "404: Resource not found")
			return
		}
		c.JSON(http.StatusOK, resources.GetAll())
	}
}

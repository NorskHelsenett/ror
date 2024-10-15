// TODO: Describe package
package metrics

import (
	metricsservice "github.com/NorskHelsenett/ror/cmd/api/services/metricsService"
	"net/http"

	"github.com/NorskHelsenett/ror/pkg/context/gincontext"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/helpers/rorerror"

	"github.com/gin-gonic/gin"
)

// TODO: Describe function
//
//	@Summary	Get metrics for clusters
//	@Schemes
//	@Description	Get metrics for clusters
//	@Tags			metrics
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200						{object}	apicontracts.MetricList
//	@Failure		403						{string}	Forbidden
//	@Failure		401						{string}	Unauthorized
//	@Failure		500						{string}	Failure	message
//	@Router			/v1/metrics/clusters	[get]
//	@Security		ApiKey || AccessToken
func GetForClusters() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		metrics, err := metricsservice.GetForClusters(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, rorerror.RorError{
				Status:  http.StatusInternalServerError,
				Message: "Could not fetch metrics for clusters",
			})
			return
		}

		if metrics == nil {
			empty := apicontracts.MetricList{}
			c.JSON(http.StatusNotFound, empty)
			return
		}

		c.JSON(http.StatusOK, metrics)
	}
}

// TODO: Describe function
//
//	@Summary	Get metrics for clusters by workspace
//	@Schemes
//	@Description	Get metrics for clusters by workspace
//	@Tags			metrics
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200												{object}	apicontracts.MetricList
//	@Failure		403												{string}	Forbidden
//	@Failure		401												{string}	Unauthorized
//	@Failure		500												{string}	Failure	message
//	@Param			workspaceName									path		string	true	"workspaceName"
//	@Router			/v1/metrics/clusters/workspace/{workspaceName}	[get]
//	@Security		ApiKey || AccessToken
func GetForClustersByWorkspace() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		workspaceName := c.Param("workspaceName")
		defer cancel()

		// importing apicontracts for swagger
		var _ apicontracts.Datacenter
		results, err := metricsservice.GetForClustersByWorkspace(ctx, workspaceName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, rorerror.RorError{
				Status:  http.StatusInternalServerError,
				Message: "Could not fetch metrics for clusters by workspace",
			})
			return
		}

		if results == nil {
			empty := apicontracts.MetricList{}
			c.JSON(http.StatusNotFound, empty)
			return
		}

		c.JSON(http.StatusOK, results)
	}
}

// TODO: Describe function
//
//	@Summary	Get metrics for clusterid
//	@Schemes
//	@Description	Get metrics for clusterid
//	@Tags			metrics
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200	{object}	apicontracts.MetricItem
//	@Failure		403	{string}	Forbidden
//	@Failure		401	{string}	Unauthorized
//	@Failure		500	{string}	Failure	message
//	@Router			/v1/metrics/cluster/{clusterId} [get]
//	@Param			clusterId	path	string	true	"clusterId"
//	@Security		ApiKey || AccessToken
func GetByClusterId() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		clusterId := c.Param("clusterId")
		defer cancel()

		result, err := metricsservice.GetForClusterid(ctx, clusterId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, rorerror.RorError{
				Status:  http.StatusInternalServerError,
				Message: "Could not fetch metrics for clusterid",
			})
			return
		}

		if result == nil {
			empty := apicontracts.MetricItem{}
			c.JSON(http.StatusNotFound, empty)
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

// TODO: Describe function
//
//	@Summary	Get metrics for clusters by a property
//	@Schemes
//	@Description	Get metrics for clusters by a property
//	@Tags			metrics
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200	{object}	apicontracts.MetricsCustom
//	@Failure		403	{string}	Forbidden
//	@Failure		401	{string}	Unauthorized
//	@Failure		500	{string}	Failure	message
//	@Router			/v1/metrics/custom/cluster/{property} [get]
//	@Param			property	path	string	true	"property"
//	@Security		ApiKey || AccessToken
func MetricsForClustersByProperty() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		property := c.Param("property")
		defer cancel()

		result, err := metricsservice.ForClustersByProperty(ctx, property)
		if err != nil {
			c.JSON(http.StatusInternalServerError, rorerror.RorError{
				Status:  http.StatusInternalServerError,
				Message: "Could not fetch custom metrics for clusters by property",
			})
			return
		}

		if result == nil {
			c.JSON(http.StatusOK, apicontracts.MetricsCustom{
				Data: make([]apicontracts.MetricsCustomItem, 0),
			})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

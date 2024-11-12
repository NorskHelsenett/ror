// TODO: Describe function
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
//	@Summary	Get metrics for datacenters
//	@Schemes
//	@Description	Get metrics for datacenters
//	@Tags			metrics
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200						{object}	apicontracts.MetricList
//	@Failure		403						{string}	Forbidden
//	@Failure		401						{string}	Unauthorized
//	@Failure		500						{string}	Failure	message
//	@Router			/v1/metrics/datacenters	[get]
//	@Security		ApiKey || AccessToken
func GetForDatacenters() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		// importing apicontracts for swagger
		var _ apicontracts.MetricList
		metrics, err := metricsservice.GetForDatacenters(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, rorerror.RorError{
				Status:  http.StatusInternalServerError,
				Message: "Could not fetch metrics",
			})
			return
		}

		if metrics == nil {
			empty := apicontracts.MetricList{}
			c.JSON(http.StatusOK, empty)
			return
		}

		c.JSON(http.StatusOK, metrics)
	}
}

// TODO: Describe function
//
//	@Summary	Get metrics for datacenter name
//	@Schemes
//	@Description	Get metrics for datacenter name
//	@Tags			metrics
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200										{object}	apicontracts.MetricItem
//	@Failure		403										{string}	Forbidden
//	@Failure		401										{string}	Unauthorized
//	@Failure		500										{string}	Failure	message
//	@Router			/v1/metrics/datacenter/{datacenterName}	[get]
//	@Param			datacenterName							path	string	true	"datacenterName"
//	@Security		ApiKey || AccessToken
func GetByDatacenterName() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		datacenterName := c.Param("datacenterName")
		defer cancel()

		// importing apicontracts for swagger
		var _ apicontracts.MetricItem

		metrics, err := metricsservice.GetForDatacenterName(ctx, datacenterName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, rorerror.RorError{
				Status:  http.StatusInternalServerError,
				Message: "Could not fetch metrics",
			})
			return
		}

		if metrics == nil {
			empty := apicontracts.MetricItem{}
			c.JSON(http.StatusOK, empty)
			return
		}

		c.JSON(http.StatusOK, metrics)
	}
}

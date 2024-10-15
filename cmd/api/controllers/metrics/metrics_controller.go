// TODO: Describe package
package metrics

import (
	"net/http"

	"github.com/NorskHelsenett/ror/cmd/api/responses"
	metricsservice "github.com/NorskHelsenett/ror/cmd/api/services/metricsService"
	aclservice "github.com/NorskHelsenett/ror/internal/acl/services"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/context/gincontext"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/helpers/rorerror"

	"github.com/gin-gonic/gin"
)

// TODO: Describe function
//
//	@Summary	Get metrics by user
//	@Schemes
//	@Description	Get metrics by user
//	@Tags			metrics
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200			{object}	apicontracts.MetricsTotal
//	@Failure		403			{string}	Forbidden
//	@Failure		401			{string}	Unauthorized
//	@Failure		500			{string}	Failure	message
//	@Router			/v1/metrics	[get]
//	@Security		ApiKey || AccessToken
func GetTotalByUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		// importing apicontracts for swagger
		var _ apicontracts.MetricsTotal

		metricsTotal, err := metricsservice.GetTotalByUser(ctx)
		if err != nil {
			c.JSON(http.StatusUnauthorized, rorerror.RorError{
				Status:  http.StatusUnauthorized,
				Message: "Could not fetch user",
			})
			return
		}

		if metricsTotal == nil {
			empty := apicontracts.MetricsTotal{}
			c.JSON(http.StatusOK, empty)
			return
		}

		c.JSON(http.StatusOK, metricsTotal)
	}
}

// TODO: Describe function
//
//	@Summary	Get metrics
//	@Schemes
//	@Description	Get metrics
//	@Tags			metrics
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200					{object}	apicontracts.MetricsTotal
//	@Failure		403					{string}	Forbidden
//	@Failure		401					{string}	Unauthorized
//	@Failure		500					{string}	Failure	message
//	@Router			/v1/metrics/total	[get]
//	@Security		ApiKey || AccessToken
func GetTotal() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		// importing apicontracts for swagger
		var _ apicontracts.MetricsTotal

		metrics, err := metricsservice.GetTotal(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, rorerror.RorError{
				Status:  http.StatusInternalServerError,
				Message: "Could not fetch metrics",
			})
			return
		}

		if metrics == nil {
			empty := apicontracts.MetricsTotal{}
			c.JSON(http.StatusOK, empty)
			return
		}

		c.JSON(http.StatusOK, metrics)
	}
}

// Registers metrics from ror-agent
//
//	@Summary	Register metrics
//	@Schemes
//	@Description	Register metrics
//	@Tags			metrics
//	@Accept			application/json
//	@Produce		application/json
//	@Param			metrics		body		apicontracts.MetricsReport	true	"MetricsReport"
//	@Success		200			{object}	apicontracts.MetricsTotal
//	@Failure		403			{string}	Forbidden
//	@Failure		401			{string}	Unauthorized
//	@Failure		500			{string}	Failure	message
//	@Router			/v1/metrics	[post]
//	@Security		ApiKey || AccessToken
func RegisterResourceMetricsReport() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		var input apicontracts.MetricsReport
		defer cancel()

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
		ownerref := apiresourcecontracts.ResourceOwnerReference{
			Scope:   input.Owner.Scope,
			Subject: string(input.Owner.Subject),
		}
		accessObject := aclservice.CheckAccessByOwnerref(ctx, ownerref)
		if !accessObject.Update {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		err := metricsservice.ProcessMetricReport(ctx, &input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Cluster{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, nil)
	}
}

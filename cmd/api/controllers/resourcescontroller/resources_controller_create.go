package resourcescontroller

import (
	"net/http"

	"github.com/NorskHelsenett/ror/cmd/api/responses"
	resourcesservice "github.com/NorskHelsenett/ror/cmd/api/services/resourcesService"
	aclservice "github.com/NorskHelsenett/ror/internal/acl/services"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/context/gincontext"

	aclmodels "github.com/NorskHelsenett/ror/pkg/models/aclmodels"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
)

// Register a new resource, the resource is in the payload.
// Parameter clusterid must match authorized clusterid
//
//	@Summary	Register  resource
//	@Schemes
//	@Description	Registers a  resource
//	@Tags			resources
//	@Accept			application/json
//	@Produce		application/json
//	@Param			resourcereport	body		apiresourcecontracts.ResourceUpdateModel	true	"ResourceUpdate"
//	@Success		200				{bool}		bool
//	@Failure		403				{string}	Forbidden
//	@Failure		401				{string}	Unauthorized
//	@Failure		500				{string}	Failure	message
//	@Router			/v1/resources [post]
//	@Security		ApiKey || AccessToken
func NewResource() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		ctx, span := otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "Resource create controller")
		defer span.End()
		var input apiresourcecontracts.ResourceUpdateModel

		_, span1 := otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "Validate request")
		defer span1.End()

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

		span1.AddEvent("Request validated")
		span1.End()
		_, span2 := otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "Check access")
		defer span2.End()
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

		span2.AddEvent("Access checked")
		span2.End()
		_, span3 := otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "Run service: resourceservice.ResourceNewCreateService")
		defer span3.End()

		err := resourcesservice.ResourceNewCreateService(ctx, input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Cluster{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		span3.AddEvent("Resource created")
		span3.End()
		_, span4 := otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "Return response")
		defer span4.End()

		c.JSON(http.StatusCreated, nil)
	}
}

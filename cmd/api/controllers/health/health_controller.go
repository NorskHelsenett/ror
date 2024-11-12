// TODO: Describe package
package health

import (
	"context"
	"fmt"
	"net/http"

	"github.com/NorskHelsenett/ror/cmd/api/apiconnections"
	"github.com/NorskHelsenett/ror/cmd/api/responses"
	"strings"
	"time"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"github.com/NorskHelsenett/ror/pkg/clients/redisdb"

	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
)

//var tracer trace.Tracer = otel.GetTracerProvider().Tracer("HealthCheck")

// TODO: Describe function
func Ping(url string) (int, error) {
	var client = http.Client{
		Timeout: 2 * time.Second,
	}
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}
	_ = resp.Body.Close()
	return resp.StatusCode, nil
}

// TODO: Describe function
func PingAndUpdateStatus(url string, ctx context.Context) responses.HealthStatusCode {
	statusCode, err := Ping(url)
	if statusCode == 0 {
		rlog.Errorc(ctx, "", err)
		return responses.StatusUnableToPing
	}
	if statusCode == -1 {
		rlog.Errorc(ctx, "", err)
		return responses.StatusNotConnected
	}
	return responses.StatusOK
}

// TODO: Describe function
func GetMongoDBStatus(ctx context.Context) responses.HealthStatusCode {
	if !mongodb.Ping() {
		rlog.Errorc(ctx, "could not ping mongodb", fmt.Errorf(""))
		return responses.StatusNotConnected
	}
	return responses.StatusOK
}

// TODO: Describe function
func GetRabbitMQStatus(ctx context.Context) responses.HealthStatusCode {
	if !apiconnections.RabbitMQConnection.Ping() {
		rlog.Errorc(ctx, "could not ping rabbitmq", fmt.Errorf(""))
		return responses.StatusNotConnected
	}
	return responses.StatusOK
}

func GetRedisStatus(ctx context.Context) responses.HealthStatusCode {
	if !redisdb.Ping() {
		rlog.Errorc(ctx, "could not ping rabbitmq", fmt.Errorf(""))
		return responses.StatusNotConnected
	}
	return responses.StatusOK
}

// TODO: Describe function
func GetTracingStatus(ctx context.Context) responses.HealthStatusCode {
	tracerProvider := otel.GetTracerProvider()
	tracerProviderType := fmt.Sprintf("%T", tracerProvider)
	if strings.Contains(tracerProviderType, "global") {
		rlog.Errorc(ctx, "opentelemetry not connected", fmt.Errorf("opentelemetry not connected"))
		return responses.StatusNotConnected
	}
	return responses.StatusOK
}

// TODO: Describe function
//
//	@BasePath	/
//	@Summary	Health status
//	@Schemes
//	@Description	Get health status for NHN-ROR-API
//	@Tags			health status
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200	{object}	apicontracts.HealthStatus
//	@Failure		500	{string}	Failure	message
//	@Router			/health [get]
func GetHealthStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var healthstatus responses.HealthStatus
		var services []responses.Services
		httpStatus := http.StatusOK
		services = append(services, responses.Services{Name: "identityProvider", Status: PingAndUpdateStatus(viper.GetString(configconsts.OIDC_PROVIDER), c.Request.Context())})
		services = append(services, responses.Services{Name: "database", Status: GetMongoDBStatus(c.Request.Context())})
		services = append(services, responses.Services{Name: "messageBus", Status: GetRabbitMQStatus(c.Request.Context())})
		services = append(services, responses.Services{Name: "secrets", Status: PingAndUpdateStatus(viper.GetString(configconsts.VAULT_URL), c.Request.Context())})
		//services = append(services, responses.Services{Name: "redis", Status: GetRedisStatus(c.Request.Context())})

		// services = append(services, responses.Services{Name: "tracing", Status: GetTracingStatus(c.Request.Context())})
		healthstatus.Services = services

		for _, v := range services {
			if v.Status != responses.StatusOK && v.Name != "messageBus" && v.Name != "tracing" {
				httpStatus = http.StatusInternalServerError
			}
		}

		// importing apicontracts for swagger
		var _ apicontracts.HealthStatus
		c.JSON(httpStatus,
			healthstatus,
		)
	}
}

package health

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/api/responses"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"go.opentelemetry.io/otel"
)

var mockDB *mongo.Client

func GetMongoDBStatusMock() responses.HealthStatusCode {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := mockDB.Ping(ctx, nil)
	if err != nil {
		rlog.Errorc(ctx, "could not ping", err)
		return responses.StatusNotConnected
	}
	return responses.StatusOK
}

func GetRabbitMQStatusMock() responses.HealthStatusCode {
	return responses.StatusOK
}

func GetTracingStatusMock() responses.HealthStatusCode {
	tracerProvider := otel.GetTracerProvider()
	tracerProviderType := fmt.Sprintf("%T", tracerProvider)
	if !strings.Contains(tracerProviderType, "global") {
		return responses.StatusNotConnected
	}
	return responses.StatusOK
}

func GetHealthStatusMock() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		var healthstatus responses.HealthStatus
		var services []responses.Services
		services = append(services, responses.Services{Name: "identityProvider", Status: PingAndUpdateStatus(viper.GetString(configconsts.OIDC_PROVIDER), ctx)})
		services = append(services, responses.Services{Name: "database", Status: GetMongoDBStatusMock()})
		services = append(services, responses.Services{Name: "messageBus", Status: GetRabbitMQStatusMock()})
		services = append(services, responses.Services{Name: "secrets", Status: PingAndUpdateStatus(viper.GetString(configconsts.VAULT_URL), ctx)})
		services = append(services, responses.Services{Name: "tracing", Status: GetTracingStatusMock()})
		healthstatus.Services = services

		c.JSON(http.StatusOK,
			healthstatus,
		)
	}
}

func init() {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	viper.Set(configconsts.OIDC_PROVIDER, testServer.URL)
	viper.Set(configconsts.VAULT_URL, testServer.URL)
	_ = mtest.Setup()
}
func TestGetHealthStatus(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	testTable := map[string]struct {
		want responses.HealthStatus
	}{
		"Success": {
			want: responses.HealthStatus{
				Services: []responses.Services{
					{Name: "identityProvider", Status: responses.StatusOK},
					{Name: "database", Status: responses.StatusOK},
					{Name: "messageBus", Status: responses.StatusOK},
					{Name: "secrets", Status: responses.StatusOK},
					{Name: "tracing", Status: responses.StatusOK},
				},
			},
		},
	}

	for name, tc := range testTable {
		mt.Run(name, func(mt *mtest.T) {
			mt.AddMockResponses(mtest.CreateSuccessResponse())
			mockDB = mt.Client

			router := gin.New()
			router.Use(gin.Logger())
			router.Use(gin.Recovery())

			_ = router.SetTrustedProxies([]string{"localhost"})

			router.GET("/health", GetHealthStatusMock())

			rr := httptest.NewRecorder()

			request, err := http.NewRequest(http.MethodGet, "/health", nil)
			assert.NoError(mt, err)

			router.ServeHTTP(rr, request)

			response := responses.HealthStatus{}
			responseBytes, _ := io.ReadAll(rr.Result().Body)

			err = json.Unmarshal(responseBytes, &response)
			if err != nil {
				rlog.Error("could not unmarshal", err)
			}

			assert.Equal(mt, http.StatusOK, rr.Result().StatusCode)
			assert.Equal(mt, tc.want, response)

		})
	}
}

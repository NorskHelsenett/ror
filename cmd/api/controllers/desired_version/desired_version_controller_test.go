package desired_version

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/api/responses"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var mockData []apicontracts.DesiredVersion

func MockGetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, mockData)
	}
}

func MockCreateOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input apicontracts.DesiredVersion

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.Cluster{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		for _, v := range mockData {
			if v.Key == input.Key {
				c.JSON(http.StatusBadRequest, responses.Cluster{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": fmt.Sprintf("Desired version with key: %s already exists.", input.Key)}})
				return
			}
		}

		mockData = append(mockData, input)

		c.Status(200)
	}
}

func MockGetOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Param("key")

		var result apicontracts.DesiredVersion

		for _, v := range mockData {
			if v.Key == key {
				result = v
			}
		}

		c.JSON(200, result)
	}
}

func TestInitialRequest(t *testing.T) {

	gin.SetMode(gin.TestMode)

	// Insert mock data
	{
		mockData = append(mockData, apicontracts.DesiredVersion{Key: "kubernetes", Value: "20.alpha"})
		mockData = append(mockData, apicontracts.DesiredVersion{Key: "ror-agent", Value: "30.beta"})
		mockData = append(mockData, apicontracts.DesiredVersion{Key: "vsphere", Value: "1.alpha"})
	}

	t.Run("Success", func(t *testing.T) {
		router := gin.New()
		// Set up the routes we are testing
		{

			router.Use(gin.Logger())
			router.Use(gin.Recovery())

			_ = router.SetTrustedProxies([]string{"localhost"})

			router.GET("/v1/workspaces", MockGetAll())
			router.POST("/v1/workspaces", MockCreateOne())
			router.GET("/v1/workspaces/:key", MockGetOne())
		}

		{
			rr := httptest.NewRecorder()

			request, err := http.NewRequest(http.MethodGet, "/v1/workspaces", nil)

			assert.NoError(t, err)

			router.ServeHTTP(rr, request)

			assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
		}

		// Test that we fail if we insert a desired version we already have
		{
			rr := httptest.NewRecorder()

			input := apicontracts.DesiredVersion{
				Key:   "vsphere",
				Value: "30",
			}

			raw, err := json.Marshal(input)
			assert.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, "/v1/workspaces", bytes.NewBuffer(raw))

			assert.NoError(t, err)

			router.ServeHTTP(rr, request)

			assert.Equal(t, http.StatusBadRequest, rr.Result().StatusCode)

		}

		{
			rr := httptest.NewRecorder()

			input := apicontracts.DesiredVersion{
				Key:   "dex",
				Value: "30",
			}

			raw, err := json.Marshal(input)
			assert.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, "/v1/workspaces", bytes.NewBuffer(raw))

			assert.NoError(t, err)

			router.ServeHTTP(rr, request)

			assert.Equal(t, http.StatusOK, rr.Result().StatusCode)

		}

		{
			rr := httptest.NewRecorder()

			request, err := http.NewRequest(http.MethodGet, "/v1/workspaces/vsphere", nil)

			assert.NoError(t, err)

			router.ServeHTTP(rr, request)

			assert.Equal(t, http.StatusOK, rr.Result().StatusCode)

		}
	})
}

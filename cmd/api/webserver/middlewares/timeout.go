package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func timeoutResponse(c *gin.Context) {
	json := map[string]any{"success": false, "message": "timeout"}
	c.JSON(http.StatusRequestTimeout, json)
}

func TimeoutMiddleware(duration time.Duration) gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(duration),
		timeout.WithHandler(func(c *gin.Context) {
			c.Set("timeout", duration)
			c.Next()
		}),
		timeout.WithResponse(timeoutResponse),
	)
}

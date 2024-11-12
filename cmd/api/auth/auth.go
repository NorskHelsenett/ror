package auth

import (
	"strings"

	"github.com/NorskHelsenett/ror/pkg/helpers/rorerror"
	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware(c *gin.Context) {
	xapikey := c.Request.Header.Get("X-API-KEY")
	if len(xapikey) > 0 {
		ApiKeyAuth(c)
		return
	}

	authorization := c.Request.Header.Get("Authorization")
	if strings.HasPrefix(authorization, "Bearer ") {
		DexMiddleware(c)
		return
	}

	rorerror := rorerror.RorError{
		Status:  401,
		Message: "Authorization provider not supported",
	}
	c.AbortWithStatusJSON(rorerror.Status, rorerror)
}

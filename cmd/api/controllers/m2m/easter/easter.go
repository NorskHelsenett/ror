package easter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterM2m() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://youtu.be/ZCFlT_FYnEE")
	}
}

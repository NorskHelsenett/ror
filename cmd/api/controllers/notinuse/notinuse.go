package notinuse

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func NotInUse() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("Hei %s du har kommet til et nummer som ikke er i bruk: %s%s\n", c.Request.Header["X-Forwarded-For"], c.Request.Host, c.Request.URL)
		for key, header := range c.Request.Header {
			fmt.Printf("%s: %s\n", key, header)
		}
	}
}

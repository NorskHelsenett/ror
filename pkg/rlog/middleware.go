package rlog

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		ctx := c.Request.Context()
		start := time.Now()
		c.Next()
		end := time.Now()
		statusCode := c.Writer.Status()
		if statusCode > 299 {
			latency := end.Sub(start)
			raw := c.Request.URL.RawQuery
			clientIP := c.ClientIP()
			forwardedHdr := c.Request.Header["X-Forwarded-For"]
			forwardedFor := ""
			if len(forwardedHdr) > 0 {
				forwardedFor = forwardedHdr[0]
			}
			userAgent := c.Request.Header["User-Agent"]
			method := c.Request.Method
			var errors []error
			for _, v := range c.Errors {
				errors = append(errors, v.Err)
			}
			bodySize := c.Writer.Size()

			if raw != "" {
				path = path + "?" + raw
			}

			Infoc(
				ctx,
				"",
				zap.String("method", method),
				zap.Int("status_code", statusCode),
				zap.String("path", path),
				zap.Duration("latency", latency),
				zap.String("userAgent", userAgent[0]),
				zap.String("client_ip", clientIP),
				zap.String("forwarded_for", forwardedFor),
				zap.Int("body_size", bodySize),
				zap.Errors("errors", errors),
			)
		}
	}

}

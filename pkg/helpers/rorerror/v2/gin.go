package rorerror

import (
	"strings"

	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/gin-gonic/gin"
)

type Field = rlog.Field

func GinHandleErrorAndAbort(c *gin.Context, status int, err error, fields ...Field) bool {
	if err != nil {
		rorerror := NewRorErrorFromError(status, err)
		fields = append(fields, rlog.Int("statuscode", status))
		rorerror.GinLogErrorAbort(c, fields...)
		return true
	}
	return false
}

func maskValue(value string) string {
	maskedKey := string(value[0:2]) + strings.Repeat("*", len(value)-4) + string(value[len(value)-2:])
	return maskedKey
}

// GinLogErrorJSON logs the error and returns the request with the rorerror struct as a JSON response
func (e ErrorData) GinLogErrorJSON(c *gin.Context, fields ...Field) {
	e.logError(c, fields...)
	c.JSON(e.Status, e)
}

// GinLogErrorAbort logs the error and aborts the request with the rorerror struct as a JSON response
func (e ErrorData) GinLogErrorAbort(c *gin.Context, fields ...Field) {
	e.logError(c, fields...)
	c.AbortWithStatusJSON(e.Status, e)
}

// logError logs the error with the given fields
func (e ErrorData) logError(c *gin.Context, fields ...Field) {
	if len(e.errors) > 0 {
		for _, errs := range e.errors {
			fields = append(fields, rlog.String("error", errs.Error()))
		}
	}
	for i, field := range fields {
		if field.Key == "apikey" {
			fields[i].String = maskValue(fields[i].String)
		}
	}
	rlog.Errorc(c.Request.Context(), "error", e, fields...)
}

package rorerror

import (
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/gin-gonic/gin"
)

// GinLogErrorJSON logs the error and returns the request with the rorerror struct as a JSON response
func (e RorError) GinLogErrorJSON(c *gin.Context, fields ...Field) {
	e.logError(c, fields...)
	c.JSON(e.Status, e)
}

// GinLogErrorAbort logs the error and aborts the request with the rorerror struct as a JSON response
func (e RorError) GinLogErrorAbort(c *gin.Context, fields ...Field) {
	e.logError(c, fields...)
	c.AbortWithStatusJSON(e.Status, e)
}

// logError logs the error with the given fields
func (e RorError) logError(c *gin.Context, fields ...Field) {
	zfields := make([]rlog.Field, len(fields))
	if len(e.errors) > 0 {
		for _, errs := range e.errors {
			zfields = append(zfields, rlog.String("error", errs.Error()))
		}
	}

	for _, field := range fields {
		zfields = append(zfields, field.ToRlog())
	}
	rlog.Errorc(c.Request.Context(), "error", e, zfields...)
}

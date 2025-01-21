package rorerror

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/gin-gonic/gin"
)

type RorError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	errors  []error
}

func NewRorErrorFromError(status int, err error) RorError {
	rorerror := RorError{
		Status:  status,
		Message: err.Error(),
	}
	return rorerror
}
func NewRorError(status int, err string, errors ...error) RorError {
	rorerror := RorError{
		Status:  status,
		Message: err,
		errors:  errors,
	}
	return rorerror
}

func GinHandleErrorAndAbort(c *gin.Context, status int, err error, fields ...Field) bool {
	if err != nil {
		rorerror := NewRorErrorFromError(status, err)
		fields = append(fields, String("statuscode", fmt.Sprintf("%d", status)))
		zfields := make([]rlog.Field, len(fields))
		for i, field := range fields {
			if field.Key != "apikey" {
				zfields[i] = field.ToRlog()
			} else {
				zfields[i] = rlog.String(field.Key, maskApiKey(field.Value))
			}

		}
		rlog.Errorc(c.Request.Context(), "error:", err, zfields...)
		c.AbortWithStatusJSON(rorerror.Status, rorerror)
		return true
	}
	return false
}

func maskApiKey(apikey string) string {
	maskedKey := string(apikey[0:2]) + strings.Repeat("*", len(apikey)-4) + string(apikey[len(apikey)-2:])
	return maskedKey
}

func (e RorError) Error() string {
	return fmt.Sprintf("Status: %d, Message: %s", e.Status, e.Message)
}

func (e RorError) AsJson() []byte {
	ret, _ := json.Marshal(e)
	return ret
}

func (e RorError) AsString() string {
	return string(e.AsJson())
}
func (e RorError) GinLogErrorJSON(c *gin.Context, fields ...Field) {
	e.logError(c, fields...)
	c.JSON(e.Status, e)
}
func (e RorError) GinLogErrorAbort(c *gin.Context, fields ...Field) {
	e.logError(c, fields...)
	c.AbortWithStatusJSON(e.Status, e)
}
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

func (e RorError) GinLogErrorAndAbortWithMessage(c *gin.Context, message string, fields ...rlog.Field) {
	rlog.Errorc(c.Request.Context(), message, e, fields...)
	c.AbortWithStatusJSON(e.Status, e)
}

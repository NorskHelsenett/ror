package rorginerror

import (
	"github.com/NorskHelsenett/ror/pkg/helpers/rorerror/v2"
	"github.com/gin-gonic/gin"
)

type RorGinError interface {
	rorerror.RorError

	// GinLogErrorAbort logs the error with context and aborts the Gin request
	// with a JSON response containing the error details.
	GinLogErrorAbort(c *gin.Context, fields ...Field)

	// GinLogErrorJSON logs the error with context and returns a JSON response
	// containing the error details without aborting the request.
	GinLogErrorJSON(c *gin.Context, fields ...Field)
}
type RorGinErrorData struct {
	rorerror.RorError
}

func NewRorGinErrorFromError(status int, err error) RorGinError {
	rorerror := RorGinErrorData{
		RorError: rorerror.NewRorErrorFromError(status, err),
	}
	return rorerror
}

func NewRorGinError(status int, err string, errors ...error) RorGinError {
	rorerror := RorGinErrorData{
		RorError: rorerror.NewRorError(status, err, errors...),
	}
	return rorerror
}

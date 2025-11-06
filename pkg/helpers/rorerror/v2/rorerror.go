package rorerror

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type RorError interface {
	GetStatusCode() int
	GetMessage() string
	Error() string
	GinLogErrorAbort(c *gin.Context, fields ...Field)
	GinLogErrorJSON(c *gin.Context, fields ...Field)
}

type ErrorData struct {
	Status  int    `json:"status" example:"400"`          // HTTP status code
	Message string `json:"message" example:"Bad Request"` // Error message
	errors  []error
}

func NewRorErrorFromError(status int, err error) RorError {
	rorerror := ErrorData{
		Status:  status,
		Message: err.Error(),
	}
	return rorerror
}
func NewRorError(status int, err string, errors ...error) ErrorData {
	rorerror := ErrorData{
		Status:  status,
		Message: err,
		errors:  errors,
	}
	return rorerror
}

func (e ErrorData) Error() string {
	return fmt.Sprintf("Status: %d, Message: %s", e.Status, e.Message)
}

func (e ErrorData) GetStatusCode() int {
	return e.Status
}

func (e ErrorData) GetMessage() string {
	return e.Message
}

func (e ErrorData) RorErrorData() ErrorData {
	return e
}

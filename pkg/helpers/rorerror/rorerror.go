package rorerror

import (
	"encoding/json"
	"fmt"
)

var NoRorError = RorError{}

type RorError struct {
	Status  int    `json:"status" example:"400"`          // HTTP status code
	Message string `json:"message" example:"Bad Request"` // Error message
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

func (e RorError) IsError() bool {
	return e.Status != 0
}

package rorerror

import (
	"encoding/json"
	"fmt"
)

type RorError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
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
	}
	if len(errors) > 0 {
		for _, errs := range errors {
			rorerror.Message = fmt.Sprintf("%s Error: %s", rorerror.Message, errs.Error())
		}
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

package m2m

import (
	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

func init() {
	validate = validator.New()
}

package validationService

import (
	"github.com/go-playground/validator/v10"
)

type ValidationService interface {
	Validate(data interface{}) error
}
type validationSrv struct {
}

func (v *validationSrv) Validate(data interface{}) error {
	validate := validator.New()
	return validate.Struct(data)

}

func NewValidationSrv() ValidationService {
	return &validationSrv{}
}

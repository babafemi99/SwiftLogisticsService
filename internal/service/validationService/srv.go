package validationService

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"sls/internal/entity/adminEntity"
	"sls/internal/service/timeService"
)

type ValidationService interface {
	Validate(data interface{}) error
	ValidateDPR(dpr *adminEntity.DPR) error
	//ValidateCoupon(code, userId string) error
}
type validationSrv struct {
	timeSrv timeService.TimeSrv
}

func NewValidationSrv(timeSrv timeService.TimeSrv) ValidationService {
	return &validationSrv{timeSrv: timeSrv}
}

func (v *validationSrv) ValidateDPR(dpr *adminEntity.DPR) error {
	switch dpr.Status {
	case "ALL":
		if dpr.Seconds != 0 && dpr.Kilometers != 0 {
			return nil
		} else {
			return errors.New("validation error, check seconds and kilometers fields")
		}
	case "KMS":
		if dpr.Kilometers != 0 {
			dpr.Seconds = 0
			return nil
		} else {
			return errors.New("validation error, check kilometers field")
		}
	case "SECS":
		if dpr.Seconds != 0 {
			dpr.Kilometers = 0
			return nil
		} else {
			return errors.New("validation error, check seconds field")
		}
	case "NIL":
		dpr.Kilometers = 0
		dpr.Seconds = 0
		return nil
	default:
		return errors.New("wrong status input")
	}

}

func (v *validationSrv) Validate(data interface{}) error {
	validate := validator.New()
	return validate.Struct(data)

}

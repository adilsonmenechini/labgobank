package utils

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func ValidateStruct(req interface{}) error {
	_, err := govalidator.ValidateStruct(req)

	if err != nil {
		return err
	}

	return nil

}

package validatorx

import "github.com/go-playground/validator/v10"

var (
	validate = validator.New()
)

func Validate(data interface{}) error {
	return validate.Struct(data)
}

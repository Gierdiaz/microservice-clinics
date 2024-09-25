package validator

import (
	"strings"

	"github.com/Gierdiaz/diagier-clinics/internal/domain/user"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func InitValidator() {
	validate = validator.New()
}

func Validate(i interface{}) error {
	return validate.Struct(i)
}

func ValidateRegister(req user.AuthRequest) map[string]string {
	err := validate.Struct(req)
	if err != nil {
		return TranslateValidationErrors(err)
	}
	return nil
}

func TranslateValidationErrors(err error) map[string]string {
	errs := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		field := strings.ToLower(err.Field())
		errs[field] = err.Tag()
	}
	return errs
}
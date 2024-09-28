package validator

import (
	"strings"

	userApplication "github.com/Gierdiaz/diagier-clinics/internal/application/DTO"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func InitValidator() {
	validate = validator.New()
}

func Validate(i interface{}) error {
	return validate.Struct(i)
}

func ValidateRegister(req userApplication.AuthRequest) map[string]string {
	if err := validate.Struct(req); err != nil {
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

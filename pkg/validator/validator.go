package validator

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func InitValidator() {
	validate = validator.New()
}

// Função para validar qualquer estrutura
func Validate(i interface{}) error {
	return validate.Struct(i)
}
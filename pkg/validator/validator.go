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

// Função para validar qualquer estrutura
func Validate(i interface{}) error {
	return validate.Struct(i)
}

// Função para validar o AuthRequest específico do registro
func ValidateRegister(req user.AuthRequest) map[string]string {
	err := validate.Struct(req)
	if err != nil {
		return TranslateValidationErrors(err)
	}
	return nil
}

// Função auxiliar para traduzir erros de validação em mensagens amigáveis
func TranslateValidationErrors(err error) map[string]string {
	errs := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		field := strings.ToLower(err.Field())
		errs[field] = err.Tag() // Exemplo: retorna o nome da tag de validação como mensagem de erro
	}
	return errs
}

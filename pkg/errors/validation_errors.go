package errors

import "fmt"

var (
	ErrInvalidName    = "o nome deve ter entre 3 e 100 caracteres"
	ErrInvalidAge     = "a idade deve estar entre 1 e 150 anos"
	ErrInvalidGender  = "o gênero deve ser 'masculino', 'feminino' ou 'outro'"
	ErrInvalidAddress = "o endereço deve ter entre 3 e 100 caracteres"
	ErrInvalidPhone   = "o número de telefone deve seguir o formato internacional (+55XXXXXXXXXX)"
	ErrInvalidEmail   = "o email deve estar preenchido e seguir o formato padrão (exemplo@dominio.com)"
)

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("campo '%s': %s", e.Field, e.Message)
}

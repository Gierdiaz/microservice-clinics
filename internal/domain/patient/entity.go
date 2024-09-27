package patient

import (
	"fmt"
	"regexp"
	"time"

	"github.com/google/uuid"
)

type Patient struct {
	ID           uuid.UUID `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	Age          int       `db:"age" json:"age"`
	Gender       string    `db:"gender" json:"gender"`
	Address      string    `db:"address" json:"address"`
	Phone        string    `db:"phone" json:"phone"`
	Email        string    `db:"email" json:"email"`
	Observations string    `db:"observations" json:"observations"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

// Erros específicos de validação do domínio Patient
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

// Validate checks for the validity of Patient fields
func (p *Patient) Validate() error {
	if len(p.Name) < 3 || len(p.Name) > 100 {
		return &ValidationError{Field: "name", Message: ErrInvalidName}
	}

	if p.Age < 0 || p.Age > 150 {
		return &ValidationError{Field: "age", Message: ErrInvalidAge}
	}

	if p.Gender != "masculino" && p.Gender != "feminino" && p.Gender != "outro" {
		return &ValidationError{Field: "gender", Message: ErrInvalidGender}
	}

	if len(p.Address) < 3 || len(p.Address) > 100 {
		return &ValidationError{Field: "address", Message: ErrInvalidAddress}
	}

	if len(p.Phone) == 0 {
		return &ValidationError{Field: "phone", Message: ErrInvalidPhone}
	}
	e164Regex := regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
	if !e164Regex.MatchString(p.Phone) {
		return &ValidationError{Field: "phone", Message: ErrInvalidPhone}
	}

	if len(p.Email) == 0 {
		return &ValidationError{Field: "email", Message: ErrInvalidEmail}
	}

	return nil
}

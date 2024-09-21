package patient

import (
	"regexp"
	"time"

	"github.com/Gierdiaz/diagier-clinics/pkg/errors"
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

// Validate checks for the validity of Patient fields
func (p *Patient) Validate() error {
	if len(p.Name) < 3 || len(p.Name) > 100 {
		return &errors.ValidationError{Field: "name", Message: errors.ErrInvalidName}
	}

	if p.Age < 0 || p.Age > 150 {
		return &errors.ValidationError{Field: "age", Message: errors.ErrInvalidAge}
	}

	if p.Gender != "masculino" && p.Gender != "feminino" && p.Gender != "outro" {
		return &errors.ValidationError{Field: "gender", Message: errors.ErrInvalidGender}
	}

	if len(p.Address) < 3 || len(p.Address) > 100 {
		return &errors.ValidationError{Field: "address", Message: errors.ErrInvalidAddress}
	}

	if len(p.Phone) == 0 {
		return &errors.ValidationError{Field: "phone", Message: errors.ErrInvalidPhone}
	}
	e164Regex := regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
	if !e164Regex.MatchString(p.Phone) {
		return &errors.ValidationError{Field: "phone", Message: errors.ErrInvalidPhone}
	}

	if len(p.Email) == 0 {
		return &errors.ValidationError{Field: "email", Message: errors.ErrInvalidEmail}
	}

	return nil
}

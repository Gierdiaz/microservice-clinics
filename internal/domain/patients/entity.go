package patients

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
)

type Pacient struct {
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

func (p *Pacient) Validate() error {

	if len(p.Name) < 3 || len(p.Name) > 100 {
		return errors.New("o nome deve ter entre 3 e 100 caracteres")
	}

	if p.Age < 0 || p.Age > 150 {
		return errors.New("a idade deve ser entre 0 e 150")
	}

	if p.Gender != "male" && p.Gender != "female" && p.Gender != "other" {
		return errors.New("o genero deve ser 'male', 'female' ou 'other'")
	}

	if len(p.Address) < 3 || len(p.Address) > 100 {
		return errors.New("o endereço deve ter entre 3 e 100 caracteres")
	}

	if len(p.Phone) == 0 {
		return errors.New("o telefone deve ser preenchido")
	}

	e164Regex := regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
	if !e164Regex.MatchString(p.Phone) {
		return errors.New("phone number must be in valid E.164 format (e.g., +5511999999999)")
	}

	if len(p.Email) == 0 {
		return errors.New("o email deve ser preenchido")
	}

	return nil
}

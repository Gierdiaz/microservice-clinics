package patient

import (
	"github.com/google/uuid"
)

type PatientRepository interface {
	Index() ([]*Patient, error)
	Show(id uuid.UUID) (*Patient, error)
	Store(patient *Patient) (*Patient, error)
	Update(patient *Patient) (*Patient, error)
	Delete(id uuid.UUID) error
}

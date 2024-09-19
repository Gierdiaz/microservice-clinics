package patient

import (
	"time"

	"github.com/google/uuid"
)

type PatientService interface {
	GetAllPatients() ([]*Patient, error)
	GetPatientByID(id uuid.UUID) (*Patient, error)
	CreatePatient(dto *PatientDTO) (*Patient, error)
	UpdatePatient(id uuid.UUID, dto *PatientDTO) error
	DeletePatient(id uuid.UUID) error
}

type patientService struct {
	repository PatientRepository
}

func NewPatientService(repository PatientRepository) PatientService {
	return &patientService{repository: repository}
}

func (service *patientService) GetAllPatients() ([]*Patient, error) {
	return service.repository.Index()
}

func (service *patientService) GetPatientByID(id uuid.UUID) (*Patient, error) {
	return service.repository.Show(id)
}

func (service *patientService) CreatePatient(dto *PatientDTO) (*Patient, error) {
	patient := dto.ToEntity()
	if err := patient.Validate(); err != nil {
		return nil, err
	}
	return service.repository.Store(patient)
}

func (service *patientService) UpdatePatient(id uuid.UUID, dto *PatientDTO) error {
	patient, err := service.repository.Show(id)
	if err != nil {
		return err
	}
	patient.Name = dto.Name
	patient.Age = dto.Age
	patient.Gender = dto.Gender
	patient.Address = dto.Address
	patient.Phone = dto.Phone
	patient.Email = dto.Email
	patient.Observations = dto.Observations
	patient.UpdatedAt = time.Now()

	if err := patient.Validate(); err != nil {
		return err
	}

	_, err = service.repository.Update(patient)
	return err
}

func (service *patientService) DeletePatient(id uuid.UUID) error {
	err := service.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

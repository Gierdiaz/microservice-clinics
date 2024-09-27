package patient

import (
	"encoding/json"
	"log"
	"time"

	"github.com/Gierdiaz/diagier-clinics/pkg/messaging"
	"github.com/google/uuid"
)

type PatientService interface {
	GetAllPatients() ([]*Patient, error)
	GetPatientByID(id uuid.UUID) (*Patient, error)
	CreatePatient(patient *Patient) (*Patient, error)
	UpdatePatient(id uuid.UUID, patient *Patient) (*Patient, error)
	DeletePatient(id uuid.UUID) error
}

type patientService struct {
	repository PatientRepository
	rabbitMQ   *messaging.RabbitMQ
}

func NewPatientService(repository PatientRepository, rabbitMQ *messaging.RabbitMQ) PatientService {
	return &patientService{repository: repository, rabbitMQ: rabbitMQ}
}

func (service *patientService) GetAllPatients() ([]*Patient, error) {
	return service.repository.Index()
}

func (service *patientService) GetPatientByID(id uuid.UUID) (*Patient, error) {
	return service.repository.Show(id)
}

func (service *patientService) CreatePatient(patient *Patient) (*Patient, error) {
	if err := patient.Validate(); err != nil {
		return nil, err
	}
	createdPatient, err := service.repository.Store(patient)
	if err != nil {
		return nil, err
	}

	message, _ := json.Marshal(createdPatient)
	if err := service.rabbitMQ.Publish("patients", message); err != nil {
		log.Printf("Erro ao publicar a mensagem: %s", err)
	}

	return createdPatient, nil
}

func (service *patientService) UpdatePatient(id uuid.UUID, patient *Patient) (*Patient, error) {
	existingPatient, err := service.repository.Show(id)
	if err != nil {
		return nil, err
	}

	existingPatient.Name = patient.Name
	existingPatient.Age = patient.Age
	existingPatient.Gender = patient.Gender
	existingPatient.Address = patient.Address
	existingPatient.Phone = patient.Phone
	existingPatient.Email = patient.Email
	existingPatient.Observations = patient.Observations
	existingPatient.UpdatedAt = time.Now()

	if err := existingPatient.Validate(); err != nil {
		return nil, err
	}

	updatedPatient, err := service.repository.Update(existingPatient)
	if err != nil {
		return nil, err
	}

	message, _ := json.Marshal(updatedPatient)
	if err := service.rabbitMQ.Publish("patients_update", message); err != nil {
		log.Printf("Erro ao publicar a mensagem de atualização: %s", err)
	}

	return updatedPatient, nil
}

func (service *patientService) DeletePatient(id uuid.UUID) error {
	return service.repository.Delete(id)
}

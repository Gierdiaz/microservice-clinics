package patient

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PatientRepository interface {
	Index() ([]*Patient, error)
	Show(id uuid.UUID) (*Patient, error)
	Store(patient *Patient) (*Patient, error)
	Update(patient *Patient) (*Patient, error)
	Delete(id uuid.UUID) error
}

type patientRepository struct {
	db *sqlx.DB
}

func NewPatientRepository(db *sqlx.DB) PatientRepository {
	return &patientRepository{db: db}
}

func (repo *patientRepository) Index() ([]*Patient, error) {
	var patients []*Patient
	err := repo.db.Select(&patients, "SELECT * FROM patients")
	if err != nil {
		return nil, err
	}
	return patients, nil
}

func (repo *patientRepository) Show(id uuid.UUID) (*Patient, error) {
	var patient Patient
	err := repo.db.Get(&patient, "SELECT * FROM patients WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

func (repo *patientRepository) Store(patient *Patient) (*Patient, error) {
	_, err := repo.db.NamedExec(`
		INSERT INTO patients (id, name, age, gender, address, phone, email, observations, created_at, updated_at)
		VALUES (:id, :name, :age, :gender, :address, :phone, :email, :observations, :created_at, :updated_at)`,
		patient)
	if err != nil {
		return nil, err
	}
	return patient, nil
}

func (repo *patientRepository) Update(patient *Patient) (*Patient, error) {
	_, err := repo.db.NamedExec(`
		UPDATE patients SET name = :name, age = :age, gender = :gender, address = :address, phone = :phone, 
		email = :email, observations = :observations, updated_at = :updated_at WHERE id = :id`,
		patient)
	if err != nil {
		return nil, err
	}
	return patient, nil
}

func (repo *patientRepository) Delete(id uuid.UUID) error {
	_, err := repo.db.Exec("DELETE FROM patients WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}


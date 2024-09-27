package repository

import (
	"database/sql"
	"fmt"

	"github.com/Gierdiaz/diagier-clinics/internal/domain/patient"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type patientRepository struct {
	db *sqlx.DB
}

func NewPatientRepository(db *sqlx.DB) patient.PatientRepository {
	return &patientRepository{db: db}
}

func (repo *patientRepository) Index() ([]*patient.Patient, error) {
	var patients []*patient.Patient
	if err := repo.db.Select(&patients, "SELECT * FROM patients"); err != nil {
		return nil, err
	}
	return patients, nil
}

func (repo *patientRepository) Show(id uuid.UUID) (*patient.Patient, error) {
	var p patient.Patient
	err := repo.db.Get(&p, "SELECT * FROM patients WHERE id = $1", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("patient with id %s not found", id)
		}
		return nil, err
	}
	return &p, nil
}

func (repo *patientRepository) Store(p *patient.Patient) (*patient.Patient, error) {
	_, err := repo.db.NamedExec(`
		INSERT INTO patients (id, name, age, gender, address, phone, email, observations, created_at, updated_at) 
		VALUES (:id, :name, :age, :gender, :address, :phone, :email, :observations, NOW(), NOW())`, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (repo *patientRepository) Update(p *patient.Patient) (*patient.Patient, error) {
	_, err := repo.db.NamedExec(`
		UPDATE patients 
		SET name = :name, age = :age, gender = :gender, address = :address, phone = :phone, email = :email, observations = :observations, updated_at = NOW() 
		WHERE id = :id`, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (repo *patientRepository) Delete(id uuid.UUID) error {
	if _, err := repo.db.Exec("DELETE FROM patients WHERE id = $1", id); err != nil {
		return err
	}
	return nil
}

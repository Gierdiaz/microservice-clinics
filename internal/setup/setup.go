package setup

import (
	"github.com/Gierdiaz/diagier-clinics/internal/domain/patients"
	"github.com/Gierdiaz/diagier-clinics/internal/handler"
	"github.com/jmoiron/sqlx"
)

func SetupServices(db *sqlx.DB) *handler.PatientsHandler {
	patientRepo := patients.NewPatientRepository(db)
	patientService := patients.NewPatientService(patientRepo)
	patientHandler := handler.NewPatientsHandler(patientService)
	return patientHandler
}

package setup

import (
	"github.com/Gierdiaz/diagier-clinics/internal/domain/patient"
	"github.com/Gierdiaz/diagier-clinics/internal/domain/user"
	"github.com/Gierdiaz/diagier-clinics/internal/handler"
	"github.com/jmoiron/sqlx"
)

func SetupServices(db *sqlx.DB) *handler.PatientsHandler {
	patientRepo := patient.NewPatientRepository(db)
	patientService := patient.NewPatientService(patientRepo)
	patientHandler := handler.NewPatientsHandler(patientService)
	return patientHandler
}

func SetupUserServices(db *sqlx.DB) *handler.UserHandler {
    userRepo := user.NewUserRepository(db)
    userService := user.NewService(userRepo)
    userHandler := handler.NewUserHandler(userService)
    return userHandler
}

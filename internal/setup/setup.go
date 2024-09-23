package setup

import (
	"github.com/Gierdiaz/diagier-clinics/internal/domain/patient"
	"github.com/Gierdiaz/diagier-clinics/internal/domain/user"
	"github.com/Gierdiaz/diagier-clinics/internal/handler"
	"github.com/Gierdiaz/diagier-clinics/pkg/messaging"
	"github.com/jmoiron/sqlx"
)


func SetupServices(db *sqlx.DB, rabbit *messaging.RabbitMQ) *handler.PatientsHandler {
	patientRepo := patient.NewPatientRepository(db)
	patientService := patient.NewPatientService(patientRepo, rabbit)
	patientHandler := handler.NewPatientsHandler(patientService)
	return patientHandler
}

func SetupUserServices(db *sqlx.DB) *handler.UserHandler {
    userRepo := user.NewUserRepository(db)
    userService := user.NewService(userRepo)
    userHandler := handler.NewUserHandler(userService)
    return userHandler
}

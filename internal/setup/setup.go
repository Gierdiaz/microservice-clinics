package setup

import (
	patientRepo "github.com/Gierdiaz/diagier-clinics/infrastructure/repositories/patient"
	userRepo "github.com/Gierdiaz/diagier-clinics/infrastructure/repositories/user"
	"github.com/Gierdiaz/diagier-clinics/internal/domain/patient"
	"github.com/Gierdiaz/diagier-clinics/internal/domain/user"
	"github.com/Gierdiaz/diagier-clinics/internal/handler"
	"github.com/Gierdiaz/diagier-clinics/pkg/messaging"
	"github.com/jmoiron/sqlx"
)

func SetupServices(db *sqlx.DB, rabbit *messaging.RabbitMQ) *handler.PatientsHandler {
	patientRepository := patientRepo.NewPatientRepository(db)
	patientService := patient.NewPatientService(patientRepository, rabbit)
	patientHandler := handler.NewPatientsHandler(patientService)
	return patientHandler
}

func SetupUserServices(db *sqlx.DB) *handler.UserHandler {
	userRepository := userRepo.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	return userHandler
}

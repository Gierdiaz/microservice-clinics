package seeders

import (
	"log"

	"github.com/Gierdiaz/diagier-clinics/internal/domain/patient"
	"github.com/jmoiron/sqlx"
)

// RunSeeds - Função responsável por rodar todas as seeds
func RunSeeds(db *sqlx.DB) {
	patientRepo := patient.NewPatientRepository(db)

	if err := SeedPatients(patientRepo); err != nil {
		log.Fatalf("Erro ao rodar seeds de patients: %v", err)
		return
	}

	log.Println("Todas as seeds foram aplicadas com sucesso.")
}

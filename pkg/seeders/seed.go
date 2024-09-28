package seeders

import (
	"log"

	"github.com/Gierdiaz/diagier-clinics/infrastructure/repositories"
	"github.com/jmoiron/sqlx"
)

// RunSeeds - Função responsável por rodar todas as seeds
func RunSeeds(db *sqlx.DB) error {
	patientRepo := repositories.NewPatientRepository(db)

	if err := SeedPatients(patientRepo); err != nil {
		log.Fatalf("Erro ao rodar seeds de patients: %v", err)
		return err
	}

	log.Println("Todas as seeds foram aplicadas com sucesso.")
	return nil
}

package seeders

import (
	"fmt"
	"log"
	"time"

	"github.com/Gierdiaz/diagier-clinics/internal/domain/patient"

	"github.com/google/uuid"
)

func SeedPatients(repository patient.PatientRepository) error {

	existingPatients, err := repository.Index()
	if err != nil {
		return fmt.Errorf("erro ao verificar se há pacientes existentes: %v", err)
	}

	if len(existingPatients) > 0 {
		log.Println("A tabela de pacientes já está populada. Seed não será executado.")
		return nil
	}

	patients := []*patient.Patient{
		{
			ID:           uuid.New(),
			Name:         "John Doe",
			Age:          35,
			Gender:       "masculino",
			Address:      "123 Main Street",
			Phone:        "+5511999999999",
			Email:        "johndoe@example.com",
			Observations: "Paciente saudável, sem condições crônicas.",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			ID:           uuid.New(),
			Name:         "Jane Smith",
			Age:          28,
			Gender:       "feminino",
			Address:      "456 Oak Avenue",
			Phone:        "+5511988888888",
			Email:        "janesmith@example.com",
			Observations: "Paciente com alergias sazonais.",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			ID:           uuid.New(),
			Name:         "Alex Johnson",
			Age:          42,
			Gender:       "outro",
			Address:      "789 Pine Road",
			Phone:        "+5511977777777",
			Email:        "alexjohnson@example.com",
			Observations: "Histórico de hipertensão.",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	for _, p := range patients {
		_, err := repository.Store(p)
		if err != nil {
			log.Printf("Erro ao semear paciente %s: %v", p.Name, err)
			return fmt.Errorf("erro ao semear paciente %s: %v", p.Name, err)
		} else {
			log.Printf("Paciente %s semeado com sucesso!", p.Name)
		}
	}

	return nil
}

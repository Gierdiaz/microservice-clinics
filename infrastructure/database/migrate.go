package database

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func RunMigrate(db *sqlx.DB) error {
	log.Println("Iniciando as migrations...")
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("erro ao criar driver Postgres para migrations: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///app/infrastructure/database/migrations", "postgres", driver)
		
	if err != nil {
		return fmt.Errorf("erro ao carregar as migrations: %w", err)
	}
	
	log.Println("Aplicando as migrations...")
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {	
		return errors.Wrap(err, "Erro na migrate.go ao aplicar as migrations")
	}

	log.Println("Migrations aplicadas com sucesso")
	return nil
}

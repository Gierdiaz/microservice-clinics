package database

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func RunMigrate(db *sqlx.DB) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("erro ao criar driver Postgres para migrations: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///app/infrastructure/database/migrations", "postgres", driver)
	if err != nil {
		return fmt.Errorf("erro ao carregar as migrations: %w", err)
	}
	
	log.Printf("Tentando aplicar as migrations do diretório: %s", "file:///app/infrastructure/database/migrations")
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {	
		return errors.Wrap(err, "Erro ao aplicar as migrations")
	}

	log.Println("Migrations aplicadas com sucesso")
	return nil
}

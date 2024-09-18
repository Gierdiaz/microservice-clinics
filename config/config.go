package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Definindo a estrutura do servidor
type Server struct {
	APP_SERVER string
}

// Definindo a estrutura do banco de dados
type Database struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

// Estrutura principal de configuração
type Config struct {
	Server   Server
	Database Database
}

// Função para carregar o arquivo .env e preencher as configurações
func LoadConfig() (*Config, error) {
	// Carregar o arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Preenchendo a estrutura Config com as variáveis de ambiente
	config := &Config{
		Server: Server{
			APP_SERVER: os.Getenv("APP_SERVER"),
		},
		Database: Database{
			DB_HOST:     os.Getenv("DB_HOST"),
			DB_PORT:     os.Getenv("DB_PORT"),
			DB_USER:     os.Getenv("DB_USER"),
			DB_PASSWORD: os.Getenv("DB_PASSWORD"),
			DB_NAME:     os.Getenv("DB_NAME"),
		},
	}

	// Validar as configurações
	if err := validateConfig(config); err != nil {
		return nil, err
	}

	return config, nil
}

// Função para validar as configurações carregadas
func validateConfig(cfg *Config) error {
	// Validar o servidor
	if cfg.Server.APP_SERVER == "" {
		return errors.New("APP_SERVER not configured")
	}

	// Validar o banco de dados
	if cfg.Database.DB_HOST == "" {
		return errors.New("DB_HOST not configured")
	}
	if cfg.Database.DB_PORT == "" {
		return errors.New("DB_PORT not configured")
	}
	if cfg.Database.DB_USER == "" {
		return errors.New("DB_USER not configured")
	}
	if cfg.Database.DB_PASSWORD == "" {
		return errors.New("DB_PASSWORD not configured")
	}
	if cfg.Database.DB_NAME == "" {
		return errors.New("DB_NAME not configured")
	}

	return nil
}

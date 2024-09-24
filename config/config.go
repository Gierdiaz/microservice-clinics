package config

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Server struct {
	APP_SERVER string
}

type Database struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

type JWT struct {
	Secret   string
	ExpHours int
}

type RabbitMQ struct {
	URL string
}


// Estrutura principal de configuração
type Config struct {
	Server   Server
	Database Database
	JWT      JWT
	RabbitMQ RabbitMQ
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	expHours, err := strconv.Atoi(os.Getenv("ExpHours"))
	if err != nil {
		expHours = 24
	}

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
		JWT: JWT{
			Secret:   os.Getenv("JWT_SECRET"),
			ExpHours: expHours,
		},
		RabbitMQ: RabbitMQ{
			URL: os.Getenv("RABBITMQ_URL"),
		},
	}

	if err := validateConfig(config); err != nil {
		return nil, err
	}

	return config, nil
}

func validateConfig(cfg *Config) error {
	if cfg.Server.APP_SERVER == "" {
		return errors.New("APP_SERVER não configurado")
	}

	// Validar o banco de dados
	if cfg.Database.DB_HOST == "" {
		return errors.New("DB_HOST não configurado")
	}
	if cfg.Database.DB_PORT == "" {
		return errors.New("DB_PORT não configurado")
	}
	if cfg.Database.DB_USER == "" {
		return errors.New("DB_USER não configurado")
	}
	if cfg.Database.DB_PASSWORD == "" {
		return errors.New("DB_PASSWORD não configurado")
	}
	if cfg.Database.DB_NAME == "" {
		return errors.New("DB_NAME não configurado")
	}

	// Validar JWT
	if cfg.JWT.Secret == "" {
		return errors.New("JWT_SECRET não configurado")
	}

	// Validar RabbitMQ
	if cfg.RabbitMQ.URL == "" {
		return errors.New("RABBITMQ_URL não configurado")
	}

	return nil
}

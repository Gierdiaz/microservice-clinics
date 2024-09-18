package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &Config{
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
	}

	if err := validateConfig(config); err != nil {
		return nil, err
	}

	return config, nil
}

func validateConfig(cfg *Config) error {
	if cfg.DB_HOST == "" {
		return errors.New("DB_HOST not configured")
	}
	if cfg.DB_PORT == "" {
		return errors.New("DB_PORT not configured")
	}
	if cfg.DB_USER == "" {
		return errors.New("DB_USER not configured")
	}
	if cfg.DB_PASSWORD == "" {
		return errors.New("DB_PASSWORD not configured")
	}
	if cfg.DB_NAME == "" {
		return errors.New("DB_NAME not configured")
	}

	return nil
}
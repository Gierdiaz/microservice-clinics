package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Gierdiaz/diagier-clinics/config"
	"github.com/Gierdiaz/diagier-clinics/infrastructure/database"
	"github.com/Gierdiaz/diagier-clinics/internal/endpoint"
	"github.com/Gierdiaz/diagier-clinics/pkg/logger"
	"github.com/Gierdiaz/diagier-clinics/pkg/messaging"
	"github.com/Gierdiaz/diagier-clinics/pkg/middleware"
	"github.com/Gierdiaz/diagier-clinics/pkg/seeders"
	"github.com/Gierdiaz/diagier-clinics/pkg/validator"
)

func main() {
	// Inicializando o logger
	logger := logger.NewLogger()

	// Carregando as configurações
	config, err := config.LoadConfig()
	if err != nil {
		logger.Log("level", "error", "msg", "Erro ao carregar configurações", err)
		os.Exit(1)
	}
	logger.Log("level", "info", "msg", "Configurações carregadas com sucesso", "port: ", config.Database.DB_PORT)
	fmt.Printf("Conectando na porta %s...\n", config.Database.DB_PORT)

	// Inicializando o validador
	validator.InitValidator()

	// Initialize JWT middleware
	middleware.InitJWT(config)

	// Conectando ao banco de dados
	db, err := database.InitDatabase(config)
	if err != nil {
		logger.Log("level", "error", "msg", "Erro ao conectar ao banco de dados", err)
		os.Exit(1)
	}
	defer db.Close()

	logger.Log("level", "info", "msg", "Conexão com o banco de dados estabelecida")

	// Executando as migrations
	err = database.RunMigrate(db)
	if err != nil {
		logger.Log("level", "error", "msg", "EEEEEEEEErro ao aplicar as migrations", "error", err.Error())
		os.Exit(1)
	}

	logger.Log("level", "info", "msg", "Migrations aplicadas com sucesso")

	// Executando as seeds
	seeders.RunSeeds(db)

	// Inicializando a conexão RabbitMQ
	rabbitMQ, err := messaging.NewRabbitMQ(config.RabbitMQ.URL) // Use a URL de conexão do RabbitMQ aqui
	if err != nil {
		logger.Log("level", "info", "msg", "Conectando ao RabbitMQ", "url", config.RabbitMQ.URL)
		os.Exit(1)
	}
	defer rabbitMQ.Close() // Fechar a conexão ao finalizar o programa

	// Inicializando o roteador Gin
	router := endpoint.Router(db, rabbitMQ)

	// Rodando o servidor HTTP na porta 8080
	err = http.ListenAndServe(config.Server.APP_SERVER, router)
	if err != nil {
		logger.Log("level", "error", "msg", "Erro ao iniciar o servidor", err)
		os.Exit(1)
	}
}

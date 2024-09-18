package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Gierdiaz/diagier-clinics/config"
	"github.com/Gierdiaz/diagier-clinics/infrastructure/database"
	"github.com/Gierdiaz/diagier-clinics/internal/endpoint"
	"github.com/Gierdiaz/diagier-clinics/pkg/logger"
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

	// Conectando ao banco de dados
	db, err := database.InitDatabase(config)
	if err != nil {
		logger.Log("level", "error", "msg", "Erro ao conectar ao banco de dados", err)
		os.Exit(1)
	}
	defer db.Close()

	logger.Log("level", "info", "msg", "Conexão com o banco de dados estabelecida")

	// Inicializando o roteador Gin
	router := endpoint.Router()

	// Rodando o servidor HTTP na porta 8080
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		logger.Log("level", "error", "msg", "Erro ao iniciar o servidor", err)
		os.Exit(1)
	}
	

}
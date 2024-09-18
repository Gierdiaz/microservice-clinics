package main

import (
	"fmt"
	"os"

	"github.com/Gierdiaz/diagier-clinics/config"
	"github.com/Gierdiaz/diagier-clinics/pkg/logger"
)

func main() {

	logger := logger.NewLogger()

	config, err := config.LoadConfig()
	if err != nil {
		logger.Log("level", "error", "msg", "Erro ao carregar configurações", err)
		os.Exit(1)
	}
	logger.Log("level", "info", "msg", "Configurações carregadas com sucesso", "port: ", config.DB_PORT)
	fmt.Printf("Conectando na porta %s...\n", config.DB_PORT)
}
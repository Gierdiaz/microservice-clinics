# Nome do binário
BINARY_NAME=myapp

# Variáveis
DOCKER_COMPOSE = docker-compose
DOCKER_SERVICE = microservice_diagier_clinics
DOCKER_API = docker exec -it $(DOCKER_SERVICE)
USE_DOCKER = true
DOCKER_CMD = $(if $(USE_DOCKER),$(DOCKER_API),)

# Comando para rodar seeds
seed: ## Rodar as seeds
	$(DOCKER_CMD) go run cmd/seed/main.go

# Comando para rodar testes
test: ## Rodar os testes
	$(DOCKER_CMD) go test -v ./...

# Comando para rodar a aplicação
run: ## Rodar a aplicação
	$(DOCKER_COMPOSE) up

# Comando para rodar a aplicação em modo detach
run-detach: ## Rodar a aplicação em modo detach
	$(DOCKER_COMPOSE) up -d

# Comando para derrubar os serviços
down: ## Derrubar os serviços
	$(DOCKER_COMPOSE) down

# Comando para visualizar logs
logs: ## Visualizar os logs
	$(DOCKER_COMPOSE) logs -f

# Comando para compilar a aplicação
build: ## Compilar a aplicação
	$(DOCKER_CMD) go build -o $(BINARY_NAME) ./cmd/main.go

# Comando para limpar os binários
clean: ## Limpar os binários
	rm -f $(BINARY_NAME)


FROM golang:1.23-alpine AS builder

# Instala dependências essenciais
RUN apk update && apk add --no-cache git

# Cria o diretório de trabalho
WORKDIR /app

# Desativa o proxy Go
ENV GOPROXY=direct

# Copia os arquivos do projeto para dentro do contêiner
COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download

# Copia o código fonte
COPY . .

# Compila o binário
RUN go build -o main ./cmd/api

# Verifica se o binário foi gerado
RUN ls -la /app

# Etapa de execução
FROM alpine:latest

# Cria o diretório de trabalho no contêiner final
WORKDIR /app

# Copia o binário compilado da etapa de construção
COPY --from=builder /app/main /app/main

# Copia o arquivo .env para o contêiner
COPY .env /app/.env

# Expõe a porta que o serviço vai usar
EXPOSE 8080

# Definindo a variável de ambiente com a localização do arquivo .env
ENV GO_ENV=production

# Comando que será executado ao iniciar o contêiner
CMD ["/app/main"]


FROM golang:1.23-alpine AS builder

# Instala dependências essenciais
RUN apk update && apk add --no-cache git

# Cria o diretório de trabalho
WORKDIR /app

# Copia os arquivos do projeto para dentro do contêiner
COPY go.mod go.sum ./
RUN go mod download

# Copia o código fonte
COPY . .

# Compila o binário
RUN go build -o main ./cmd/api

# Etapa de execução
FROM alpine:latest

# Copia o binário compilado da etapa de construção
COPY --from=builder /app/main /app/main

# Define o diretório de trabalho no contêiner final
WORKDIR /app

# Expõe a porta que o serviço vai usar
EXPOSE 8080

# Comando que será executado ao iniciar o contêiner
CMD ["./main"]

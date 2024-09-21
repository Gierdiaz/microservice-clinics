# Comandos para Gerenciamento do Banco de Dados PostgreSQL

Este documento contém comandos úteis para gerenciar o banco de dados PostgreSQL no seu projeto.

## Acessando o Contêiner PostgreSQL

```bash
docker compose exec -u postgres diagier_clinics_db psql -d diagier-clinics
```

## Listar Bancos de Dados

```bash
\l
```

## Listar Tabelas

```bash
\dt
```

## Colunas de uma Tabela

```bash
\d patients
```

## Sair do Prompt do PostgreSQL
```bash
\q
```

-- Habilita a extensão para geração de UUIDs se não existir
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Remove a tabela users se já existir
DROP TABLE IF EXISTS users;

-- Criação da tabela users
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

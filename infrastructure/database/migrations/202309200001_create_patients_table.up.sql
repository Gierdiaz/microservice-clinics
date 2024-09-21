-- Habilita a extensão para geração de UUIDs se não existir
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Remove a tabela patients se já existir
DROP TABLE IF EXISTS patients;

-- Criação da tabela patients
CREATE TABLE patients (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- UUID gerado automaticamente
    name VARCHAR(100) NOT NULL,
    age INT NOT NULL CHECK (age > 0 AND age < 150),
    gender VARCHAR(10) NOT NULL CHECK (gender IN ('masculino', 'feminino', 'outro')),
    address VARCHAR(100) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    observations TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

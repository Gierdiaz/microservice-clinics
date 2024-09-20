-- Criação da extensão para gerar UUIDs
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

DROP TABLE IF EXISTS patients;

-- Criação da tabela patients
CREATE TABLE IF NOT EXISTS patients (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    age INT NOT NULL CHECK (
        age > 0
        AND age < 150
    ),
    gender VARCHAR(10) NOT NULL CHECK (gender IN ('male', 'female', 'other')),
    address VARCHAR(100) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    observations TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
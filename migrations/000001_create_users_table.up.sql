
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    phone_number VARCHAR(16) NOT NULL UNIQUE,

    username VARCHAR(32) NOT NULL UNIQUE,

    email VARCHAR(64) NOT NULL UNIQUE,

    is_verified BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_users_phone_number
ON users(phone_number);

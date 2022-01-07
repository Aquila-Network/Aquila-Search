CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- CREATE TABLE users
-- (
--     id              SERIAL PRIMARY KEY,
--     first_name      varchar(100) not null,
--     last_name       varchar(100) not null,
--     email           varchar(50) not null unique,
--     password        varchar(255) not null,
--     is_admin  BOOLEAN NOT NULL DEFAULT FALSE,
--     is_active BOOLEAN NOT NULL DEFAULT TRUE,
--     created_at TIMESTAMP NOT NULL DEFAULT NOW()
-- );

CREATE TABLE customers_temp
(
    id SERIAL PRIMARY KEY,
    customer_id uuid DEFAULT uuid_generate_v1() unique,
    first_name VARCHAR(20) NOT NULL,
    last_name VARCHAR(20),
    avatar BYTEA,
    secret_key VARCHAR(20) NOT NULL unique,
    aquila_db_database_name uuid DEFAULT uuid_generate_v1() unique,
    shared_hash uuid DEFAULT uuid_generate_v1() unique,
    is_sharable BOOLEAN NOT NULL DEFAULT TRUE,
    document_number INTEGER DEFAULT 0,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE customers
(
    id SERIAL PRIMARY KEY,
    customer_id uuid DEFAULT uuid_generate_v1() unique,
    first_name VARCHAR(20) NOT NULL,
    last_name VARCHAR(20) NOT NULL,
    email varchar(50) not null unique,
    description varchar(255),
    avatar BYTEA,
    secret_key VARCHAR(20) NOT NULL unique,
    aquila_db_database_name uuid DEFAULT uuid_generate_v1() unique,
    shared_hash uuid DEFAULT uuid_generate_v1() unique,
    is_sharable BOOLEAN NOT NULL DEFAULT TRUE,
    document_number INTEGER DEFAULT 0,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
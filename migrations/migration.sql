CREATE TABLE users
(
    id              SERIAL PRIMARY KEY,
    first_name      varchar(100) not null,
    last_name       varchar(100) not null,
    email           varchar(50) not null unique,
    password        varchar(255) not null,
    is_admin  BOOLEAN NOT NULL DEFAULT FALSE,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

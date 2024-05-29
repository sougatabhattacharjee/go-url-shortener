-- init.sql
CREATE TABLE IF NOT EXISTS urls
(
    id         SERIAL PRIMARY KEY,
    short_url  VARCHAR(255) NOT NULL UNIQUE,
    long_url   TEXT         NOT NULL,
    clicks     INT                   DEFAULT 0,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

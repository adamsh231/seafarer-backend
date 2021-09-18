-- +migrate Up
CREATE TABLE IF NOT EXISTS "admins"
(
    "id"          char(36) PRIMARY KEY DEFAULT (uuid_generate_v4()),
    "name"        varchar(255) NOT NULL,
    "email"       varchar(255) NOT NULL UNIQUE,
    "password"    varchar(255) NOT NULL,
    "company_id"    char(36)   NOT NULL,
    "created_at"  timestamp    NOT NULL,
    "updated_at"  timestamp    NOT NULL,
    "deleted_at"  timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "admins";
-- +migrate Up
CREATE TABLE IF NOT EXISTS "users"
(
    "id"          char(36) PRIMARY KEY DEFAULT (uuid_generate_v4()),
    "name"        varchar(255) NOT NULL,
    "email"       varchar(255) NOT NULL UNIQUE,
    "password"    varchar(255) NOT NULL,
    "created_at"  timestamp    NOT NULL,
    "is_verified" bool         NOT NULL,
    "updated_at"  timestamp    NOT NULL,
    "deleted_at"  timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "users";
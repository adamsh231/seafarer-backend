-- +migrate Up
CREATE TABLE IF NOT EXISTS "recruitments"
(
    "id"            char(36)     PRIMARY KEY DEFAULT (uuid_generate_v4()),
    "user_uuid"     char(36)     NOT NULL,
    "expect_salary" real         NOT NULL,
    "salary"        real,
    "position"      varchar(255),
    "available_on"  date,
    "sign_on"       date,
    "ship"          varchar(255),
    "letter"        varchar(255),
    "status"        varchar(255)
    );

-- +migrate Down
DROP TABLE IF EXISTS "recruitments";
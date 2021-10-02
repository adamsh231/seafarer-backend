-- +migrate Up
CREATE TABLE IF NOT EXISTS "recruitments"
(
    "id"           char(36)     PRIMARY KEY DEFAULT (uuid_generate_v4()),
    "user_uuid"    char(36)     NOT NULL,
    "salary"       float(11)    NOT NULL,
    "position"     varchar(255) NOT NULL,
    "available_on" date,
    "sign_on"      date,
    "ship"         varchar(255),
    "letter"       varchar(255),
    "status"       varchar(255)
    );

-- +migrate Down
DROP TABLE IF EXISTS "recruitments";
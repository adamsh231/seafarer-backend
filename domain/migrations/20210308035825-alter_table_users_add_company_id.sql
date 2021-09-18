-- +migrate Up
ALTER TABLE users
ADD COLUMN company_id char(36);

-- +migrate Down
ALTER TABLE users
DROP COLUMN company_id;
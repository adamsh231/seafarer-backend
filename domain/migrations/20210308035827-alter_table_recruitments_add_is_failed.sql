-- +migrate Up
ALTER TABLE recruitments
ADD COLUMN is_failed bool;

-- +migrate Down
ALTER TABLE recruitments
DROP COLUMN is_failed;
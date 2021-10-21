-- +migrate Up
ALTER TABLE users
ADD COLUMN "recruitment_id" char(36);

-- +migrate Down
ALTER TABLE users
DROP COLUMN recruitment_id;
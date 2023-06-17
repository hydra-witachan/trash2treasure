-- migrate:up
ALTER TABLE users
    ADD COLUMN points BIGINT NOT NULL DEFAULT '0';

-- migrate:down
ALTER TABLE users
    DROP COLUMN points;

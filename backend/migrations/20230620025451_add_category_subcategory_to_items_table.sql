-- migrate:up
ALTER TABLE items
ADD category VARCHAR(128) NOT NULL DEFAULT "",
ADD sub_category VARCHAR(128) NOT NULL DEFAULT "";

-- migrate:down
ALTER TABLE items
DROP COLUMN category,
DROP COLUMN sub_category;

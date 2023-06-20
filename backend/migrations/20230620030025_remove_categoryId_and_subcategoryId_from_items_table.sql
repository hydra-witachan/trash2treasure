-- migrate:up
ALTER TABLE items
DROP FOREIGN KEY fk_category_id,
DROP FOREIGN KEY fk_sub_category_id,
DROP COLUMN category_id,
DROP COLUMN sub_category_id;

-- migrate:down
ALTER TABLE items
ADD category_id CHAR(36);

ALTER TABLE items
ADD CONSTRAINT fk_category_id
FOREIGN KEY (category_id)
REFERENCES categories (id);

ALTER TABLE items
ADD sub_category_id CHAR(36);

ALTER TABLE items
ADD CONSTRAINT fk_sub_category_id
FOREIGN KEY (sub_category_id)
REFERENCES sub_categories (id);
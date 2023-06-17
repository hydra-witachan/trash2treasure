-- migrate:up
INSERT INTO categories (name)
VALUES ('Plastic'),
    ('Organic'),
    ('Paper'),
    ('Glass');

-- migrate:down


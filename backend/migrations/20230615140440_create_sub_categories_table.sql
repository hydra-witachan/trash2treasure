-- migrate:up
CREATE TABLE sub_categories (
    id CHAR(36) NOT NULL DEFAULT UUID(),

    category_id CHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    deleted_at TIMESTAMP,

    PRIMARY KEY (id),
    INDEX (deleted_at),
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);

-- migrate:down
DROP TABLE sub_categories;

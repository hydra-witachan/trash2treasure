-- migrate:up
CREATE TABLE categories (
    id CHAR(36) NOT NULL DEFAULT UUID(),

    name VARCHAR(255) NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    deleted_at TIMESTAMP,

    PRIMARY KEY (id),
    INDEX (deleted_at)
);

-- migrate:down
DROP TABLE categories;

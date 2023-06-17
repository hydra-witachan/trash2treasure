-- migrate:up
CREATE TABLE items (
    id CHAR(36) NOT NULL DEFAULT UUID(),

    author_id CHAR(36) NOT NULL,
    author_name VARCHAR(256) NOT NULL,
    item_name VARCHAR(256) NOT NULL,
    description VARCHAR(512),
    points INT NOT NULL,
    image_url VARCHAR(256),
    needed_amount INT NOT NULL,
    fullfiled_amount INT NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    deleted_at TIMESTAMP,

    PRIMARY KEY (id),
    INDEX (deleted_at),
    FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE
);

-- migrate:down
DROP TABLE items;

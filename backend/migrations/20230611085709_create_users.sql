-- migrate:up
CREATE TABLE users (
    id CHAR(36) NOT NULL DEFAULT UUID(),

    username VARCHAR(64) NOT NULL,
    full_name VARCHAR(256) NOT NULL,
    email VARCHAR(128) NOT NULL,
    address VARCHAR(512) NOT NULL,
    password VARCHAR(128) NOT NULL,
    role VARCHAR(12) NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    deleted_at TIMESTAMP,

    PRIMARY KEY (id),
    INDEX (deleted_at),
    CHECK (role IN ('donator', 'collector'))
);

-- migrate:down
DROP TABLE users;

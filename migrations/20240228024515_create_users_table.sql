-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    name       TEXT,
    email      varchar(255) NOT NULL,
    password   varchar(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    UNIQUE (email)
);
INSERT INTO users (name, email, password)
VALUES ('test', 'test@test.test', '$2a$10$Z.rlid5FtJlteNtKs9i6meTVgnj.gyhU8nNyeE54BNkH/ALkhfN0.');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd

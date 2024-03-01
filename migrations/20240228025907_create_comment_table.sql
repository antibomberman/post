-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS comments
(
    id         SERIAL PRIMARY KEY,
    user_id    INTEGER   NOT NULL,
    post_id    INTEGER   NOT NULL,
    Content       TEXT      NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
INSERT INTO comments (user_id, post_id, Content) VALUES (1, 1, 'This is a comment');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS comments;
-- +goose StatementEnd

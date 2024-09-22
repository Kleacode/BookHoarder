-- +goose Up
-- +goose StatementBegin
CREATE TABLE tags(
    id        SERIAL PRIMARY KEY,

    name      VARCHAR(50),
    user_id	INT REFERENCES users(id) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tags;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE books(
    id        SERIAL PRIMARY KEY,

    title     VARCHAR(100),
    user_id	INT REFERENCES users(id) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE books;
-- +goose StatementEnd

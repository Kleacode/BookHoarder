-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
    id        SERIAL PRIMARY KEY,

    name      VARCHAR(20)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd

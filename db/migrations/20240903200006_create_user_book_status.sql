-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_book_status(
    id        SERIAL PRIMARY KEY,

    user_id	INT REFERENCES users(id) NOT NULL,
    book_id	INT REFERENCES books(id) NOT NULL,
    status_id	INT REFERENCES status(id) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_book_status;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE book_tag(
    id SERIAL PRIMARY KEY,
    
    book_id	INT REFERENCES books(id) NOT NULL,
    tag_id	INT REFERENCES tags(id) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE book_tag;
-- +goose StatementEnd

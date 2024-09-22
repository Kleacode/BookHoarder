-- +goose Up
-- +goose StatementBegin
CREATE TABLE hoarder_tag(
    id SERIAL PRIMARY KEY,
    
    hoarder_id	INT REFERENCES user_book_status(id) NOT NULL,
    tag_id	INT REFERENCES tags(id) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE hoarder_tag;
-- +goose StatementEnd

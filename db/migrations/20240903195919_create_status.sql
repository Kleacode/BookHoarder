-- +goose Up
-- +goose StatementBegin
CREATE TABLE status(
    id        SERIAL PRIMARY KEY,

    name      VARCHAR(20)
);
INSERT INTO status (name) VALUES ('todo');
INSERT INTO status (name) VALUES ('wip');
INSERT INTO status (name) VALUES ('done');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE status;
-- +goose StatementEnd

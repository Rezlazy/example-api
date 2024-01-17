-- +goose Up
-- +goose StatementBegin
CREATE TABLE list
(
    id   BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE list;
-- +goose StatementEnd

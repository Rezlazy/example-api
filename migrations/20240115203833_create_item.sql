-- +goose Up
-- +goose StatementBegin
CREATE TABLE item
(
    id   BIGSERIAL PRIMARY KEY,
    title VARCHAR NOT NULL,
    description VARCHAR,
    list_id BIGINT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE item;
-- +goose StatementEnd

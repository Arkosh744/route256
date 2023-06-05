-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS items
(
    user_id SERIAL NOT NULL,
    sku     BIGINT NOT NULL,
    count   INT    NOT NULL,
    PRIMARY KEY (user_id, sku)
);

CREATE INDEX IF NOT EXISTS items_user_id_idx ON items (user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS items_user_id_idx;
DROP TABLE IF EXISTS items;
-- +goose StatementEnd
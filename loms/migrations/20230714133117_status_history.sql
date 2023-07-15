-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS msg_history
(
    user_id    BIGINT                      NOT NULL,
    order_id   BIGINT                      NOT NULL,
    status     TEXT                        NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);
Create INDEX IF NOT EXISTS msg_history_user_id_idx ON msg_history (user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS order_items_order_id_idx;
DROP TABLE IF EXISTS order_items;
-- +goose StatementEnd
-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders
(
    order_id   BIGSERIAL PRIMARY KEY,
    user_id    BIGINT    NOT NULL,
    status     text      NOT NULL default 'unknown',
    created_at TIMESTAMP NOT NULL default NOW()
);

CREATE TABLE IF NOT EXISTS reservations
(
    order_id     BIGINT  NOT NULL REFERENCES orders (order_id),
    sku          INTEGER NOT NULL,
    count        INTEGER NOT NULL,
    warehouse_id INTEGER NOT NULL,
    PRIMARY KEY (order_id, sku)
);

Create INDEX IF NOT EXISTS reservations_order_id_idx ON reservations (order_id);

CREATE TABLE IF NOT EXISTS order_items
(
    order_id     BIGINT  NOT NULL REFERENCES orders (order_id),
    sku          INTEGER NOT NULL,
    count        INTEGER NOT NULL,
    warehouse_id INTEGER NOT NULL,
    PRIMARY KEY (order_id, sku)
);

Create INDEX IF NOT EXISTS order_items_order_id_idx ON order_items (order_id);

CREATE TABLE IF NOT EXISTS stocks
(
    sku          INTEGER NOT NULL,
    warehouse_id INTEGER NOT NULL,
    count        INTEGER NOT NULL,
    PRIMARY KEY (sku, warehouse_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS stocks;
DROP INDEX IF EXISTS order_items_order_id_idx;
DROP TABLE IF EXISTS order_items;
DROP INDEX IF EXISTS reservations_order_id_idx;
DROP TABLE IF EXISTS reservations;
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
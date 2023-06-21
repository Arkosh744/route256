-- +goose Up
-- +goose StatementBegin

-- we can add user_id to the table to make it more granular
CREATE TABLE IF NOT EXISTS rate_limiter_data
(
    last_time  TIMESTAMP WITHOUT TIME ZONE,
    prev_count INT,
    cur_count  INT
);

INSERT INTO rate_limiter_data (last_time, prev_count, cur_count)
VALUES (now(), 0, 0)
ON CONFLICT DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS rate_limiter_data;
-- +goose StatementEnd

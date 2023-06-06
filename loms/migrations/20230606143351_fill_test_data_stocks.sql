-- +goose Up
-- +goose StatementBegin
INSERT INTO stocks VALUES (3596599, 1, 32);
INSERT INTO stocks VALUES (19366373, 3, 231);
INSERT INTO stocks VALUES (19366373, 8, 15);
INSERT INTO stocks VALUES (19366373, 13, 12);
INSERT INTO stocks VALUES (25475334, 11,33);
INSERT INTO stocks VALUES (24167411, 4,33);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM stocks;
-- +goose StatementEnd
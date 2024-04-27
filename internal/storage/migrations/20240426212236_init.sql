-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS wallets
(
    id   INTEGER PRIMARY KEY,
    user_name VARCHAR(32),
    balance   INT DEFAULT 0
);

CREATE TYPE OPERATION AS ENUM ('deposit', 'debit');

CREATE TABLE IF NOT EXISTS operations
(
    id        SERIAL PRIMARY KEY,
    user_id   INTEGER REFERENCES wallets (id),
    operation OPERATION,
    amount    INTEGER,
    reason    VARCHAR(100)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS operations;
DROP TABLE IF EXISTS wallets;
-- +goose StatementEnd

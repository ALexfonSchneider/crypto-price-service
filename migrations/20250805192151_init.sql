-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS coins
(
    id         TEXT PRIMARY KEY,
    name       TEXT        NOT NULL UNIQUE, -- например, "Bitcoin"
    symbol     TEXT        NOT NULL,        -- например, "BTC"
    is_active  BOOLEAN     NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_coins_symbol ON coins (symbol);

CREATE TABLE IF NOT EXISTS prices
(
    id        BIGSERIAL PRIMARY KEY,
    coin_id   TEXT      NOT NULL REFERENCES coins (id) ON DELETE CASCADE,
    price     FLOAT       NOT NULL,
    created_at TIMESTAMPTZ NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_prices_coin_id_timestamp ON prices (coin_id, created_at);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS coins;
DROP TABLE IF EXISTS prices;
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS coins (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,               -- например, "Bitcoin"
    symbol TEXT NOT NULL UNIQUE,      -- например, "BTC"
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_coins_symbol ON coins(symbol);

CREATE TABLE IF NOT EXISTS prices (
    id TEXT PRIMARY KEY,
    coin_id TEXT NOT NULL REFERENCES coins(id) ON DELETE CASCADE,
    price FLOAT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_prices_coin_id_timestamp ON prices(coin_id, timestamp);
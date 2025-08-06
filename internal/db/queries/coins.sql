-- name: CreateCoin :exec
INSERT INTO coins (id, name, symbol, is_active, created_at)
VALUES (@id, @name, @symbol, @is_active, @created_at);

-- name: GetCoin :one
SELECT id, name, symbol, is_active, created_at
FROM coins
WHERE id = @id;

-- name: GetCoinsBySymbols :many
SELECT id, name, symbol, is_active, created_at
FROM coins
WHERE symbol = ANY(@symbols::TEXT[]);

-- name: GetCoinBySymbol :one
SELECT id, name, symbol, is_active, created_at
FROM coins
WHERE symbol = @symbol;

-- name: ListActiveCoins :many
SELECT id, name, symbol, is_active, created_at
FROM coins
WHERE is_active = true
ORDER BY name;

-- name: DeactivateCoin :exec
UPDATE coins
SET is_active = false
WHERE id = @id;

-- name: ActivateCoin :exec
UPDATE coins
SET is_active = true
WHERE id = @id;
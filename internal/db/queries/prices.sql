-- name: CreatePrice :exec
INSERT INTO prices (coin_id, price, created_at)
VALUES (@coin_id, @price, @created_at);

-- name: CreatePrices :copyfrom
INSERT INTO prices (coin_id, price, created_at)
VALUES (@coin_id, @price, @created_at);

-- name: GetClosestPriceByCoinID :one
SELECT id, coin_id, price, created_at
FROM prices
WHERE coin_id = @coin_id
ORDER BY ABS(EXTRACT(EPOCH FROM (created_at - @created_at::timestamptz)))
LIMIT 1;

-- name: GetAllPricesForCoinByCoinID :many
SELECT id, coin_id, price, created_at
FROM prices
WHERE coin_id = @coin_id
ORDER BY created_at desc;
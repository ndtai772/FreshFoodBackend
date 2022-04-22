-- name: ListOrders :many
SELECT *
FROM orders
ORDER BY id DESC;

-- name: GetOrder :one
SELECT *
FROM orders
WHERE id = $1;

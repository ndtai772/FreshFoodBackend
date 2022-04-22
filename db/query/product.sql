-- name: ListProducts :many
SELECT *
FROM products
ORDER BY id DESC;

-- name: GetProduct :one
SELECT *
FROM products
WHERE id = $1;

-- name: CreateProduct :one
INSERT INTO products (
    title,
    price,
    description,
    image_0,
    image_1,
    image_2,
    image_3,
    valid_until
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING *;
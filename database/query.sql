-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- name: ListProducts :many
SELECT * FROM products
ORDER BY id;

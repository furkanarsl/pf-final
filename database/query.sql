-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- name: ListProducts :many
SELECT * FROM products
ORDER BY id;

-- name: ListProductsByID :many
SELECT * FROM products
WHERE id = ANY(@IDS::bigint[]);

-- name: GetCartForUser :one
select * from carts
where user_id = $1;

-- name: ListCartItems :many
select cp.id, p."name", p.price, p.vat, p.id as product_id from cart_products cp
left join carts c on cp.cart_id = c.id
left join products p on p.id = cp.product_id
where c.id = $1;

-- name: GetCartItem :one
SELECT * FROM cart_products
WHERE id = $1;

-- name: AddToCart :one
INSERT INTO cart_products(product_id,cart_id) values($1,$2) RETURNING *;

-- name: DeleteCartItem :exec
DELETE FROM cart_products
WHERE cart_id = $1 AND id = $2;

-- name: GetOrderForUser :one
SELECT * FROM orders
WHERE id = $1 AND user_id = $2;

-- name: GetOrdersForUser :many
SELECT * FROM orders
WHERE user_id = $1;

-- name: GetOrdersBetweenDate :many
SELECT * FROM orders
WHERE user_id = $1 AND ordered_at >= @start_date AND ordered_at < @end_date;

-- name: GetOrderCountBetweenDate :one
SELECT COUNT(*) FROM orders
WHERE user_id = $1 AND ordered_at >= @start_date AND ordered_at < @end_date;

-- name: GetOrderTotalBetweenDate :one
SELECT SUM(o.total_paid)::float FROM orders o
WHERE user_id = $1 AND ordered_at >= @start_date AND ordered_at < @end_date;

-- name: CreateOrder :one
INSERT INTO orders(user_id, ordered_at, total_paid) values($1,$2,$3)
RETURNING *;

-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- name: ListProducts :many
SELECT * FROM products
ORDER BY id;

-- name: ListProductsByID :many
SELECT * FROM products
WHERE id = ANY($1::bigint[]);

-- name: GetCartForUser :one
select * from carts
where user_id = $1;

-- name: ListCartItems :many
select cp.id, cp.quantity, p."name", p.price, p.vat from cart_products cp 
left join carts c on cp.cart_id = c.id
left join products p on p.id = cp.product_id
where c.id = $1;

-- name: GetCartItem :one
SELECT * FROM cart_products
WHERE id = $1;

-- name: AddToCart :one
INSERT INTO cart_products(product_id,cart_id,quantity) values($1,$2,$3) RETURNING *;

-- name: DeleteCartItem :exec
DELETE FROM cart_products
WHERE cart_id = $1 AND id = $2;

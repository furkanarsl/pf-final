// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: query.sql

package queries

import (
	"context"
	"database/sql"
	"time"
)

const addToCart = `-- name: AddToCart :one
INSERT INTO cart_products(product_id, cart_id, quantity) values($1, $2, $3) RETURNING id, product_id, cart_id, quantity
`

type AddToCartParams struct {
	ProductID int64
	CartID    int64
	Quantity  int32
}

func (q *Queries) AddToCart(ctx context.Context, arg AddToCartParams) (CartProduct, error) {
	row := q.db.QueryRow(ctx, addToCart, arg.ProductID, arg.CartID, arg.Quantity)
	var i CartProduct
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.CartID,
		&i.Quantity,
	)
	return i, err
}

const createOrder = `-- name: CreateOrder :one
INSERT INTO orders(user_id, ordered_at, total_paid) values($1,$2,$3)
RETURNING id, user_id, ordered_at, total_paid
`

type CreateOrderParams struct {
	UserID    int64
	OrderedAt time.Time
	TotalPaid float64
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRow(ctx, createOrder, arg.UserID, arg.OrderedAt, arg.TotalPaid)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.OrderedAt,
		&i.TotalPaid,
	)
	return i, err
}

const deleteCartItem = `-- name: DeleteCartItem :exec
DELETE FROM cart_products
WHERE cart_id = $1 AND id = $2
`

type DeleteCartItemParams struct {
	CartID int64
	ID     int64
}

func (q *Queries) DeleteCartItem(ctx context.Context, arg DeleteCartItemParams) error {
	_, err := q.db.Exec(ctx, deleteCartItem, arg.CartID, arg.ID)
	return err
}

const emptyCart = `-- name: EmptyCart :exec
DELETE FROM cart_products
WHERE cart_id = $1
`

func (q *Queries) EmptyCart(ctx context.Context, cartID int64) error {
	_, err := q.db.Exec(ctx, emptyCart, cartID)
	return err
}

const getCartForUser = `-- name: GetCartForUser :one
select id, user_id from carts
where user_id = $1
`

func (q *Queries) GetCartForUser(ctx context.Context, userID int64) (Cart, error) {
	row := q.db.QueryRow(ctx, getCartForUser, userID)
	var i Cart
	err := row.Scan(&i.ID, &i.UserID)
	return i, err
}

const getCartItem = `-- name: GetCartItem :one
SELECT id, product_id, cart_id, quantity FROM cart_products
WHERE id = $1
`

func (q *Queries) GetCartItem(ctx context.Context, id int64) (CartProduct, error) {
	row := q.db.QueryRow(ctx, getCartItem, id)
	var i CartProduct
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.CartID,
		&i.Quantity,
	)
	return i, err
}

const getCartItemByProductID = `-- name: GetCartItemByProductID :one
SELECT id, product_id, cart_id, quantity FROM cart_products
WHERE cart_id = $1 AND product_id = $2
`

type GetCartItemByProductIDParams struct {
	CartID    int64
	ProductID int64
}

func (q *Queries) GetCartItemByProductID(ctx context.Context, arg GetCartItemByProductIDParams) (CartProduct, error) {
	row := q.db.QueryRow(ctx, getCartItemByProductID, arg.CartID, arg.ProductID)
	var i CartProduct
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.CartID,
		&i.Quantity,
	)
	return i, err
}

const getOrderCountBetweenDate = `-- name: GetOrderCountBetweenDate :one
SELECT COUNT(*) FROM orders
WHERE user_id = $1 AND ordered_at >= $2 AND ordered_at < $3
`

type GetOrderCountBetweenDateParams struct {
	UserID    int64
	StartDate time.Time
	EndDate   time.Time
}

func (q *Queries) GetOrderCountBetweenDate(ctx context.Context, arg GetOrderCountBetweenDateParams) (int64, error) {
	row := q.db.QueryRow(ctx, getOrderCountBetweenDate, arg.UserID, arg.StartDate, arg.EndDate)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getOrderForUser = `-- name: GetOrderForUser :one
SELECT id, user_id, ordered_at, total_paid FROM orders
WHERE id = $1 AND user_id = $2
`

type GetOrderForUserParams struct {
	ID     int64
	UserID int64
}

func (q *Queries) GetOrderForUser(ctx context.Context, arg GetOrderForUserParams) (Order, error) {
	row := q.db.QueryRow(ctx, getOrderForUser, arg.ID, arg.UserID)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.OrderedAt,
		&i.TotalPaid,
	)
	return i, err
}

const getOrderTotalBetweenDate = `-- name: GetOrderTotalBetweenDate :one
SELECT SUM(o.total_paid)::float FROM orders o
WHERE user_id = $1 AND ordered_at >= $2 AND ordered_at < $3
`

type GetOrderTotalBetweenDateParams struct {
	UserID    int64
	StartDate time.Time
	EndDate   time.Time
}

func (q *Queries) GetOrderTotalBetweenDate(ctx context.Context, arg GetOrderTotalBetweenDateParams) (float64, error) {
	row := q.db.QueryRow(ctx, getOrderTotalBetweenDate, arg.UserID, arg.StartDate, arg.EndDate)
	var column_1 float64
	err := row.Scan(&column_1)
	return column_1, err
}

const getOrdersBetweenDate = `-- name: GetOrdersBetweenDate :many
SELECT id, user_id, ordered_at, total_paid FROM orders
WHERE user_id = $1 AND ordered_at >= $2 AND ordered_at < $3
`

type GetOrdersBetweenDateParams struct {
	UserID    int64
	StartDate time.Time
	EndDate   time.Time
}

func (q *Queries) GetOrdersBetweenDate(ctx context.Context, arg GetOrdersBetweenDateParams) ([]Order, error) {
	rows, err := q.db.Query(ctx, getOrdersBetweenDate, arg.UserID, arg.StartDate, arg.EndDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.OrderedAt,
			&i.TotalPaid,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOrdersForUser = `-- name: GetOrdersForUser :many
SELECT id, user_id, ordered_at, total_paid FROM orders
WHERE user_id = $1
`

func (q *Queries) GetOrdersForUser(ctx context.Context, userID int64) ([]Order, error) {
	rows, err := q.db.Query(ctx, getOrdersForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.OrderedAt,
			&i.TotalPaid,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProduct = `-- name: GetProduct :one
SELECT id, name, price, vat FROM products
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProduct(ctx context.Context, id int64) (Product, error) {
	row := q.db.QueryRow(ctx, getProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.Vat,
	)
	return i, err
}

const listCartItems = `-- name: ListCartItems :many
select cp.id, cp.quantity, p."name", p.price, p.vat, p.id as product_id from cart_products cp
left join carts c on cp.cart_id = c.id
left join products p on p.id = cp.product_id
where c.id = $1
`

type ListCartItemsRow struct {
	ID        int64
	Quantity  int32
	Name      sql.NullString
	Price     sql.NullFloat64
	Vat       sql.NullInt16
	ProductID sql.NullInt64
}

func (q *Queries) ListCartItems(ctx context.Context, id int64) ([]ListCartItemsRow, error) {
	rows, err := q.db.Query(ctx, listCartItems, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListCartItemsRow
	for rows.Next() {
		var i ListCartItemsRow
		if err := rows.Scan(
			&i.ID,
			&i.Quantity,
			&i.Name,
			&i.Price,
			&i.Vat,
			&i.ProductID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listProducts = `-- name: ListProducts :many
SELECT id, name, price, vat FROM products
ORDER BY id
`

func (q *Queries) ListProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.Query(ctx, listProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Price,
			&i.Vat,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listProductsByID = `-- name: ListProductsByID :many
SELECT id, name, price, vat FROM products
WHERE id = ANY($1::bigint[])
`

func (q *Queries) ListProductsByID(ctx context.Context, ids []int64) ([]Product, error) {
	rows, err := q.db.Query(ctx, listProductsByID, ids)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Price,
			&i.Vat,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCartItemQuantity = `-- name: UpdateCartItemQuantity :one
UPDATE cart_products
SET quantity = $1
WHERE id = $2 RETURNING id, product_id, cart_id, quantity
`

type UpdateCartItemQuantityParams struct {
	Quantity int32
	ID       int64
}

func (q *Queries) UpdateCartItemQuantity(ctx context.Context, arg UpdateCartItemQuantityParams) (CartProduct, error) {
	row := q.db.QueryRow(ctx, updateCartItemQuantity, arg.Quantity, arg.ID)
	var i CartProduct
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.CartID,
		&i.Quantity,
	)
	return i, err
}

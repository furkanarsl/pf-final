// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package queries

import (
	"time"
)

type Cart struct {
	ID     int64
	UserID int64
}

type CartProduct struct {
	ID        int64
	ProductID int64
	CartID    int64
	Quantity  int32
}

type Order struct {
	ID        int64
	UserID    int64
	OrderedAt time.Time
	TotalPaid float64
}

type Product struct {
	ID    int64
	Name  string
	Price float64
	Vat   int16
}

type User struct {
	ID   int64
	Name string
}

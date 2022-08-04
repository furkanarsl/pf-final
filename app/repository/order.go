package repository

import (
	"context"
	"time"

	"github.com/furkanarsl/pf-final/database"
	"github.com/furkanarsl/pf-final/pkg/queries"
)

type orderRepo struct {
	database.DbQueries
}

type OrderRepo interface {
	CreateOrder(userID int64, totalPaid float64) (queries.Order, error)
	CustomerOrderTotalMonthly(userID int64) float64
	CustomerOrderCountMonthly(userID int64) int64
}

func NewOrderRepo(queries database.DbQueries) *orderRepo {
	return &orderRepo{queries}
}

func (r *orderRepo) CreateOrder(userID int64, totalPaid float64) (queries.Order, error) {
	orderedAt := time.Now()
	args := queries.CreateOrderParams{UserID: userID, OrderedAt: orderedAt, TotalPaid: totalPaid}
	order, err := r.Queries.CreateOrder(context.Background(), args)

	if err != nil {
		return order, err
	}
	return order, nil
}

func (r *orderRepo) CustomerOrderTotalMonthly(userID int64) float64 {
	startOfMonth, endOfMonth := getCurrentMonth()
	args := queries.GetOrderTotalBetweenDateParams{StartDate: startOfMonth, EndDate: endOfMonth}
	result, _ := r.GetOrderTotalBetweenDate(context.Background(), args)
	return result
}

func (r *orderRepo) CustomerOrderCountMonthly(userID int64) int64 {
	startOfMonth, endOfMonth := getCurrentMonth()
	args := queries.GetOrderCountBetweenDateParams{StartDate: startOfMonth, EndDate: endOfMonth}
	result, _ := r.GetOrderCountBetweenDate(context.Background(), args)
	return result
}

func getCurrentMonth() (startOfMonth time.Time, endOfMonth time.Time) {
	now := time.Now()
	year, month, _ := now.Date()
	loc := now.Location()
	startOfMonth = time.Date(year, month, 1, 0, 0, 0, 0, loc)
	endOfMonth = startOfMonth.AddDate(0, 1, -1)
	return
}

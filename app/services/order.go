package services

import (
	"github.com/furkanarsl/pf-final/app/repository"
	"github.com/furkanarsl/pf-final/pkg/queries"
)

type OrderSvc struct {
	orderRepo   repository.OrderRepo
	cartService CartService
}

type OrderService interface {
	CompleteOrder(userID int64) (queries.Order, error)
}

func NewOrderService(orderRepo repository.OrderRepo, cartService CartService) *OrderSvc {
	return &OrderSvc{orderRepo: orderRepo, cartService: cartService}
}

func (s *OrderSvc) CompleteOrder(userID int64) (queries.Order, error) {
	// Get total of cart for user and check for empty cart etc.
	// empty the cart
	order, err := s.orderRepo.CreateOrder(userID, 10.10)
	if err != nil {
		return order, err
	}
	return order, nil
}

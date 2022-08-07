package services

import (
	"errors"

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
	cart, err := s.cartService.ListCart(userID)
	if err != nil {
		return queries.Order{}, err
	}
	if cart.CartSummary.FinalPrice < 1 {
		return queries.Order{}, errors.New("cannot complete order with empty cart")
	}
	order, err := s.orderRepo.CreateOrder(userID, cart.CartSummary.FinalPrice)

	if err != nil {
		return order, err
	}
	s.cartService.ClearCart(userID)
	return order, nil
}

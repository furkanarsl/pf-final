package services

import (
	"errors"

	"github.com/furkanarsl/pf-final/app/entity"
	"github.com/furkanarsl/pf-final/app/repository"
)

type OrderSvc struct {
	orderRepo   repository.OrderRepo
	cartService CartService
}

type OrderService interface {
	CompleteOrder(userID int64) (entity.Order, error)
}

func NewOrderService(orderRepo repository.OrderRepo, cartService CartService) *OrderSvc {
	return &OrderSvc{orderRepo: orderRepo, cartService: cartService}
}

func (s *OrderSvc) CompleteOrder(userID int64) (entity.Order, error) {
	cart, err := s.cartService.ListCart(userID)
	if err != nil {
		return entity.Order{}, err
	}
	if cart.CartSummary.FinalPrice < 1 {
		return entity.Order{}, errors.New("cannot complete order with empty cart")
	}
	order, err := s.orderRepo.CreateOrder(userID, cart.CartSummary.FinalPrice)

	if err != nil {
		return entity.Order{}, err
	}
	s.cartService.ClearCart(userID)
	result := entity.Order{
		ID:        order.ID,
		UserID:    order.UserID,
		OrderedAt: order.OrderedAt,
		Total:     order.TotalPaid,
	}

	return result, nil
}

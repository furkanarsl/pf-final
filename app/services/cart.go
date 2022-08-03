package services

import (
	"github.com/furkanarsl/pf-final/app/entity"
	"github.com/furkanarsl/pf-final/app/repository"
	"github.com/furkanarsl/pf-final/pkg/queries"
)

type CartSvc struct {
	CartRepo    repository.CartRepo
	ProductRepo repository.ProductRepo
}

type CartService interface {
	ListCart(userID int64) (entity.UserCart, error)
	AddToCart(userID int64, productID int64, quantity int32) (queries.CartProduct, error)
}

func NewCartService(CartRepo repository.CartRepo, ProductRepo repository.ProductRepo) *CartSvc {
	return &CartSvc{CartRepo: CartRepo, ProductRepo: ProductRepo}
}

func (s *CartSvc) ListCart(userID int64) (entity.UserCart, error) {
	userCart := entity.UserCart{}
	cart, err := s.CartRepo.GetCartForUser(userID)
	if err != nil {
		return userCart, err
	}
	userCart.ID = cart.ID

	cartItems, _ := s.CartRepo.GetCartItems(userID)

	if len(cartItems) < 1 {
		userCart.Items = []entity.CartItem{}
		return userCart, nil
	}
	var cartTotal float64 = 0
	var taxTotal float64 = 0
	for i := range cartItems {
		// TODO: Move these calculations to a different function to make it easier to also calculate discounts
		item := cartItems[i]
		taxAmount := item.Price.Float64 * float64(item.Vat.Int16) / 100
		taxedPrice := item.Price.Float64 + taxAmount
		itemTotal := float64(item.Quantity) * taxedPrice

		cartItem := entity.CartItem{
			ID:        item.ID,
			OrgPrice:  item.Price.Float64,
			Price:     taxedPrice,
			Quantity:  item.Quantity,
			Vat:       item.Vat.Int16,
			TaxAmount: taxAmount,
			Total:     itemTotal}
		userCart.Items = append(userCart.Items, cartItem)
		cartTotal += itemTotal
		taxTotal += taxAmount
	}

	userCart.TotalPrice = cartTotal
	userCart.TotalTax = taxTotal
	return userCart, nil
}

func (s *CartSvc) AddToCart(userID int64, productID int64, quantity int32) (queries.CartProduct, error) {
	cart, err := s.CartRepo.GetCartForUser(userID)
	if err != nil {
		return queries.CartProduct{}, err
	}

	product, err := s.ProductRepo.FindOne(productID)

	if err != nil {
		return queries.CartProduct{}, err
	}

	result, err := s.CartRepo.AddToCart(cart.ID, product.ID, quantity)
	if err != nil {
		return result, err
	}

	return result, nil
}

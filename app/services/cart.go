package services

import (
	"github.com/furkanarsl/pf-final/app/entity"
	"github.com/furkanarsl/pf-final/app/repository"
	"github.com/furkanarsl/pf-final/pkg/queries"
)

type cartSvc struct {
	CartRepo    repository.CartRepo
	ProductRepo repository.ProductRepo
}

type CartService interface {
	ListCart(userID int64) (entity.UserCart, error)
	AddToCart(userID int64, productID int64) (queries.CartProduct, error)
	RemoveFromCart(userID, itemID int64) error
}

func NewCartService(CartRepo repository.CartRepo, ProductRepo repository.ProductRepo) *cartSvc {
	return &cartSvc{CartRepo: CartRepo, ProductRepo: ProductRepo}
}

func (s *cartSvc) ListCart(userID int64) (entity.UserCart, error) {
	userCart := entity.UserCart{}
	cart, err := s.CartRepo.GetCartForUser(userID)
	if err != nil {
		return userCart, err
	}
	userCart.ID = cart.ID

	cartItems, _ := s.CartRepo.GetCartItems(userID)
	s.calculateCartItems(&cartItems, &userCart)

	if len(cartItems) < 1 {
		userCart.Items = []entity.CartItem{}
		return userCart, nil
	}
	return userCart, nil
}

func (s *cartSvc) AddToCart(userID int64, productID int64) (queries.CartProduct, error) {
	cart, err := s.CartRepo.GetCartForUser(userID)
	if err != nil {
		return queries.CartProduct{}, err
	}

	product, err := s.ProductRepo.FindOne(productID)

	if err != nil {
		return queries.CartProduct{}, err
	}

	result, err := s.CartRepo.AddToCart(cart.ID, product.ID)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *cartSvc) RemoveFromCart(userID, itemID int64) error {
	cart, err := s.CartRepo.GetCartForUser(userID)
	if err != nil {
		return err
	}

	err = s.CartRepo.RemoveFromCart(cart.UserID, itemID)
	if err != nil {
		return err
	}
	return nil
}

func (s *cartSvc) calculateCartItems(cartItems *[]queries.ListCartItemsRow, userCart *entity.UserCart) {
	var cartTotal float64 = 0
	var cartTaxTotal float64 = 0

	for i := range *cartItems {
		item := (*cartItems)[i]

		taxAmount := calculatePercent(item.Price.Float64, item.Vat.Int16)
		taxedPrice := item.Price.Float64 + taxAmount

		product := entity.Product{ID: item.ProductID.Int64, Name: item.Name.String, Vat: item.Vat.Int16, Price: item.Price.Float64}

		cartItem := entity.CartItem{
			ID:        item.ID,
			Product:   product,
			Price:     taxedPrice,
			DiscPrice: taxedPrice,
			OrgTax:    taxAmount,
			DiscTax:   taxAmount,
		}

		userCart.Items = append(userCart.Items, cartItem)
		cartTotal += cartItem.Price
		cartTaxTotal += taxAmount
	}
	userCart.TotalTax = cartTaxTotal
	userCart.TotalTaxDisc = cartTaxTotal
	userCart.TotalPrice = cartTotal
	userCart.TotalPriceDisc = cartTotal
}

func calculatePercent(price float64, percent int16) float64 {
	return (price * (float64(percent) / 100))
}

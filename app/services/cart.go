package services

import (
	"github.com/furkanarsl/pf-final/app/entity"
	"github.com/furkanarsl/pf-final/app/repository"
	"github.com/furkanarsl/pf-final/app/utils"
	"github.com/furkanarsl/pf-final/pkg/queries"
)

type cartSvc struct {
	CartRepo        repository.CartRepo
	ProductRepo     repository.ProductRepo
	OrderRepo       repository.OrderRepo
	DiscountService DiscountService
}

type CartService interface {
	ListCart(userID int64) (entity.UserCart, error)
	AddToCart(userID, productID int64, quantity int32) (entity.CartAddResult, error)
	ClearCart(userID int64) error
	RemoveFromCart(userID, itemID int64) error
}

func NewCartService(cartRepo repository.CartRepo, productRepo repository.ProductRepo, orderRepo repository.OrderRepo, discountService DiscountService) *cartSvc {
	return &cartSvc{
		CartRepo:        cartRepo,
		ProductRepo:     productRepo,
		DiscountService: discountService,
		OrderRepo:       orderRepo,
	}
}

func (s *cartSvc) ListCart(userID int64) (entity.UserCart, error) {
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

	s.calculateCartItems(&cartItems, &userCart)
	args := DiscountConditions{
		CustomerTotalMonthly:         s.OrderRepo.CustomerOrderTotalMonthly(userID),
		CustomerPurchaseCountMonthly: s.OrderRepo.CustomerOrderCountMonthly(userID),
	}
	userCart = s.DiscountService.ApplyDiscount(userCart, args)
	return userCart, nil
}

func (s *cartSvc) AddToCart(userID, productID int64, quantity int32) (entity.CartAddResult, error) {
	cart, err := s.CartRepo.GetCartForUser(userID)
	if err != nil {
		return entity.CartAddResult{}, err
	}

	product, err := s.ProductRepo.FindOne(productID)
	if err != nil {
		return entity.CartAddResult{}, err
	}

	cartItem, err := s.CartRepo.GetCartItemByProductID(cart.ID, product.ID)

	if err != nil {
		cartItem, err = s.CartRepo.AddToCart(cart.ID, product.ID, quantity)
	} else {
		cartItem, err = s.CartRepo.UpdateCartItemQuantity(cartItem, cartItem.Quantity+quantity)
	}

	r := entity.CartAddResult{}
	if err != nil {
		return r, err
	}

	r.Product = entity.Product(product)
	r.CartID = cartItem.CartID
	r.ID = cartItem.ID

	return r, nil
}

func (s *cartSvc) ClearCart(userID int64) error {
	cart, err := s.CartRepo.GetCartForUser(userID)
	if err != nil {
		return err
	}

	s.CartRepo.EmptyCart(cart.ID)
	return nil
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
	summary := entity.CartSummary{}

	for i := range *cartItems {
		item := (*cartItems)[i]

		taxAmount := utils.CalculatePercent(item.Price.Float64, item.Vat.Int16)
		taxedPrice := item.Price.Float64 + taxAmount
		product := entity.Product{ID: item.ProductID.Int64, Name: item.Name.String, Vat: item.Vat.Int16, Price: item.Price.Float64}

		cartItem := entity.CartItem{
			ID:         item.ID,
			Quantity:   item.Quantity,
			Product:    product,
			Price:      taxedPrice,
			OrgTax:     taxAmount,
			TotalPrice: float64(item.Quantity) * taxedPrice,
		}

		userCart.Items = append(userCart.Items, cartItem)
		cartTotal += cartItem.TotalPrice
		cartTaxTotal += taxAmount * float64(item.Quantity)
	}

	summary.ProductTotal = cartTotal
	summary.TaxTotal = cartTaxTotal
	summary.FinalPrice = summary.ProductTotal
	userCart.CartSummary = summary
}

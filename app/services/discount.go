package services

import (
	"github.com/furkanarsl/pf-final/app/entity"
)

// Discount amounts for
var discountAmount = map[int16]int16{
	0:  0,
	8:  10,
	18: 15,
}

type DiscountSvc struct {
}

type DiscountService interface {
	ApplyDiscount(userCart entity.UserCart)
}

func NewDiscountService() *DiscountSvc {
	return &DiscountSvc{}
}

func (s *DiscountSvc) ApplyDiscount(userCart entity.UserCart) {
	//TODO: Choose which discount is best for given cart and apply it

}

func (s *DiscountSvc) calculateFourthOrderDiscount(userCart entity.UserCart, discountThreshold float64, orderCount int) entity.UserCart {
	if orderCount%4 != 0 || userCart.TotalPrice < discountThreshold {
		return userCart
	}

	userCart.TotalPriceDisc = 0
	userCart.TotalTaxDisc = 0

	for i := range userCart.Items {
		var discount int16 = 0
		item := &userCart.Items[i]
		if val, ok := discountAmount[item.Product.Vat]; ok {
			discount = val
		}
		applyDiscount(item, &userCart, discount)
	}
	return userCart
}
func (s *DiscountSvc) calculateFourthItemDiscount(userCart entity.UserCart) {

}

func (s *DiscountSvc) calculateMonthlyDiscount(userCart entity.UserCart) {

}

func applyDiscount(item *entity.CartItem, userCart *entity.UserCart, discount int16) {
	discOrgPrice := item.Product.Price - calculatePercent(item.Product.Price, discount)
	item.DiscTax = calculatePercent(discOrgPrice, item.Product.Vat)
	item.DiscPrice = discOrgPrice + item.DiscTax
	userCart.TotalPriceDisc += item.DiscPrice
	userCart.TotalTaxDisc += item.DiscTax
}

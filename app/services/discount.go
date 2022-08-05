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

}

func (s *DiscountSvc) calculateFourthOrderDiscount(userCart entity.UserCart, discountThreshold float64, orderCount int) entity.UserCart {

	if orderCount%4 != 0 || userCart.TotalPrice < discountThreshold {
		return userCart
	}
	userCart.TotalPriceDisc = 0

	for i := range userCart.Items {
		var discount int16 = 0
		item := &userCart.Items[i]
		if val, ok := discountAmount[item.Vat]; ok {
			discount = val
		}
		applyDiscount(item, discount)
		userCart.TotalPriceDisc += item.DiscTotal
	}
	return userCart
}
func (s *DiscountSvc) calculateFourthItemDiscount(userCart entity.UserCart) {

}

func (s *DiscountSvc) calculateMonthlyDiscount(userCart entity.UserCart) {

}

func applyDiscount(item *entity.CartItem, discount int16) {
	item.DiscOrgPrice = item.OrgPrice - calculatePercent(item.OrgPrice, discount)
	item.DiscTax = calculatePercent(item.DiscOrgPrice, item.Vat)
	item.DiscPrice = item.DiscOrgPrice + item.DiscTax
	item.DiscTotal = item.DiscPrice * float64(item.Quantity)
}

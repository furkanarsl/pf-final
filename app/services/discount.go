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

func (s *DiscountSvc) calculateFourthItemDiscount(userCart entity.UserCart) entity.UserCart {
	var discount int16 = 0
	productCount := make(map[int64]int)
	userCart.TotalPriceDisc = 0
	userCart.TotalTaxDisc = 0
	for i := range userCart.Items {
		item := &userCart.Items[i]
		// Add product to count map
		if _, ok := productCount[item.Product.ID]; !ok {
			productCount[item.Product.ID] = 1
		} else {
			productCount[item.Product.ID] += 1
		}
		if productCount[item.Product.ID] > 3 {
			discount = 8
		}
		applyDiscount(item, &userCart, discount)
		discount = 0
	}
	return userCart
}

func (s *DiscountSvc) calculateMonthlyDiscount(userCart entity.UserCart) entity.UserCart {
	// TODO: add check to see if user made purchase more than given amount
	var discount int16 = 10
	userCart.TotalPriceDisc = 0
	userCart.TotalTaxDisc = 0
	for i := range userCart.Items {
		item := &userCart.Items[i]
		applyDiscount(item, &userCart, discount)
	}
	return userCart
}

func applyDiscount(item *entity.CartItem, userCart *entity.UserCart, discount int16) {
	discOrgPrice := item.Product.Price - calculatePercent(item.Product.Price, discount)
	item.DiscTax = calculatePercent(discOrgPrice, item.Product.Vat)
	item.DiscPrice = discOrgPrice + item.DiscTax
	userCart.TotalPriceDisc += item.DiscPrice
	userCart.TotalTaxDisc += item.DiscTax
}

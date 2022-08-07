package services

import (
	"github.com/furkanarsl/pf-final/app/entity"
	"github.com/furkanarsl/pf-final/app/utils"
)

type discountSvc struct {
	discountThreshold float64
}

type DiscountService interface {
	ApplyDiscount(userCart entity.UserCart)
}

func NewDiscountService(discountThreshold float64) *discountSvc {
	return &discountSvc{discountThreshold: discountThreshold}
}

type DiscountConditions struct {
	CustomerTotalMonthly         float64
	CustomerPurchaseCountMonthly int
}

func (s *discountSvc) ApplyDiscount(userCart entity.UserCart, args DiscountConditions) entity.UserCart {
	discounters := []Discounter{
		MonthlyDiscount{DiscountThreshold: s.discountThreshold, CustomerTotal: args.CustomerTotalMonthly},
		FourthPurchaseMonthlyDiscount{DiscountThreshold: s.discountThreshold, OrderCount: args.CustomerPurchaseCountMonthly},
		FourthItemDiscount{},
	}

	discountedSummaries := []entity.CartSummary{}

	for _, discounter := range discounters {

		result := discounter.CalculateCart(userCart)

		summary := entity.CartSummary{
			ProductTotal:   userCart.CartSummary.ProductTotal,
			DiscountAmount: result.DiscountAmount,
			TaxTotal:       result.TaxAmount,
			FinalPrice:     userCart.CartSummary.ProductTotal - result.DiscountAmount + (userCart.CartSummary.TaxTotal - result.TaxAmount),
		}

		discountedSummaries = append(discountedSummaries, summary)
	}

	for _, summ := range discountedSummaries {
		if summ.FinalPrice < userCart.CartSummary.FinalPrice {
			userCart.CartSummary = summ
		}
	}

	return userCart
}

// Discount amounts for
var discountAmount = map[int16]int16{
	0:  0,
	8:  10,
	18: 15,
}

type Discounter interface {
	CalculateCart(userCart entity.UserCart) DiscountResult
}

type DiscountResult struct {
	DiscountAmount float64
	TaxAmount      float64
}

type FourthPurchaseMonthlyDiscount struct {
	DiscountThreshold float64
	OrderCount        int
}

func (d FourthPurchaseMonthlyDiscount) CalculateCart(userCart entity.UserCart) DiscountResult {

	if d.OrderCount%4 != 0 || userCart.CartSummary.ProductTotal < d.DiscountThreshold {
		return DiscountResult{}
	}

	var discountPercent int16 = 0
	var discountResult float64 = 0
	var taxResult float64 = 0

	for i := range userCart.Items {
		item := &userCart.Items[i]
		if val, ok := discountAmount[item.Product.Vat]; ok {
			discountPercent = val
			discount := calculateDiscount(item, discountPercent)
			newTax := utils.CalculatePercent(item.Product.Price-discount, item.Product.Vat)
			discountResult += discount * float64(item.Quantity)

			taxResult += newTax * float64(item.Quantity)
		}
	}
	return DiscountResult{DiscountAmount: discountResult, TaxAmount: taxResult}
}

type FourthItemDiscount struct{}

func (FourthItemDiscount) CalculateCart(userCart entity.UserCart) DiscountResult {
	var discountPercent int16 = 0
	var discountResult float64 = 0
	var taxResult float64 = 0

	for i := range userCart.Items {
		item := &userCart.Items[i]
		quantity := item.Quantity
		if quantity > 3 {
			discountPercent = 8
			quantity -= 3
			discount := calculateDiscount(item, discountPercent)
			newTax := utils.CalculatePercent(item.Product.Price-discount, item.Product.Vat)
			discountResult += discount * float64(quantity-3)
			taxResult += newTax * float64(quantity)
		}
		taxResult += item.OrgTax * float64(quantity)
		discountPercent = 0
	}
	return DiscountResult{DiscountAmount: discountResult, TaxAmount: taxResult}
}

type MonthlyDiscount struct {
	DiscountThreshold float64
	CustomerTotal     float64
}

func (d MonthlyDiscount) CalculateCart(userCart entity.UserCart) DiscountResult {
	if d.CustomerTotal < d.DiscountThreshold {
		return DiscountResult{TaxAmount: userCart.CartSummary.TaxTotal}
	}

	var discountPercent int16 = 10
	var discountResult float64 = 0
	var taxResult float64 = 0

	for i := range userCart.Items {
		item := &userCart.Items[i]
		discount := calculateDiscount(item, discountPercent)
		newTax := utils.CalculatePercent(item.Product.Price-discount, item.Product.Vat)
		discountResult += discount * float64(item.Quantity)
		taxResult += newTax * float64(item.Quantity)
	}

	return DiscountResult{DiscountAmount: discountResult, TaxAmount: taxResult}
}

func calculateDiscount(item *entity.CartItem, discountPercent int16) float64 {
	return utils.CalculatePercent(item.Product.Price, discountPercent)
}

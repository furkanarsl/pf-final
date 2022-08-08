package services

import (
	"github.com/furkanarsl/pf-final/app/entity"
	"testing"
)

func TestFourthItemDiscount_CalculateCart(t *testing.T) {
	f := FourthItemDiscount{}
	cart := entity.UserCart{
		ID: 1,
		Items: []entity.CartItem{
			{
				ID: 1,
				Product: entity.Product{
					ID:    1,
					Name:  "p1",
					Price: 100,
					Vat:   18,
				},
				Price:      100,
				OrgTax:     18,
				Quantity:   4,
				TotalPrice: 118,
			},
		},
		CartSummary: entity.CartSummary{
			ProductTotal:   472,
			FinalPrice:     472,
			TaxTotal:       72,
			DiscountAmount: 0,
		},
	}
	result := f.CalculateCart(cart)
	if result.DiscountAmount != 8 {
		t.Errorf("result.DiscountAmount: %v want: %v ", result.DiscountAmount, 8)
	}
}

func TestFourthPurchaseMonthlyDiscount_CalculateCart(t *testing.T) {
	f := FourthPurchaseMonthlyDiscount{
		DiscountThreshold: 50,
		OrderCount:        3,
	}
	cart := entity.UserCart{
		ID: 1,
		Items: []entity.CartItem{
			{
				ID: 1,
				Product: entity.Product{
					ID:    1,
					Name:  "p1",
					Price: 100,
					Vat:   18,
				},
				Price:      100,
				OrgTax:     18,
				Quantity:   1,
				TotalPrice: 118,
			},
		},
		CartSummary: entity.CartSummary{
			ProductTotal:   118,
			FinalPrice:     118,
			TaxTotal:       18,
			DiscountAmount: 0,
		},
	}
	result := f.CalculateCart(cart)
	if result.DiscountAmount != 0 {
		t.Errorf("result.DiscountAmount: %v want: %v ", result.DiscountAmount, 0)
	}
	f.OrderCount = 4
	result = f.CalculateCart(cart)

	if result.DiscountAmount != 15 {
		t.Errorf("result.DiscountAmount: %v want: %v ", result.DiscountAmount, 15)
	}
}

func TestMonthlyDiscount_CalculateCart(t *testing.T) {
	f := MonthlyDiscount{
		DiscountThreshold: 150,
		CustomerTotal:     100,
	}
	cart := entity.UserCart{
		ID: 1,
		Items: []entity.CartItem{
			{
				ID: 1,
				Product: entity.Product{
					ID:    1,
					Name:  "p1",
					Price: 100,
					Vat:   18,
				},
				Price:      100,
				OrgTax:     18,
				Quantity:   1,
				TotalPrice: 118,
			},
		},
		CartSummary: entity.CartSummary{
			ProductTotal:   118,
			FinalPrice:     118,
			TaxTotal:       18,
			DiscountAmount: 0,
		},
	}
	result := f.CalculateCart(cart)
	if result.DiscountAmount != 0 {
		t.Errorf("result.DiscountAmount: %v want: %v ", result.DiscountAmount, 0)
	}
	f.CustomerTotal = 200
	result = f.CalculateCart(cart)

	if result.DiscountAmount != 10 {
		t.Errorf("result.DiscountAmount: %v want: %v ", result.DiscountAmount, 10)
	}
}

package services

import (
	"database/sql"
	"github.com/furkanarsl/pf-final/app/entity"
	"github.com/furkanarsl/pf-final/pkg/queries"
	"testing"
)

// Makes sure that tax and total is calculated correctly
func TestCartSvc_calculateCart(t *testing.T) {
	userCart := entity.UserCart{
		ID:          1,
		Items:       []entity.CartItem{},
		CartSummary: entity.CartSummary{},
	}
	itemsFromDB := []queries.ListCartItemsRow{
		{
			ID:       0,
			Quantity: 1,
			Name: sql.NullString{
				String: "Product1",
				Valid:  true,
			},
			Price: sql.NullFloat64{
				Float64: 100,
				Valid:   true,
			},
			Vat: sql.NullInt16{
				Int16: 18,
				Valid: true,
			},
			ProductID: sql.NullInt64{
				Int64: 1,
				Valid: true,
			},
		},
	}
	calculateCartItems(&itemsFromDB, &userCart)
	if userCart.CartSummary.DiscountAmount != 0 {
		t.Error()
	}
	if userCart.CartSummary.ProductTotal != 118 {
		t.Errorf("userCart.CartSummary.ProductTotal: %v, want: 118 ", userCart.CartSummary.ProductTotal)
	}

	if userCart.CartSummary.TaxTotal != 18 {
		t.Errorf("userCart.CartSummary.TaxTotal: %v, want: 18 ", userCart.CartSummary.TaxTotal)
	}

	if userCart.CartSummary.FinalPrice != 118 {
		t.Errorf("userCart.CartSummary.FinalPrice: %v, want: 18 ", userCart.CartSummary.FinalPrice)
	}
}

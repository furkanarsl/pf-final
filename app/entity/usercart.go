package entity

type UserCart struct {
	ID          int64       `json:"id"`
	Items       []CartItem  `json:"items"`
	CartSummary CartSummary `json:"summary"`
}

type CartItem struct {
	ID         int64   `json:"id"`
	Product    Product `json:"product"`
	Price      float64 `json:"sell_price"`
	OrgTax     float64 `json:"tax"`
	Quantity   int32   `json:"quantity"`
	TotalPrice float64 `json:"total"`
}

type CartSummary struct {
	ProductTotal float64 `json:"product_total"`
	FinalPrice   float64 `json:"final_price"`
	TaxTotal     float64 `json:"tax_total"`

	DiscountAmount float64 `json:"discount_amount"`
}

type Product struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Vat   int16   `json:"vat"`
}

type CartAddResult struct {
	ID      int64   `json:"id"`
	Product Product `json:"product"`
	CartID  int64   `json:"cart_id"`
}

package entity

type UserCart struct {
	ID         int64      `json:"id"`
	Items      []CartItem `json:"cart_items"`
	TotalPrice float64    `json:"total_price"`
	TotalTax   float64    `json:"tax_total"`
}

type CartItem struct {
	ID        int64   `json:"id"`
	Price     float64 `json:"price"`
	Quantity  int32   `json:"quantity"`
	OrgPrice  float64 `json:"original_price"`
	Total     float64 `json:"total_price"`
	Vat       int16   `json:"vat_percent"`
	TaxAmount float64 `json:"tax_amount"`
}

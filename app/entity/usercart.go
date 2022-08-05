package entity

type UserCart struct {
	ID             int64      `json:"id"`
	Items          []CartItem `json:"items"`
	TotalPriceDisc float64    `json:"disc_total_price"`
	TotalTaxDisc   float64    `json:"disc_total_tax"`
	TotalPrice     float64    `json:"total_price"`
	TotalTax       float64    `json:"total_tax"`
}

type CartItem struct {
	ID      int64   `json:"id"`
	Product Product `json:"product"`
	Price   float64 `json:"sell_price"`
	OrgTax  float64 `json:"tax"`

	DiscPrice float64 `json:"disc_sell_price"`
	DiscTax   float64 `json:"disc_tax"`
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

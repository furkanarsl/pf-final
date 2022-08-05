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
	ID       int64 `json:"id"`
	Quantity int32 `json:"quantity"`

	Price    float64 `json:"sell_price"`
	OrgPrice float64 `json:"org_price"`
	OrgTax   float64 `json:"tax"`

	DiscPrice    float64 `json:"disc_sell_price"`
	DiscOrgPrice float64 `json:"disc_org_price"`
	DiscTax      float64 `json:"disc_tax"`

	Total     float64 `json:"total_price"`
	DiscTotal float64 `json:"disc_total_price"`
	Vat       int16   `json:"tax_percent"`
}

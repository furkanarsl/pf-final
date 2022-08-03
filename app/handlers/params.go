package handlers

type AddToCartParams struct {
	ProductID int64 `json:"product_id"`
	Quantity  int32 `json:"quantity"`
}

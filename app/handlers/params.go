package handlers

type AddToCartParams struct {
	ProductID int64 `json:"product_id"`
	Quantity  int32 `json:"quantity"`
}

type RemoveFromCartParams struct {
	ProductID int64 `json:"product_id"`
	CartID    int64 `json:"cart_id"`
}

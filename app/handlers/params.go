package handlers

type AddToCartParams struct {
	ProductID int64 `json:"product_id"`
}

type RemoveFromCartParams struct {
	ProductID int64 `json:"product_id"`
	CartID    int64 `json:"cart_id"`
}

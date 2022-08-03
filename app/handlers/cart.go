package handlers

import "github.com/gin-gonic/gin"

type cartHandler struct {
}

func NewCartHandler(r *gin.RouterGroup) {

	handler := cartHandler{}

	r.GET("/cart", handler.ListCart)
	r.POST("/cart", handler.AddToCart)
	r.DELETE("/cart/item/:id", handler.DeleteFromCart)
}

func (h *cartHandler) ListCart(c *gin.Context) {
	items := []string{"123"}
	c.JSON(200, gin.H{"items": items, "total_price": 123, "total_selling_price": 150, "total_vat": 15})
}

func (h *cartHandler) AddToCart(c *gin.Context) {
	c.JSON(200, gin.H{"item": "Add an item to cart return the added item"})
}

func (h *cartHandler) DeleteFromCart(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"item": "Delete item from cart", "id": id, "status": true})
}

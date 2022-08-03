package handlers

import "github.com/gin-gonic/gin"

type orderHandler struct {
}

func NewOrderHandler(r *gin.RouterGroup) {

	handler := orderHandler{}

	r.POST("/order/complete", handler.CompleteOrder)
}

func (h *orderHandler) CompleteOrder(c *gin.Context) {
	c.JSON(200, gin.H{"status": "Order completed"})
}

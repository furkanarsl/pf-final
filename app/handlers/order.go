package handlers

import (
	"net/http"
	"strconv"

	"github.com/furkanarsl/pf-final/app/services"
	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	orderService services.OrderService
}

func NewOrderHandler(r *gin.RouterGroup, orderService services.OrderService) {

	handler := orderHandler{orderService: orderService}

	r.POST("/order/complete", handler.CompleteOrder)
}

func (h *orderHandler) CompleteOrder(c *gin.Context) {
	qID, _ := c.GetQuery("user_id")
	userID, err := strconv.ParseInt(qID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid user id"})
		return
	}
	order, err := h.orderService.CompleteOrder(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	c.JSON(200, order)
}

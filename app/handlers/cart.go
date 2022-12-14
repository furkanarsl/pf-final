package handlers

import (
	"net/http"
	"strconv"

	"github.com/furkanarsl/pf-final/app/services"
	"github.com/gin-gonic/gin"
)

type cartHandler struct {
	cartService services.CartService
}

func NewCartHandler(r *gin.RouterGroup, cartService services.CartService) {

	handler := cartHandler{cartService: cartService}

	r.GET("/cart", handler.ListCart)
	r.POST("/cart", handler.AddToCart)
	r.DELETE("/cart/item/:id", handler.DeleteFromCart)
}

func (h *cartHandler) ListCart(c *gin.Context) {

	qID, _ := c.GetQuery("user_id")
	userID, err := strconv.ParseInt(qID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid user id"})
		return
	}

	cart, err := h.cartService.ListCart(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "User doesn't exists"})
		return
	}

	c.JSON(200, cart)
}

func (h *cartHandler) AddToCart(c *gin.Context) {

	qID, _ := c.GetQuery("user_id")
	userID, err := strconv.ParseInt(qID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid user id"})
		return
	}

	addParams := AddToCartParams{}
	err = c.BindJSON(&addParams)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid body"})
		return
	}

	if addParams.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Provide a valid quantity number"})
		return
	}

	res, err := h.cartService.AddToCart(userID, addParams.ProductID, addParams.Quantity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Failed to add to cart"})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *cartHandler) DeleteFromCart(c *gin.Context) {
	id := c.Param("id")

	qID, _ := c.GetQuery("user_id")
	userID, err := strconv.ParseInt(qID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid user id"})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid body"})
	}

	itemID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid itemID"})
		return
	}

	err = h.cartService.RemoveFromCart(userID, itemID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Failed to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id, "status": "Success"})
}

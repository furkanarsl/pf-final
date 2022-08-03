package handlers

import (
	"github.com/furkanarsl/pf-final/app/services"
	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService services.ProductService
}

func NewProductHandler(r *gin.RouterGroup, productService services.ProductService) {

	handler := productHandler{productService: productService}

	r.GET("/products", handler.ListProducts)
	r.GET("/products/:id", handler.GetProduct)
}

func (h *productHandler) ListProducts(c *gin.Context) {
	products := h.productService.ListProducts()
	c.JSON(200, gin.H{"products": products})
}

func (h *productHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"products": "Details of a product", id: id})
}

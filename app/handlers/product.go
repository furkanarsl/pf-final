package handlers

import (
	"net/http"
	"strconv"

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
	products, err := h.productService.ListProducts()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "No products found"})
		return
	}
	c.JSON(200, products)
}

func (h *productHandler) GetProduct(c *gin.Context) {
	// TODO: Later change how we handle the errors
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Please enter a valid id"})
		return
	}

	product, err := h.productService.GetProduct(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "Product not found"})
		return
	}
	c.JSON(200, product)
}

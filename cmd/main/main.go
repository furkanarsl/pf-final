package main

import (
	"github.com/furkanarsl/pf-final/app/handlers"
	"github.com/furkanarsl/pf-final/app/repository"
	"github.com/furkanarsl/pf-final/app/services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api := r.Group("/api/v1")
	//Create repositories
	productRepo := repository.NewProductRepo()

	//Create services
	productSvc := services.NewProductService(productRepo)

	//Register handlers
	handlers.NewProductHandler(api, productSvc)
	handlers.NewCartHandler(api)
	handlers.NewOrderHandler(api)
	r.Run()
}

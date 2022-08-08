package main

import (
	"context"
	"os"
	"strconv"

	"github.com/furkanarsl/pf-final/app/handlers"
	"github.com/furkanarsl/pf-final/app/repository"
	"github.com/furkanarsl/pf-final/app/services"
	"github.com/furkanarsl/pf-final/database"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := gin.Default()

	api := r.Group("/api/v1")
	//Connect to DB
	// Maybe move this to another file
	db, err := database.Open(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	queries := database.NewQueries(db)
	defer db.Close(context.Background())
	//Create repositories
	productRepo := repository.NewProductRepo(queries)
	cartRepo := repository.NewCartRepo(queries)
	orderRepo := repository.NewOrderRepo(queries)
	//Create services
	dcThreshold, _ := strconv.ParseFloat(os.Getenv("DISCOUNT_THRESHOLD_AMOUNT"), 64)
	discountSvc := services.NewDiscountService(dcThreshold)
	productSvc := services.NewProductService(productRepo)
	cartSvc := services.NewCartService(cartRepo, productRepo, orderRepo, discountSvc)
	orderSvc := services.NewOrderService(orderRepo, cartSvc)
	//Register handlers
	handlers.NewProductHandler(api, productSvc)
	handlers.NewCartHandler(api, cartSvc)
	handlers.NewOrderHandler(api, orderSvc)
	r.Run()
}

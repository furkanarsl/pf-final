package main

import (
	"context"
	"os"

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
	db, err := database.Open(os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}

	queries := database.NewQueries(db)
	defer db.Close(context.Background())
	//Create repositories
	productRepo := repository.NewProductRepo(queries)
	cartRepo := repository.NewCartRepo(queries)

	//Create services
	productSvc := services.NewProductService(productRepo)
	cartSvc := services.NewCartService(cartRepo, productRepo)

	//Register handlers
	handlers.NewProductHandler(api, productSvc)
	handlers.NewCartHandler(api, cartSvc)
	handlers.NewOrderHandler(api)
	r.Run()
}

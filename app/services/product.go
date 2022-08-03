package services

import (
	"github.com/furkanarsl/pf-final/app/repository"
	"github.com/furkanarsl/pf-final/pkg/queries"
)

type productSvc struct {
	productRepo repository.ProductRepo
}

type ProductService interface {
	ListProducts() ([]queries.Product, error) // TODO: Maybe change to a dto so we can set field names
	GetProduct(id int64) (queries.Product, error)
}

func NewProductService(productRepo repository.ProductRepo) *productSvc {
	return &productSvc{productRepo: productRepo}
}

func (s *productSvc) ListProducts() ([]queries.Product, error) {
	products, _ := s.productRepo.FindAll()
	return products, nil
}

func (s *productSvc) GetProduct(id int64) (queries.Product, error) {
	product, err := s.productRepo.FindOne(id)
	if err != nil {
		return product, err
	}
	return product, nil
}

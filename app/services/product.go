package services

import (
	"github.com/furkanarsl/pf-final/app/entity"
	"github.com/furkanarsl/pf-final/app/repository"
)

type productSvc struct {
	productRepo repository.ProductRepo
}

type ProductService interface {
	ListProducts() ([]entity.Product, error)
	GetProduct(id int64) (entity.Product, error)
}

func NewProductService(productRepo repository.ProductRepo) *productSvc {
	return &productSvc{productRepo: productRepo}
}

func (s *productSvc) ListProducts() ([]entity.Product, error) {
	products, _ := s.productRepo.FindAll()
	results := []entity.Product{}
	for _, product := range products {
		results = append(results, entity.Product(product))
	}
	return results, nil
}

func (s *productSvc) GetProduct(id int64) (entity.Product, error) {
	product, err := s.productRepo.FindOne(id)
	result := entity.Product(product)
	if err != nil {
		return result, err
	}
	return result, nil
}

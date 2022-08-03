package services

import "github.com/furkanarsl/pf-final/app/repository"

type productSvc struct {
	productRepo repository.ProductRepo
}

type ProductService interface {
	ListProducts() []string // TODO: Change to product model later
}

func NewProductService(productRepo repository.ProductRepo) *productSvc {
	return &productSvc{productRepo: productRepo}
}

func (s *productSvc) ListProducts() []string {
	products := s.productRepo.FindAll()
	return products
}

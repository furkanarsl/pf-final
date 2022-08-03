package repository

import (
	"context"

	"github.com/furkanarsl/pf-final/database"
	"github.com/furkanarsl/pf-final/pkg/queries"
)

type productRepo struct {
	database.DbQueries
}

type ProductRepo interface {
	FindAll() ([]queries.Product, error)
	FindOne(id int64) (queries.Product, error)
}

func NewProductRepo(queries database.DbQueries) *productRepo {
	return &productRepo{queries}
}

func (r *productRepo) FindAll() ([]queries.Product, error) {
	products, err := r.ListProducts(context.Background())
	if err != nil {
		return products, err
	}
	return products, nil
}

func (r *productRepo) FindOne(id int64) (queries.Product, error) {
	product, err := r.GetProduct(context.Background(), id)
	if err != nil {
		return product, err
	}
	return product, nil
}

package repository

type productRepo struct {
}

type ProductRepo interface {
	FindAll() []string
}

func NewProductRepo() *productRepo {
	return &productRepo{}
}

func (r *productRepo) FindAll() []string {
	products := []string{"product1", "pro2", "p3"}
	return products
}

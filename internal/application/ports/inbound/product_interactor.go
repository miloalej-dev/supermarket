package inbound

import "github.com/miloalej-dev/supermarket/internal/domain"

type ProductInteractor interface {
	CreateProduct(domain.Product) (domain.Product, error)
	GetProduct(id int) (domain.Product, error)
	ListProducts() ([]domain.Product, error)
	UpdateProduct(id int, updates map[string]interface{}) (*domain.Product, error)
	DeleteProduct(id int) error
	CreateOrUpdate(id int, newProduct domain.Product) (domain.Product, error)
}

package domain

import "github.com/tomazcx/products-api/internal/entity"

type IProductRepository interface {
	FindById(id string) (*entity.Product, error)
	FindAll(limit, size int, sort string) ([]entity.Product, error)
	Create(product *entity.Product) error
	Update(product *entity.Product) error
	Delete(id string) error
}

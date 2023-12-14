package domain

import "github.com/tomazcx/products-api/internal/entity"

type IProductRepository interface {
	Exists(id string) bool
	FindById(id string) (*entity.Product, error)
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	Create(product *entity.Product) error
	Update(product *entity.Product) error
	Delete(id string) error
}

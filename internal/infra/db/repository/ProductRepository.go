package repository

import (
	"github.com/tomazcx/products-api/internal/entity"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (p *ProductRepository) FindById(id string) (*entity.Product, error) {
	var product entity.Product
	err := p.DB.First(&product).Where("id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductRepository) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *ProductRepository) Update(product *entity.Product) error {
	return p.DB.Save(product).Error
}

func (p *ProductRepository) Delete(id string) error {
	return p.DB.Where("id = ?", id).Delete(&entity.Product{}).Error
}

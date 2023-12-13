package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomazcx/products-api/internal/entity"
	"github.com/tomazcx/products-api/internal/infra/db/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestFindProductById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Test product", 9.99)
	assert.Nil(t, err)

	err = db.Create(product).Error
	assert.Nil(t, err)

	productRepository := repository.NewProductRepository(db)
	productFound, err := productRepository.FindById(product.ID.String())

	assert.Nil(t, err)
	assert.Equal(t, productFound.ID, product.ID)
	assert.Equal(t, productFound.Name, product.Name)
	assert.Equal(t, productFound.Price, product.Price)
	assert.NotEmpty(t, productFound.CreatedAt)
}

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Test Product", 9.99)
	assert.Nil(t, err)

	productRepository := repository.NewProductRepository(db)
	err = productRepository.Create(product)
	assert.Nil(t, err)

	var productFound entity.Product
	err = db.First(&productFound).Where("id = ?", product.ID).Error

	assert.Nil(t, err)
	assert.Equal(t, productFound.ID, product.ID)
	assert.Equal(t, productFound.Name, product.Name)
	assert.Equal(t, productFound.Price, product.Price)
	assert.NotEmpty(t, productFound.CreatedAt)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Test Product", 9.99)
	assert.Nil(t, err)

	err = db.Create(&product).Error
	assert.Nil(t, err)

	product.Name = "Test Product Updated"
	product.Price = 10.99

	productRepository := repository.NewProductRepository(db)
	err = productRepository.Update(product)

	var productFound entity.Product

	err = db.First(&productFound).Where("id = ?", product.ID).Error

	assert.Nil(t, err)
	assert.Equal(t, productFound.Name, product.Name)
	assert.Equal(t, productFound.Price, product.Price)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Test Product", 9.99)
	assert.Nil(t, err)

	err = db.Create(product).Error
	assert.Nil(t, err)

	productRepository := repository.NewProductRepository(db)
	err = productRepository.Delete(product.ID.String())

	var productFound entity.Product

	err = db.First(&productFound).Where("id = ?", product.ID).Error

	assert.NotNil(t, err)
}

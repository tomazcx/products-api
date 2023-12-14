package test

import (
	"fmt"
	"log"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomazcx/products-api/internal/entity"
	"github.com/tomazcx/products-api/internal/infra/db/repository"
	pkgEntity "github.com/tomazcx/products-api/pkg/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func setupTestProductRepository() (func(db *gorm.DB), *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entity.Product{})

	return func(db *gorm.DB) {
		sqlDB, err := db.DB()

		if err != nil {
			log.Fatalf("Failed to get the database coonection")
		}

		sqlDB.Close()
	}, db
}

func TestProductExists(t *testing.T) {
	teardown, db := setupTestProductRepository()
	defer teardown(db)

	product, err := entity.NewProduct("Test Product", 9.99)
	assert.Nil(t, err)

	err = db.Create(product).Error
	assert.Nil(t, err)

	productRepository := repository.NewProductRepository(db)
	productExists := productRepository.Exists(product.ID.String())

	assert.True(t, productExists)
}

func TestProductDontExist(t *testing.T) {
	teardown, db := setupTestProductRepository()
	defer teardown(db)

	id := pkgEntity.NewId()
	productRepository := repository.NewProductRepository(db)
	productExists := productRepository.Exists(id.String())

	assert.False(t, productExists)
}

func TestFindAll(t *testing.T) {
	teardown, db := setupTestProductRepository()
	defer teardown(db)

	for i := 1; i <= 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		err = db.Create(&product).Error
		assert.NoError(t, err)
	}

	productRepository := repository.NewProductRepository(db)

	products, err := productRepository.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productRepository.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productRepository.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 4)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 24", products[3].Name)
}

func TestFindProductById(t *testing.T) {
	teardown, db := setupTestProductRepository()
	defer teardown(db)

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
	teardown, db := setupTestProductRepository()
	defer teardown(db)

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
	teardown, db := setupTestProductRepository()
	defer teardown(db)

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
	teardown, db := setupTestProductRepository()
	defer teardown(db)

	product, err := entity.NewProduct("Test Product", 9.99)
	assert.Nil(t, err)

	err = db.Create(product).Error
	assert.Nil(t, err)

	productRepository := repository.NewProductRepository(db)
	err = productRepository.Delete(product.ID.String())

	var productFound entity.Product
	err = db.First(&productFound).Where("id = ?", product.ID).Error

	assert.Error(t, err)
}

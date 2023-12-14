package test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	usecase "github.com/tomazcx/products-api/internal/data/use-case/product"
	"github.com/tomazcx/products-api/internal/entity"
	"github.com/tomazcx/products-api/internal/infra/db/repository/mocks"
)

func TestShowManyProductsUseCase(t *testing.T) {
	repo := &mocks.ProductRepositoryMock{}

	var products []entity.Product

	for i := 0; i < 10; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product n. %d", i), 9.99)
		assert.Nil(t, err)
		products = append(products, *product)
	}

	repo.On("FindAll", 1, 10, "asc").Return(products, nil)

	showManyProductsUseCase := usecase.ShowManyProductsUseCase{Repository: repo}
	page := 1
	limit := 10
	sort := "asc"
	result, err := showManyProductsUseCase.Execute(page, limit, sort)

	assert.Equal(t, products, result)
	assert.Nil(t, err)
}

func TestShowManyProductsUseCase_When_Fails(t *testing.T) {
	repo := &mocks.ProductRepositoryMock{}

	repo.On("FindAll", 1, 10, "asc").Return([]entity.Product{}, errors.New("DB ERROR"))

	showManyProductsUseCase := usecase.ShowManyProductsUseCase{Repository: repo}
	page := 1
	limit := 10
	sort := "asc"
	result, err := showManyProductsUseCase.Execute(page, limit, sort)

	assert.NotNil(t, err)
	assert.Len(t, result, 0)
}

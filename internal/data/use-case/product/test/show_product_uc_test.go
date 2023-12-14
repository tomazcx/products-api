package test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	usecase "github.com/tomazcx/products-api/internal/data/use-case/product"
	"github.com/tomazcx/products-api/internal/entity"
	"github.com/tomazcx/products-api/internal/infra/db/repository/mocks"
	entityPkg "github.com/tomazcx/products-api/pkg/entity"
)

func TestShowProductUseCase(t *testing.T) {
	productId := entityPkg.NewId()
	expectedProduct := &entity.Product{ID: productId, Name: "Test Product", Price: 9.99, CreatedAt: time.Now()}

	repo := &mocks.ProductRepositoryMock{}
	repo.On("Exists", productId.String()).Return(true)
	repo.On("FindById", productId.String()).Return(expectedProduct, nil)

	showProductUseCase := usecase.ShowProductUseCase{Repository: repo}

	product, err := showProductUseCase.Execute(productId.String())

	assert.Nil(t, err)
	assert.Equal(t, product, expectedProduct)
	repo.AssertCalled(t, "Exists", productId.String())
	repo.AssertCalled(t, "FindById", productId.String())
}

func TestShowProductUseCase_When_Product_Dont_Exist(t *testing.T) {
	productId := entityPkg.NewId()

	repo := &mocks.ProductRepositoryMock{}
	repo.On("Exists", productId.String()).Return(false)

	showProductUseCase := usecase.ShowProductUseCase{Repository: repo}

	_, err := showProductUseCase.Execute(productId.String())

	assert.NotNil(t, err)
	repo.AssertCalled(t, "Exists", productId.String())
	repo.AssertNotCalled(t, "FindById", productId.String())
}

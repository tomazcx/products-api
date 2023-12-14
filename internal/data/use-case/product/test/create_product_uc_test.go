package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	usecase "github.com/tomazcx/products-api/internal/data/use-case/product"
	"github.com/tomazcx/products-api/internal/dto"
	"github.com/tomazcx/products-api/internal/entity"
	"github.com/tomazcx/products-api/internal/infra/db/repository/mocks"
	entityPkg "github.com/tomazcx/products-api/pkg/entity"
)

func TestCreateProductUseCase(t *testing.T) {
	repo := &mocks.ProductRepositoryMock{}

	expectedProduct := &entity.Product{
		ID:    entityPkg.NewId(),
		Name:  "Test Product",
		Price: 9.99,
	}
	repo.On("Create", mock.Anything).Return(nil)

	createProductUseCase := usecase.CreateProductUseCase{Repository: repo}

	createProductDto := dto.ProductDTO{
		Name:  "Test Product",
		Price: 9.99,
	}
	product, err := createProductUseCase.Execute(createProductDto)

	assert.Nil(t, err)
	assert.Equal(t, product.Name, expectedProduct.Name)
	assert.Equal(t, product.Price, expectedProduct.Price)
	repo.AssertCalled(t, "Create", mock.Anything)
}

func TestCreateProductUseCase_Invalid_Name_DTO(t *testing.T) {
	repo := &mocks.ProductRepositoryMock{}
	repo.On("Create", mock.Anything).Return(nil)

	createProductUseCase := usecase.CreateProductUseCase{Repository: repo}

	createProductDto := dto.ProductDTO{
		Name:  "",
		Price: 9.99,
	}
	_, err := createProductUseCase.Execute(createProductDto)

	assert.NotNil(t, err)
	repo.AssertNotCalled(t, "Create", mock.Anything)
}

func TestCreateProductUseCase_Invalid_Price_DTO(t *testing.T) {
	repo := &mocks.ProductRepositoryMock{}
	repo.On("Create", mock.Anything).Return(nil)

	createProductUseCase := usecase.CreateProductUseCase{Repository: repo}

	createProductDto := dto.ProductDTO{
		Name:  "Test Product",
		Price: 0.0,
	}
	_, err := createProductUseCase.Execute(createProductDto)

	assert.NotNil(t, err)
	repo.AssertNotCalled(t, "Create", mock.Anything)
}

package test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	usecase "github.com/tomazcx/products-api/internal/data/use-case/product"
	"github.com/tomazcx/products-api/internal/dto"
	"github.com/tomazcx/products-api/internal/entity"
	"github.com/tomazcx/products-api/internal/infra/db/repository/mocks"
	entityPkg "github.com/tomazcx/products-api/pkg/entity"
)

func TestUpdateProductUseCase(t *testing.T) {
	repo := &mocks.ProductRepositoryMock{}
	id := entityPkg.NewId()
	productToUpadte := &entity.Product{
		ID:    id,
		Name:  "Product Updated",
		Price: 10.99,
	}
	repo.On("FindById", id.String()).Return(productToUpadte, nil)
	repo.On("Update", productToUpadte).Return(nil)

	updateProductUseCase := usecase.UpdateProductUseCase{Repository: repo}

	updateProductDto := dto.ProductDTO{
		Name:  "Product Updated",
		Price: 10.99,
	}
	err := updateProductUseCase.Execute(updateProductDto, id.String())

	assert.Nil(t, err)
	repo.AssertCalled(t, "FindById", id.String())
	repo.AssertCalled(t, "Update", productToUpadte)
}

func TestUpdateProductUseCase_When_Product_Dont_Exists(t *testing.T) {
	repo := &mocks.ProductRepositoryMock{}
	id := entityPkg.NewId()
	repo.On("FindById", id.String()).Return((*entity.Product)(nil), errors.New("DB ERROR: Not found"))

	updateProductUseCase := usecase.UpdateProductUseCase{Repository: repo}

	updateProductDto := dto.ProductDTO{
		Name:  "Product Updated",
		Price: 10.99,
	}
	err := updateProductUseCase.Execute(updateProductDto, id.String())

	assert.NotNil(t, err)
	repo.AssertCalled(t, "FindById", id.String())
	repo.AssertNotCalled(t, "Update", mock.Anything)
}

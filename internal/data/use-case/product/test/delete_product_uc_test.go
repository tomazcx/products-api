package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	usecase "github.com/tomazcx/products-api/internal/data/use-case/product"
	"github.com/tomazcx/products-api/internal/infra/db/repository/mocks"
	entityPkg "github.com/tomazcx/products-api/pkg/entity"
)

func TestDeleteProductUseCase(t *testing.T) {
	repo := &mocks.ProductRepositoryMock{}
	id := entityPkg.NewId()

	repo.On("Exists", id.String()).Return(true)
	repo.On("Delete", id.String()).Return(nil)

	deleteProductUseCase := usecase.DeleteProductUseCase{Repository: repo}
	err := deleteProductUseCase.Execute(id.String())

	assert.Nil(t, err)
	repo.AssertCalled(t, "Exists", id.String())
	repo.AssertCalled(t, "Delete", id.String())
}

func TestDeleteWHenProductUseCase_When_Product_Dont_Exist(t *testing.T) {
	repo := &mocks.ProductRepositoryMock{}
	id := entityPkg.NewId()

	repo.On("Exists", id.String()).Return(false)

	deleteProductUseCase := usecase.DeleteProductUseCase{Repository: repo}
	err := deleteProductUseCase.Execute(id.String())

	assert.NotNil(t, err)
	repo.AssertCalled(t, "Exists", id.String())
	repo.AssertNotCalled(t, "Delete", mock.Anything)
}

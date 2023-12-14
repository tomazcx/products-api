package usecase

import (
	"net/http"

	"github.com/tomazcx/products-api/internal/entity"
	"github.com/tomazcx/products-api/internal/infra/db/domain"
	"github.com/tomazcx/products-api/pkg/httperr"
)

type ShowProductUseCase struct {
	Repository domain.IProductRepository
}

func (uc *ShowProductUseCase) Execute(id string) (*entity.Product, error) {
	productExists := uc.Repository.Exists(id)

	if !productExists {
		return nil, httperr.NewHttpError("Product not found", http.StatusNotFound)
	}

	product, err := uc.Repository.FindById(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

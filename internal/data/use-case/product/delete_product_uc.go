package usecase

import (
	"net/http"

	"github.com/tomazcx/products-api/internal/infra/db/domain"
	"github.com/tomazcx/products-api/pkg/httperr"
)

type DeleteProductUseCase struct {
	Repository domain.IProductRepository
}

func (uc *DeleteProductUseCase) Execute(id string) error {
	productExists := uc.Repository.Exists(id)

	if !productExists {
		return httperr.NewHttpError("Product not found", http.StatusNotFound)
	}

	err := uc.Repository.Delete(id)

	return err
}

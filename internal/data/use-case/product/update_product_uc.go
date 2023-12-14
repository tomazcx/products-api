package usecase

import (
	"net/http"

	"github.com/tomazcx/products-api/internal/dto"
	"github.com/tomazcx/products-api/internal/infra/db/domain"
	"github.com/tomazcx/products-api/pkg/httperr"
)

type UpdateProductUseCase struct {
	Repository domain.IProductRepository
}

func (uc *UpdateProductUseCase) Execute(dto dto.ProductDTO, id string) error {
	product, err := uc.Repository.FindById(id)

	if err != nil {
		return httperr.NewHttpError("Product not found", http.StatusNotFound)
	}

	product.Name = dto.Name
	product.Price = dto.Price

	err = uc.Repository.Update(product)

	return err
}

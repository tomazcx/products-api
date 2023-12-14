package usecase

import (
	"fmt"
	"net/http"

	"github.com/tomazcx/products-api/internal/dto"
	"github.com/tomazcx/products-api/internal/entity"
	"github.com/tomazcx/products-api/internal/infra/db/domain"
	"github.com/tomazcx/products-api/pkg/httperr"
)

type CreateProductUseCase struct {
	Repository domain.IProductRepository
}

func (uc *CreateProductUseCase) Execute(dto dto.ProductDTO) (*entity.Product, error) {
	product, err := entity.NewProduct(dto.Name, dto.Price)

	if err != nil {
		return nil, httperr.NewHttpError(fmt.Sprintf("Error creating the product: %v", err), http.StatusUnprocessableEntity)
	}

	err = uc.Repository.Create(product)

	if err != nil {
		return nil, err
	}

	return product, nil
}

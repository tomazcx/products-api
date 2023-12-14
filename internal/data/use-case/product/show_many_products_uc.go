package usecase

import (
	"github.com/tomazcx/products-api/internal/entity"
	"github.com/tomazcx/products-api/internal/infra/db/domain"
)

type ShowManyProductsUseCase struct {
	Repository domain.IProductRepository
}

func (uc ShowManyProductsUseCase) Execute(page, limit int, sort string) ([]entity.Product, error) {
	if page == 0 {
		page = 1
	}

	products, err := uc.Repository.FindAll(page, limit, sort)

	if err != nil {
		return []entity.Product{}, err
	}

	return products, nil
}

package factory

import (
	"github.com/tomazcx/products-api/configs"
	usecase "github.com/tomazcx/products-api/internal/data/use-case/product"
	"github.com/tomazcx/products-api/internal/infra/db/repository"
)

type ProductFactory struct{}

func (f ProductFactory) ShowManyProductsUseCase() *usecase.ShowManyProductsUseCase {
	db := configs.GetDBInstance()
	repo := repository.NewProductRepository(db)
	return &usecase.ShowManyProductsUseCase{Repository: repo}
}

func (f ProductFactory) ShowProductUseCase() *usecase.ShowProductUseCase {
	db := configs.GetDBInstance()
	repo := repository.NewProductRepository(db)
	return &usecase.ShowProductUseCase{Repository: repo}
}

func (f ProductFactory) CreateProductUseCase() *usecase.CreateProductUseCase {
	db := configs.GetDBInstance()
	repo := repository.NewProductRepository(db)
	return &usecase.CreateProductUseCase{Repository: repo}
}

func (f ProductFactory) DeleteProductUseCase() *usecase.DeleteProductUseCase {
	db := configs.GetDBInstance()
	repo := repository.NewProductRepository(db)
	return &usecase.DeleteProductUseCase{Repository: repo}
}

func (f ProductFactory) UpdateProductUseCase() *usecase.UpdateProductUseCase {
	db := configs.GetDBInstance()
	repo := repository.NewProductRepository(db)
	return &usecase.UpdateProductUseCase{Repository: repo}
}

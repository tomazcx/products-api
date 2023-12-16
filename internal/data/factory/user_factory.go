package factory

import (
	"github.com/tomazcx/products-api/configs"
	usecase "github.com/tomazcx/products-api/internal/data/use-case/user"
	"github.com/tomazcx/products-api/internal/infra/db/repository"
)

type UserFactory struct{}

func (f UserFactory) CreateUserUseCase() *usecase.CreateUserUseCase {
	db := configs.GetDBInstance()
	repo := repository.NewUserRepository(db)
	return &usecase.CreateUserUseCase{Repository: repo}
}

func (f UserFactory) AuthUserUseCase() *usecase.AuthUserUseCase {
	db := configs.GetDBInstance()
	repo := repository.NewUserRepository(db)
	return &usecase.AuthUserUseCase{Repository: repo}
}

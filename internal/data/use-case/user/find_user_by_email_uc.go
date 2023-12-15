package usecase

import (
	"net/http"

	"github.com/tomazcx/products-api/internal/entity"
	"github.com/tomazcx/products-api/internal/infra/db/domain"
	"github.com/tomazcx/products-api/pkg/httperr"
)

type FindUserByEmailUseCase struct {
	Repository domain.IUserRepository
}

func (uc *FindUserByEmailUseCase) Execute(email string) (*entity.User, error) {
	user, err := uc.Repository.FindByEmail(email)

	if err != nil {
		return nil, httperr.NewHttpError("User not found", http.StatusNotFound)
	}

	return user, nil
}

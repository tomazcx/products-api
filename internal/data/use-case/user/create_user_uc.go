package usecase

import (
	"fmt"
	"net/http"

	"github.com/tomazcx/products-api/internal/dto"
	"github.com/tomazcx/products-api/internal/entity"
	"github.com/tomazcx/products-api/internal/infra/db/domain"
	"github.com/tomazcx/products-api/pkg/httperr"
)

type CreateUserUseCase struct {
	Repository domain.IUserRepository
}

func (uc *CreateUserUseCase) Execute(dto dto.UserDTO) (*entity.User, error) {
	user, err := entity.NewUser(dto.Name, dto.Email, dto.Password)

	if err != nil {
		return nil, httperr.NewHttpError(fmt.Sprintf("Error creating the user: %v", err), http.StatusUnprocessableEntity)
	}

	err = uc.Repository.Create(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

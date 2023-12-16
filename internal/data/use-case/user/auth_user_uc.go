package usecase

import (
	"github.com/tomazcx/products-api/internal/dto"
	"github.com/tomazcx/products-api/internal/infra/db/domain"
)

type AuthUserUseCase struct {
	Repository domain.IUserRepository
}

func (uc *AuthUserUseCase) Execute(credentials dto.AuthDTO) bool {
	user, err := uc.Repository.FindByEmail(credentials.Email)

	if err != nil {
		return false
	}

	return user.ValidatePassword(credentials.Password)
}

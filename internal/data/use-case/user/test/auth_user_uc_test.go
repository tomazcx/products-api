package test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	usecase "github.com/tomazcx/products-api/internal/data/use-case/user"
	"github.com/tomazcx/products-api/internal/dto"
	"github.com/tomazcx/products-api/internal/entity"
	"github.com/tomazcx/products-api/internal/infra/db/repository/mocks"
)

func TestAuthUserUseCase(t *testing.T) {
	repo := &mocks.UserRepositoryMock{}

	user, _ := entity.NewUser("John Doe", "john@email.com", "12345")
	repo.On("FindByEmail", mock.Anything).Return(user, nil)

	authUserDTO := dto.AuthDTO{
		Email:    "john@email.com",
		Password: "12345",
	}

	authUserUseCase := usecase.AuthUserUseCase{Repository: repo}
	isValid := authUserUseCase.Execute(authUserDTO)

	assert.True(t, isValid)
	repo.AssertCalled(t, "FindByEmail", authUserDTO.Email)
}

func TestAuthUserUseCase_Invalid_Email(t *testing.T) {
	repo := &mocks.UserRepositoryMock{}

	repo.On("FindByEmail", mock.Anything).Return((*entity.User)(nil), errors.New("DB ERROR"))

	authUserDTO := dto.AuthDTO{
		Email:    "john@email.com",
		Password: "12345",
	}

	authUserUseCase := usecase.AuthUserUseCase{Repository: repo}
	isValid := authUserUseCase.Execute(authUserDTO)

	assert.False(t, isValid)
	repo.AssertCalled(t, "FindByEmail", authUserDTO.Email)
}

func TestAuthUserUseCase_Invalid_Password(t *testing.T) {
	repo := &mocks.UserRepositoryMock{}

	user, _ := entity.NewUser("John Doe", "john@email.com", "123456")
	repo.On("FindByEmail", mock.Anything).Return(user, nil)

	authUserDTO := dto.AuthDTO{
		Email:    "john@email.com",
		Password: "12345",
	}

	authUserUseCase := usecase.AuthUserUseCase{Repository: repo}
	isValid := authUserUseCase.Execute(authUserDTO)

	assert.False(t, isValid)
	repo.AssertCalled(t, "FindByEmail", authUserDTO.Email)
}

package test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	usecase "github.com/tomazcx/products-api/internal/data/use-case/user"
	"github.com/tomazcx/products-api/internal/dto"
	"github.com/tomazcx/products-api/internal/infra/db/repository/mocks"
)

func TestCreateUserUseCase(t *testing.T) {
	repo := &mocks.UserRepositoryMock{}

	repo.On("Create", mock.Anything).Return(nil)

	createUserUseCase := usecase.CreateUserUseCase{Repository: repo}

	createUserDto := dto.UserDTO{
		Name:     "John Doe",
		Email:    "john@email.com",
		Password: "12345",
	}
	user, err := createUserUseCase.Execute(createUserDto)

	assert.Nil(t, err)
	assert.Equal(t, user.Name, createUserDto.Name)
	assert.Equal(t, user.Email, createUserDto.Email)
	assert.NotEmpty(t, user.Password)
	assert.NotEmpty(t, user.ID)
	repo.AssertCalled(t, "Create", mock.Anything)
}

func TestCreateUserUseCase_Fails(t *testing.T) {
	repo := &mocks.UserRepositoryMock{}

	repo.On("Create", mock.Anything).Return(errors.New("DB ERROR"))

	createUserUseCase := usecase.CreateUserUseCase{Repository: repo}

	createUserDto := dto.UserDTO{
		Name:     "John Doe",
		Email:    "john@email.com",
		Password: "12345",
	}
	user, err := createUserUseCase.Execute(createUserDto)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	repo.AssertCalled(t, "Create", mock.Anything)
}

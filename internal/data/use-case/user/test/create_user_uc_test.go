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
	entityPkg "github.com/tomazcx/products-api/pkg/entity"
)

func TestCreateUserUseCase(t *testing.T) {
	repo := &mocks.UserRepositoryMock{}

	repo.On("FindByEmail", mock.Anything).Return((*entity.User)(nil), errors.New("DB ERROR"))
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
	repo.AssertCalled(t, "FindByEmail", createUserDto.Email)
	repo.AssertCalled(t, "Create", mock.Anything)
}

func TestCreateUserUseCase_Email_Already_Registered(t *testing.T) {
	repo := &mocks.UserRepositoryMock{}

	userWithEmail := &entity.User{
		ID:       entityPkg.NewId(),
		Name:     "John Doe",
		Email:    "john@email.com",
		Password: "12345",
	}
	repo.On("FindByEmail", mock.Anything).Return(userWithEmail, nil)
	repo.On("Create", mock.Anything).Return(nil)

	createUserUseCase := usecase.CreateUserUseCase{Repository: repo}

	createUserDto := dto.UserDTO{
		Name:     "John Doe",
		Email:    "john@email.com",
		Password: "12345",
	}
	user, err := createUserUseCase.Execute(createUserDto)

	assert.NotNil(t, err)
	assert.Nil(t, user)
	repo.AssertCalled(t, "FindByEmail", createUserDto.Email)
	repo.AssertNotCalled(t, "Create", mock.Anything)
}

func TestCreateUserUseCase_Fails_Create_User(t *testing.T) {
	repo := &mocks.UserRepositoryMock{}

	repo.On("Create", mock.Anything).Return(errors.New("DB ERROR"))
	repo.On("FindByEmail", mock.Anything).Return((*entity.User)(nil), errors.New("DB ERROR"))

	createUserUseCase := usecase.CreateUserUseCase{Repository: repo}

	createUserDto := dto.UserDTO{
		Name:     "John Doe",
		Email:    "john@email.com",
		Password: "12345",
	}
	user, err := createUserUseCase.Execute(createUserDto)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	repo.AssertCalled(t, "FindByEmail", createUserDto.Email)
	repo.AssertCalled(t, "Create", mock.Anything)
}

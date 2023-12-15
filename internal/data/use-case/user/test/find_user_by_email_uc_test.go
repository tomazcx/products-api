package test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	usecase "github.com/tomazcx/products-api/internal/data/use-case/user"
	"github.com/tomazcx/products-api/internal/entity"
	"github.com/tomazcx/products-api/internal/infra/db/repository/mocks"
	entityPkg "github.com/tomazcx/products-api/pkg/entity"
)

func TestFindUserByEmail(t *testing.T) {
	repo := &mocks.UserRepositoryMock{}

	user := entity.User{
		ID:       entityPkg.NewId(),
		Name:     "John Doe",
		Email:    "john@email.com",
		Password: "1234",
	}
	repo.On("FindByEmail", user.Email).Return(&user, nil)

	findUserByEmailUseCase := usecase.FindUserByEmailUseCase{Repository: repo}
	result, err := findUserByEmailUseCase.Execute("john@email.com")

	assert.Nil(t, err)
	assert.Equal(t, &user, result)
	repo.AssertCalled(t, "FindByEmail", user.Email)
}

func TestFindUserByEmail_Fail_User_Not_Found(t *testing.T) {
	repo := &mocks.UserRepositoryMock{}

	repo.On("FindByEmail", "maria@email.com").Return((*entity.User)(nil), errors.New("DB ERROR"))

	findUserByEmailUseCase := usecase.FindUserByEmailUseCase{Repository: repo}
	result, err := findUserByEmailUseCase.Execute("maria@email.com")

	assert.NotNil(t, err)
	assert.Nil(t, result)
	repo.AssertCalled(t, "FindByEmail", "maria@email.com")
}

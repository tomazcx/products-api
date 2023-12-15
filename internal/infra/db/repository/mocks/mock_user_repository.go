package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/tomazcx/products-api/internal/entity"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) FindByEmail(email string) (*entity.User, error) {
	args := m.Called(email)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *UserRepositoryMock) Create(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

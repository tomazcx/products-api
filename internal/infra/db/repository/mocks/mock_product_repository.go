package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/tomazcx/products-api/internal/entity"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (m *ProductRepositoryMock) Exists(id string) bool {
	args := m.Called(id)
	return args.Bool(0)
}

func (m *ProductRepositoryMock) FindById(id string) (*entity.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Product), args.Error(1)
}

func (m *ProductRepositoryMock) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	args := m.Called(page, limit, sort)
	return args.Get(0).([]entity.Product), args.Error(1)
}

func (m *ProductRepositoryMock) Create(product *entity.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *ProductRepositoryMock) Update(product *entity.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *ProductRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

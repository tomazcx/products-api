package domain

import "github.com/tomazcx/products-api/internal/entity"

type IUserRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

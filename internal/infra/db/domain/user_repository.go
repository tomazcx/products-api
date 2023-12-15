package domain

import "github.com/tomazcx/products-api/internal/entity"

type IUserRepository interface {
	FindByEmail(email string) (*entity.User, error)
	Create(user *entity.User) error
}

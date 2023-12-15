package repository

import (
	"github.com/tomazcx/products-api/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (u *UserRepository) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *UserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User

	if err := u.DB.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomazcx/products-api/internal/entity"
	"github.com/tomazcx/products-api/internal/infra/db/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)

	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("John Doe", "john@email.com", "1234")
	userRepository := repository.NewUserRepository(db)
	err = userRepository.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error

	assert.Nil(t, err)
	assert.Equal(t, userFound.ID, user.ID)
	assert.Equal(t, userFound.Name, user.Name)
	assert.Equal(t, userFound.Email, user.Email)
	assert.NotEmpty(t, userFound.Password)
}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)

	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("John Doe", "john@email.com", "1234")
	err = db.Create(user).Error
	assert.Nil(t, err)

	userRepository := repository.NewUserRepository(db)
	userFound, err := userRepository.FindByEmail("john@email.com")

	assert.Nil(t, err)
	assert.Equal(t, userFound.ID, user.ID)
	assert.Equal(t, userFound.Email, user.Email)
}

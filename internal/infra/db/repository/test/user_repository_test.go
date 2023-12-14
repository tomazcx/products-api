package test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomazcx/products-api/internal/entity"
	"github.com/tomazcx/products-api/internal/infra/db/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestUserRepository() (func(db *gorm.DB), *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entity.User{})

	return func(db *gorm.DB) {
		sqlDB, err := db.DB()

		if err != nil {
			log.Fatalf("Failed to get the database coonection")
		}

		sqlDB.Close()
	}, db
}

func TestCreateUser(t *testing.T) {
	teardown, db := setupTestUserRepository()
	defer teardown(db)

	user, err := entity.NewUser("John Doe", "john@email.com", "1234")
	assert.Nil(t, err)

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
	teardown, db := setupTestUserRepository()
	defer teardown(db)

	user, err := entity.NewUser("John Doe", "john@email.com", "1234")
	assert.Nil(t, err)

	err = db.Create(user).Error
	assert.Nil(t, err)

	userRepository := repository.NewUserRepository(db)
	userFound, err := userRepository.FindByEmail("john@email.com")

	assert.Nil(t, err)
	assert.Equal(t, userFound.ID, user.ID)
	assert.Equal(t, userFound.Email, user.Email)
}

package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomazcx/products-api/internal/entity"
)

func TestNewUser(t *testing.T) {
	user, err := entity.NewUser("John Doe", "john@email.com", "1234")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john@email.com", user.Email)
}

func TestUserValidatePassword(t *testing.T) {
	user, err := entity.NewUser("John Doe", "john@email.com", "1234")
	assert.Nil(t, err)
	assert.NotEmpty(t, user.Password)
	assert.True(t, user.ValidatePassword("1234"))
	assert.False(t, user.ValidatePassword("12345"))
	assert.NotEqual(t, user.Password, "1234")
}

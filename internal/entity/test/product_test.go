package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomazcx/products-api/internal/entity"
)

func TestNewProduct(t *testing.T) {
	product, err := entity.NewProduct("Test Product", 9.9)

	assert.Nil(t, err)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, product.Name, "Test Product")
	assert.Equal(t, product.Price, 9.9)
	assert.NotEmpty(t, product.CreatedAt)
}

func TestProductWhenNameIsEmpty(t *testing.T) {
	_, err := entity.NewProduct("", 9.9)

	assert.NotNil(t, err)
	assert.Equal(t, err, entity.ErrNameIsRequired)
}

func TestProductWhenPriceIsEmpty(t *testing.T) {
	_, err := entity.NewProduct("Test Product", 0.0)

	assert.NotNil(t, err)
	assert.Equal(t, err, entity.ErrPriceIsRequired)
}

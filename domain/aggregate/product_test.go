package aggregate_test

import (
	"bookstore/domain/aggregate"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := aggregate.NewProduct("SOLID", "Principles", 1000.00, 4)
	assert.Nil(t, err, err)
	assert.NotNil(t, product)
}

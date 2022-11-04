package valueobject_test

import (
	valueobject "bookstore/domain/value-object"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func TestNewItemValueObject(t *testing.T) {
	item, err := valueobject.NewItemValueObject(faker.Name(), faker.Word(), 10, 1)
	assert.Nil(t, err, err)
	assert.NotNil(t, item)
}

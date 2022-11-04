package aggregate_test

import (
	"bookstore/domain/aggregate"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func TestNewCustomer(t *testing.T) {
	custumer, err := aggregate.NewCustomer(faker.Name(), faker.Phonenumber(), faker.Email(), faker.Password())

	assert.Nil(t, err)
	assert.NotNil(t, custumer)
}

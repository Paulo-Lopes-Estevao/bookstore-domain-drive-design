package valueobject_test

import (
	valueobject "bookstore/domain/value-object"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPersonValueObject(t *testing.T) {
	person, err := valueobject.NewPersonValueObject("John", "999625342", "Doe@gmail.com", "1234")

	assert.Nil(t, err, err)

	if person.Name != "John" {
		t.Errorf("Name should be 'John', but '%s'", person.Name)
	}
}

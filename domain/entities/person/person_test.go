package person_test

import (
	"bookstore/domain/entities/person"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func TestNewPerson(t *testing.T) {
	p, err := person.NewPerson(faker.Name(), faker.Phonenumber(), faker.Email(), faker.Password())
	assert.Nil(t, err, err)
	assert.NotNil(t, p)
}

func TestPersonVerifyPassword(t *testing.T) {

	pass := faker.Password()

	p, err := person.NewPerson(faker.Name(), faker.Phonenumber(), faker.Email(), pass)
	assert.Nil(t, err, err)

	verify := p.VerifyPassword(pass)
	assert.True(t, verify)
}

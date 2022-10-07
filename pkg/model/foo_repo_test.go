package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFooRepoSave(t *testing.T) {

	repo := NewFooRepository()

	foo := NewFoo(1, "FooOne")
	repo.Save(foo)

	foo2, ok := repo.FindById(1)

	assert.Equal(t, ok, true)
	assert.Equal(t, foo, foo2)
}

func TestFooRepoFindByIdNotFound(t *testing.T) {

	repo := NewFooRepository()

	foo, ok := repo.FindById(1)

	assert.Equal(t, ok, false)
	assert.Nil(t, foo)
}

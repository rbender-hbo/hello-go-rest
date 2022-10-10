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

func TestFooRepoFindAll(t *testing.T) {

	foo1 := NewFoo(1, "FooOne")
	foo2 := NewFoo(2, "FooTwo")
	foo3 := NewFoo(3, "FooThree")

	repo := NewFooRepository()
	repo.Save(foo1)
	repo.Save(foo2)
	repo.Save(foo3)

	allFoo := repo.FindAll()

	assert.Equal(t, 3, len(allFoo))
}

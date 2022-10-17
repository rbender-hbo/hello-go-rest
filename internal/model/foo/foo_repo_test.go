package foo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFooRepoSaveWithExistingId(t *testing.T) {

	repo := NewInMemoryFooRepository()

	foo := NewFooWithId(1, "FooOne")
	repo.Save(foo)

	assert.Equal(t, 1, foo.FooId)

	foo2, ok := repo.FindById(1)

	assert.Equal(t, ok, true)
	assert.Equal(t, foo, foo2)

	assert.Equal(t, repo.LatestFoo, foo)
}

func TestFooRepoSaveWithoutId(t *testing.T) {

	//repo := NewInMemoryFooRepository()
	var repo FooRepository = NewInMemoryFooRepository()

	foo := NewFoo("FooOne")
	repo.Save(foo)

	assert.Equal(t, 1, foo.FooId)

	foo2, ok := repo.FindById(1)

	assert.Equal(t, ok, true)
	assert.Equal(t, foo, foo2)
}

func TestFooRepoSaveWithoutIdAfterSavingWithId(t *testing.T) {

	repo := NewInMemoryFooRepository()

	foo1 := NewFooWithId(3, "FooThree")
	repo.Save(foo1)

	foo2 := NewFoo("FooFour")
	repo.Save(foo2)

	assert.Equal(t, 4, foo2.FooId)

	foundFoo, ok := repo.FindById(4)

	assert.Equal(t, ok, true)
	assert.Equal(t, foundFoo, foo2)
}

func TestFooRepoFindByIdNotFound(t *testing.T) {

	repo := NewInMemoryFooRepository()

	foo, ok := repo.FindById(1)

	assert.Equal(t, ok, false)
	assert.Nil(t, foo)
}

func TestFooRepoFindAll(t *testing.T) {

	foo1 := NewFooWithId(1, "FooOne")
	foo2 := NewFooWithId(2, "FooTwo")
	foo3 := NewFooWithId(3, "FooThree")

	repo := NewInMemoryFooRepository()
	repo.Save(foo1)
	repo.Save(foo2)
	repo.Save(foo3)

	allFoo := repo.FindAll()

	assert.Equal(t, 3, len(allFoo))
}

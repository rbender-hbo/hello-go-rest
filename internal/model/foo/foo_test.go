package foo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFoo(t *testing.T) {

	foo := NewFoo("FooOne")
	assert.Equal(t, 0, foo.FooId)
	assert.Equal(t, "FooOne", foo.Name)
}

func TestNewFooWithId(t *testing.T) {

	foo := NewFooWithId(1, "FooOne")
	assert.Equal(t, 1, foo.FooId)
	assert.Equal(t, "FooOne", foo.Name)
}

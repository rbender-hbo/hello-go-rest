package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFoo(t *testing.T) {

	foo := NewFoo(1, "FooOne")
	assert.Equal(t, 1, foo.FooId)
	assert.Equal(t, "FooOne", foo.Name)
}

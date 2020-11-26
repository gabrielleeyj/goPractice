package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert := assert.New(t)
	a := 1
	b := 1
	assert.NotEqual(a, b, "The two int is not the same.")

	r := Add(a, b)
	assert.EqualValues(r, 3)
}

func TestSubtract(t *testing.T) {
	assert := assert.New(t)
	a := 3
	b := 1
	assert.NotEqual(a, b, "Value is 0")

	r := Subtract(a, b)
	assert.EqualValues(r, 2)
}

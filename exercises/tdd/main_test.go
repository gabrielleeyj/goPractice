package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	a := 1
	b := 2

	var r int
	r = add(a, b)

	var expectedType int

	// Jest
	assert.IsType(t, expectedType, r)
	assert.Equal(t, 3, r)

	r = add(3, 5)

	assert.IsType(t, expectedType, r)
	assert.Equal(t, 8, r)
}

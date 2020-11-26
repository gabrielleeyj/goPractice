package main

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadStringInputIfReadStringReceivesNilReader(t *testing.T) {
	var r *bufio.Reader

	assert.Panics(t, func() {
		readString(r)
	})
}

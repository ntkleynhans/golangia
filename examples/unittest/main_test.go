package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, Add(1,5), 6, "Add should add the two integers correctly")
}

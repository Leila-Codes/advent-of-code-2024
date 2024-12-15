package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	input := LoadInput("example.txt", false)

	assert.Len(t, input, 4)
	assert.Equal(t, input[0], Operation{2, 4})
	assert.Equal(t, input[1], Operation{5, 5})
	assert.Equal(t, input[2], Operation{11, 8})
	assert.Equal(t, input[3], Operation{8, 5})
}

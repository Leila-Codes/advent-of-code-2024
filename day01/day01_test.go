package day01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadInput(t *testing.T) {
	l1, l2 := LoadInput("example.txt")
	assert.Equal(t, l1, []int{3, 4, 2, 1, 3, 3})
	assert.Equal(t, l2, []int{4, 3, 5, 3, 9, 3})
}

func TestPartOne(t *testing.T) {
	res := PartOne(LoadInput("example.txt"))

	assert.Equal(t, 11, res)
}

func TestPartTwo(t *testing.T) {
	res := PartTwo(LoadInput("example.txt"))

	assert.Equal(t, 31, res)
}

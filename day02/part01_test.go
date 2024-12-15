package day02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartOne(t *testing.T) {
	res := PartOne(LoadInput("example.txt"))

	assert.Equal(t, 2, res)
}

package day05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	res := PartOne(LoadInput("example.txt"))

	assert.Equal(t, 143, res)
}

package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartTwo(t *testing.T) {
	res := PartOne(LoadInput("example2.txt", true))

	assert.Equal(t, 48, res)
}

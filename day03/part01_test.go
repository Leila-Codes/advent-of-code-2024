package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	res := PartOne(LoadInput("example.txt", false))

	assert.Equal(t, 161, res)
}

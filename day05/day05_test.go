package day05

import (
	"testing"

	"github.com/leila-codes/algo-ds/set"
	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	pz := LoadInput("example.txt")

	assert.Equal(t, PageOrderRule{
		47: set.New(53, 13, 61, 29),
		97: set.New(13, 61, 47, 29, 53, 75),
		75: set.New(29, 53, 47, 61, 13),
		61: set.New(13, 53, 29),
		29: set.New(13),
		53: set.New(29, 13),
	}, pz.PageOrderRules)

	assert.Equal(t, []UpdateSequence{
		[]int{75, 47, 61, 53, 29},
		[]int{97, 61, 53, 29, 13},
		[]int{75, 29, 13},
		[]int{75, 97, 47, 61, 53},
		[]int{61, 13, 29},
		[]int{97, 13, 75, 29, 47},
	}, pz.NumbersToUpdate)
}

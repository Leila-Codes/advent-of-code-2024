package day05

import (
	"advent-of-code-2024/util"
	"strings"

	"github.com/leila-codes/algo-ds/set"
)

type PageOrderRule map[int]set.Set[int]

type UpdateSequence []int

type PuzzleInput struct {
	PageOrderRules  PageOrderRule
	NumbersToUpdate []UpdateSequence
}

func LoadInput(filename string) (pz PuzzleInput) {
	pz.PageOrderRules = make(PageOrderRule)

	f, reader := util.OpenFile(filename)
	defer f.Close()

	for reader.Scan() {
		line := reader.Text()

		if strings.Contains(line, "|") {
			x, y := util.ParseNumPair(line, "|")

			if _, exists := pz.PageOrderRules[x]; exists {
				pz.PageOrderRules[x].Add(y)
			} else {
				pz.PageOrderRules[x] = set.New[int](y)
			}

		} else if strings.Contains(line, ",") {
			pz.NumbersToUpdate = append(pz.NumbersToUpdate, util.ParseIntList(util.CsvList(line)))
		}
	}

	return
}

package day05

import (
	"advent-of-code-2024/util"
	"fmt"

	"github.com/leila-codes/algo-ds/set"
)

func IsOK(start int, updateSeq []int, dependencies set.Set[int]) bool {
	for x := start; x >= 0; x-- {

		if dependencies.Contains(updateSeq[x]) {
			// fmt.Printf("Whoops page %d should come AFTER %d\n", updateSeq[x], updateSeq[start])
			return false
		}
	}

	return true
}

func PartOne(pz PuzzleInput) int {
	total := 0

	for _, seq := range pz.NumbersToUpdate {

		seqOK := true

		for x, pageNo := range seq {
			if deps, hasDeps := pz.PageOrderRules[pageNo]; hasDeps {
				// fmt.Printf("Page %d has deps: %v\n", pageNo, deps)
				// seek backwards, are there any illegal pages before?
				if !IsOK(x, seq, deps) {
					seqOK = false
				}
			}
		}

		if seqOK {
			middle := util.MiddleNumber(seq)
			fmt.Printf("Middle Page is: %d\n", middle)
			total += middle
		}
	}

	return total
}

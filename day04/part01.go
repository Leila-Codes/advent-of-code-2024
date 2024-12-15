package day04

import "strings"

type Direction [2]int

var (
	Directions = []Direction{
		[2]int{0, 1},
		[2]int{0, -1},
		[2]int{1, 0},
		[2]int{-1, 0},
		[2]int{1, 1},
		[2]int{-1, -1},
	}

	WordToFind = "XMAS"

	matches int
)

func DepthFirstWordSearch(pz PuzzleInput, start, y, x int, match string) {
	width, height := len(pz)-1, len(pz[0])-1

	if strings.HasSuffix(match, "XMAS") {
		matches++
		return
	}

	if x < 0 || x > width || y < 0 || y > height || pz[y][x] == '.' {
		return
	}

	tmp := pz[y][x]
	pz[y][x] = '.'

	for _, dir := range Directions {
		DepthFirstWordSearch(pz, start+1, y+dir[0], x+dir[1], match+string(tmp))
	}

	pz[y][x] = tmp
}

func PartOne(pz PuzzleInput) int {
	DepthFirstWordSearch(pz, 0, 0, 0, "")

	return matches
}

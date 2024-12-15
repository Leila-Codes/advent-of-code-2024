package day04

import (
	"advent-of-code-2024/util"
)

type Row []rune

type PuzzleInput []Row

func LoadInput(filename string) (pz PuzzleInput) {
	f, reader := util.OpenFile(filename)
	defer f.Close()

	for reader.Scan() {
		pz = append(pz, Row(reader.Text()))
	}

	return
}

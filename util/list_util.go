package util

import (
	"math"
	"strconv"
	"strings"
)

func CsvList(list string) []string {
	parts := strings.Split(list, ",")
	for i := 0; i < len(parts); i++ {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

func ParseIntList(list []string) []int {
	ints := make([]int, len(list))

	var err error
	for i := 0; i < len(list); i++ {
		ints[i], err = strconv.Atoi(list[i])
		if err != nil {
			panic(err)
		}
	}

	return ints
}

func ParseNumPair(input, sep string) (a, b int) {
	aStr, bStr, ok := strings.Cut(input, sep)
	if !ok {
		panic("seperator " + sep + " not found input " + input)
	}

	var err error
	a, err = strconv.Atoi(aStr)
	if err != nil {
		panic(err)
	}
	b, err = strconv.Atoi(bStr)
	if err != nil {
		panic(err)
	}

	return
}

func MiddleNumber(seq []int) int {
	middleIdx := math.Floor(float64(len(seq)) / 2)
	return seq[int(middleIdx)]
}

package day01

import (
	"bufio"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

// parseLine parses a single puzzle input line from string to two integers
func parseLine(line string) (a, b int) {
	num1, num2, _ := strings.Cut(line, "   ")
	a, _ = strconv.Atoi(num1)
	b, _ = strconv.Atoi(num2)
	return
}

// LoadInput loads an entire puzzle input file into memory
func LoadInput(filename string) (l1, l2 []int) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewScanner(f)

	for reader.Scan() {
		num1, num2 := parseLine(reader.Text())

		l1 = append(l1, num1)
		l2 = append(l2, num2)
	}

	return
}

// delta calculates the abs diff between a and b.
func delta(a, b int) int {
	return int(math.Abs(float64(b) - float64(a)))
}

// PartOne solution to Day01
func PartOne(l1, l2 []int) int {
	slices.Sort(l1)
	slices.Sort(l2)

	sum := 0
	for i := 0; i < len(l1); i++ {
		sum += delta(l1[i], l2[i])
	}

	return sum
}

// howMany returns how many of item is in the given list
func howMany(list []int, item int) (count int) {
	for _, val := range list {
		if item == val {
			count++
		}
	}

	return count
}

// PartTwo solution to Day01
func PartTwo(l1, l2 []int) int {
	sum := 0

	for i := 0; i < len(l1); i++ {
		sum += l1[i] * howMany(l2, l1[i])
	}

	return sum
}

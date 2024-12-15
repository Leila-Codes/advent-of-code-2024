package day02

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Report []int

// parseLine converts a string text line into a report
func parseLine(line string) (r Report) {
	nums := strings.Split(line, " ")

	for _, vStr := range nums {
		vInt, _ := strconv.Atoi(vStr)
		r = append(r, vInt)
	}

	return
}

// LoadInput loads and parses puzzle input for Day02.
func LoadInput(filename string) (reports []Report) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewScanner(f)
	for reader.Scan() {
		reports = append(reports, parseLine(reader.Text()))
	}

	return
}

package day03

import (
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Operation struct {
	A, B int
}

func parseInt(input string) int {
	val, _ := strconv.Atoi(input)
	return val
}

var (
	reMulOperation  = regexp.MustCompile("mul\\(\\d+,\\d+\\)")
	reMulOperation2 = regexp.MustCompile(`(?P<donot>don't\(\)|do\(\))?.?(?P<op>mul\(\d+,\d+\))`)

	enabled = true
)

const (
	DoOperation = "do()"
	NoOperation = "don't()"
)

func parseMatch2(match string) *Operation {
	parts := reMulOperation2.FindStringSubmatch(match)

	switch len(parts) {
	case 2:
		match = parts[1]
	case 3:
		if parts[1] == NoOperation {
			enabled = false
			return nil
		} else if parts[1] == DoOperation {
			enabled = true
		}

		match = parts[2]
	}

	if !enabled {
		return nil
	}

	return parseMatch(match)
}

func parseMatch(match string) *Operation {
	_, params, _ := strings.Cut(match, "(")

	a, b, _ := strings.Cut(params, ",")
	b = strings.TrimSuffix(b, ")")

	return &Operation{
		A: parseInt(a),
		B: parseInt(b),
	}
}

func LoadInput(filename string, shouldCheckOps bool) (ops []Operation) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(f)

	var matches [][]byte
	if shouldCheckOps {
		matches = reMulOperation2.FindAll(data, -1)
	} else {
		matches = reMulOperation.FindAll(data, -1)
	}

	for _, match := range matches {

		var op *Operation
		if shouldCheckOps {
			op = parseMatch2(string(match))
		} else {
			op = parseMatch(string(match))
		}

		if op != nil {
			ops = append(ops, *op)
		}
	}

	return
}

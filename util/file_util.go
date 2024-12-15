package util

import (
	"bufio"
	"os"
)

func OpenFile(filename string) (*os.File, *bufio.Scanner) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	return f, bufio.NewScanner(f)
}

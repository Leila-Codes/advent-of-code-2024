package day03

func PartOne(input []Operation) int {
	total := 0
	for _, op := range input {
		total += (op.A * op.B)
	}

	return total
}

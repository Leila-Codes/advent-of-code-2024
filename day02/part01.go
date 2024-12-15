package day02

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func isSafe(report Report) bool {
	increasing := true

	for i := 0; i < len(report)-1; i++ {
		if i == 0 {
			if report[i] > report[i+1] {
				increasing = false
			}
		}

		delta := abs(report[i] - report[i+1])

		// if the current report, is greater than the next - we must be increasing
		if delta > 3 {
			return false
		} else if report[i] >= report[i+1] && increasing { // but if we're not increasing, fail
			return false
		} else if report[i] <= report[i+1] && !increasing {
			return false
		}
	}

	return true
}

func PartOne(reports []Report) int {
	count := 0

	for _, report := range reports {
		if isSafe(report) {
			count++
		}
	}

	return count
}

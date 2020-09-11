package iteration

// Sum soma array de numeros int
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll sum for splice
func SumAll(args ...[]int) []int {
	var result []int
	for _, arg := range args {
		result = append(result, Sum(arg))
	}

	return result
}

// SumOfEverythingElse for of every thing else
func SumOfEverythingElse(args ...[]int) []int {
	var result []int
	for _, arg := range args {
		if len(arg) == 0 {
			result = append(result, 0)
			continue
		}

		end := arg[1:]
		result = append(result, Sum(end))
	}

	return result
}

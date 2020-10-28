package arrays

// Sum sum of numbers
func Sum(numbers []int) (sum int) {
	for _, v := range numbers {
		sum += v
	}
	return
}

// SumAll sum each numbers in numbersToSum
func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

// SumAllTails sum each numbers in numbersToSum except first elem in numbers
func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers[1:]))
	}
	return
}

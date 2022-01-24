package arrays

//Sum returns the sum of the elements of the array
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

//SumAll sums all the elements in multiple slices of arrays
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, number := range numbersToSum {
		sums = append(sums, Sum(number))
	}
	return sums
}

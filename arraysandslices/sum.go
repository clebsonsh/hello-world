package arraysandslices

// Sum takes an array of numbers and returns the sum of them
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

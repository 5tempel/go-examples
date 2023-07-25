package arrays

// function takes array of integers, iterates through the array using range and return the sum of all numbers
func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

/*
// function takes array of integers and return the sum of all numbers
func Sum1(numbers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
}
*/

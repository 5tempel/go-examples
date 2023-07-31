package arrays_slices

/*
SumInitial example, using arrays and iterating with a for loop
This function takes an array of integers and return the sum of all numbers.
*/
func SumInitial(numbers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
}

/*
Refactored Sum example, using arrays and iterating with a for loop with a range.
This function takes an array of integers, iterates through the array using range and return the sum of all numbers.
*/
func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

/*
SumSlices example, using slices.
*/
func SumSlices(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

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

/*
SumAll takes a varying number of slices, returns a new slice containing the totals for each slice passed in.
*/
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = SumSlices(numbers)
	}

	return sums
}

/*
SumAllAppend takes a number of slices and appends a sum of elements in each slice to an empty slice
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices#refactor-2
*/
func SumAllAppend(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, SumSlices(numbers))
	}

	return sums
}


/*
SumAllTails takes a number of slice and removes the head item from each slice and then appends the sum of all elements in each slice to an empty slice, 
returning a slice which is a sum of the tail items of the original slices. If a slice is empty it will append a slice with one item equal 0.
*/
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			//Slices can be sliced! The syntax is slice[low:high]. 
			//If you omit the value on one of the sides of the : it captures everything to that side of it. 
			//In our case, we are saying "take from 1 to the end" with numbers[1:]. 
			tail := numbers[1:]	
			sums = append(sums, SumSlices(tail))
		}
	}

	return sums
}

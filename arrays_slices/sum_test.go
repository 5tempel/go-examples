package arrays_slices

import (
	"fmt"
	"testing"
)

/*
Arrays
*/
func TestSum(t *testing.T) {

	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

// example function, see documentation for more info: https://pkg.go.dev/testing#hdr-Examples
func ExampleSum() {
	a1 := Sum([5]int{2, 3, 1, 5, 9})
	fmt.Println(a1)
	//Output: 20
}

// benchmark function runs b.N times and measures how long it takes, see documentation for more info:
// https://pkg.go.dev/testing#hdr-Benchmarks,
// https://gobyexample.com/testing-and-benchmarking
func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([5]int{2, 3, 1, 5, 9})
	}
}

/*
	Slices.
	NOTE: to see test coverage in GO, you can run: go test -cover
*/

func TestSumSlices(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := SumSlices(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := SumSlices(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

// Next topic is to write a new function called SumAll: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices#write-the-test-first-2

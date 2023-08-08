package arrays_slices

import (
	"fmt"
	"testing"
	"reflect"
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


// TestSumAll - https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices#write-the-test-first-2
func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9}, []int{20, 5, 4, 33}, []int{-5, -20, 0, 2})
	
	// Slice can only be cmpared to nil
	t.Run("testing if slice is nil", func(t *testing.T) {
		
		if got == nil {
		t.Errorf("got = %d is nil", got)
		}
	})

	// reflect.DeepEqual is useful for seeing if any two variables are the same. 
	// Note reflect.DeepEqual is not "type safe" - the code will compile even if you did something a bit silly
	t.Run("comparing two slices with reflect.DeepEqual", func(t *testing.T) {
		want := []int{3, 9, 62, -23}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	// Next time start with Refactor https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices#refactor-2

}
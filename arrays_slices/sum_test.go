package arrays_slices

import (
	"fmt"
	"reflect"
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

// SumAllAppend
func TestSumAllAppend(t *testing.T) {

	got := SumAllAppend([]int{1, 2}, []int{0, 9}, []int{20, 5, 4, 33}, []int{-5, -20, 0, 2})

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
		//want := "[3 9 62 -23]"	/    sum_test.go:118: got [3 9 62 -23] want [3 9 62 -23]
		//--- FAIL: TestSumAllAppend (0.00s) /
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}

// SumAllTails
func TestSumAllTails(t *testing.T) {

	//assigning a function to a variable
	//By defining this function inside the test, it cannot be used by other functions in this package.
	//Hiding variables and functions that don't need to be exported is an important design consideration.
	//A handy side-effect of this is this adds a little type-safety to our code.
	//If a developer mistakenly adds a new test with checkSums(t, got, "dave") the compiler will stop them in their tracks.
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})

	/* This test will fail to compile with the following error:
		--
		arrays_slices\sum_test.go:154:3: want declared and not used
		arrays_slices\sum_test.go:155:21: cannot use "dave" (untyped string constant) as []int value in argument to checkSums
		FAIL    go-examples/arrays_slices [build failed]
		--
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, "dave")

	})
	*/

	// recap from https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices#wrapping-up
}

/*
This is TDD practice, iterating over our solutions to basic mathematic problems and learning new language features motivated by our tests.
- Declaring structs to create your own data types which lets you bundle related data together and make the intent of your code clearer
- Declaring interfaces so you can define functions that can be used by different types (parametric polymorphism)
- Adding methods so you can add functionality to your data types and so you can implement interfaces
- Table driven tests to make your assertions clearer and your test suites easier to extend & maintain
*/
package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want) // f is for float64 and the .2 means print 2 decimal places
	}
}

func TestArea(t *testing.T) {
	got := Area(12.0, 6.0)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// Struct https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces#refactor

func TestStructPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := StructPerimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestStructArea(t *testing.T) {
	rectangle := Rectangle{12.0, 6.0}
	got := StructArea(rectangle)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestMethodArea(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		got := rectangle.Area()
		want := 72.0

		if got != want {
			//Use of g will print a more precise decimal number in the error message (fmt options).
			//For example, using a radius of 1.5 in a circle area calculation, f would show 7.068583 whereas g would show 7.0685834705770345.
			t.Errorf("got %g want %g", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

}

func TestInterfaceArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

}

// Decoupling  https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces#decoupling
/*
 Using table driven test
 https://github.com/golang/go/wiki/TableDrivenTests
*/
func TestTableArea(t *testing.T) {

	// creating an anonymous struct, declaring a slice of structs by using []struct with two fields, the shape and the want
	areaTests := []struct {
		shape   Shape
		hasArea float64
	}{
		// fill the slice with cases
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		// you can optionally name the fields
		{shape: Triangle{Base: 15, Height: 8}, hasArea: 60},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			// error output example: structs_test.go:117: structs.Circle{Radius:2} got 12.566370614359172 want 314.1592653589793 / FAIL: TestTableArea (0.00s)
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}

}

// Next time start with Pointers & errors - https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors

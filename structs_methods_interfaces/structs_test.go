package structs
import (
	"testing"
)
	
func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)	// f is for float64 and the .2 means print 2 decimal places
	}
}

func TestArea(t *testing.T) {
	got := Area(12.0, 6.0)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// Next time start with Refactor struct https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces#refactor
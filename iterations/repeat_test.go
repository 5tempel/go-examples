package iterations

import (
	"fmt"
	"log"
	"testing"
)

/*
helper function accepts a "test", the actual and expected value
and fails the test with an error message when the expected value
doesn't match the actual value
*/
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestRepeatN(t *testing.T) {
	repeated := RepeatN("z")
	expected := "zzzzz"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestCharValidation(t *testing.T) {

	t.Run("testing when a single character is supplied", func(t *testing.T) {
		got, err := CharValidation("b")
		want := "b"
		if err != nil {
			log.Fatal(err)
		}
		assertCorrectMessage(t, got, want)
	})

	t.Run("testing when a multiple characters are supplied", func(t *testing.T) {
		got, err := CharValidation("abc")
		want := "This is not a character"
		got = fmt.Sprintf("%s", err)
		assertCorrectMessage(t, got, want)
	})

	t.Run("testing when no character is supplied", func(t *testing.T) {
		validated, err := CharValidation("")
		var got string
		want := "No character was provided"
		if validated == "" {
			got = fmt.Sprintf("%s", err)
		}
		assertCorrectMessage(t, got, want)
	})

	/*	t.Run("testing when a multiple characters are supplied", func(t *testing.T) {
			got, err := CharValidation("abc")
			want := "This is not a character"
			if err != nil {
				log.Fatal(err)
			}
			assertCorrectMessage(t, got, want)
		})
	*/
}

func ExampleRepeat() {
	r1 := RepeatN("x")
	fmt.Println(r1)
	// Output: xxxxx

}

/*
	r2 := RepeatN("")
	fmt.Println(r2)
	// Output2: "No character was provided"

	r3 := RepeatN("abc")
	fmt.Println(r3)
	// Output3: "This is not a character"
*/

// benchmark code runs b.N times and measures how long it task
func BenchmarkRepeatN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatN("a")
	}
}

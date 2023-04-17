package main

import "fmt"

func Hello(name string) string {
/* 	s1:= "test2"
	h:= fmt.Sprintf("Hello, %s %s", name, s1)
	h:= fmt.Sprintf("Hello, %s", name)
 */
/* 	var s2 string	// declare string
	s2 = ", "		// initialize variable, note only = was used
	// s2:=", "		// decalre and initialize, :=
	h:= "Hello" + s2 + name // decalre and initialize (string concat) */

	/* const s3 = "Hello, "
	h:= s3 + name */

	var s4 string = "Hello, " // homework var s4 = "Hello, "
	h:= s4 + name
	return h

// got %q want %q", got, want
// return "Hello, " + name

}

func main() {
	fmt.Println(Hello("apple"))
}

// Shift+Alt+/
/* func main() {
	fmt.Println(Hello("apple"))
} */
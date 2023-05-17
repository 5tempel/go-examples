package main

import "fmt"

const englishHelloPrefix = "Hello, "

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

func HelloEmptyString(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}


func HelloLanguage(name, language string) string {
	greeting := "Hello, "
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		greeting = "Hola, "
	} else if language == "Polish" {
		greeting = "Czesc, "
	}

	return greeting + name
}

func HelloLSwitch(name, language string) string {
	greeting := "Hello, "
	if name == "" {
		name = "World"
	}
	switch language {
	case "Spanish" : greeting = "Hola, "
	case "Polish" : greeting = "Czesc, "
	default : greeting = "Hello, "
	}

	return greeting + name
}

func main() {
	fmt.Println(Hello("apple"))
	fmt.Println(HelloEmptyString("EWA"))
	fmt.Println(HelloLanguage("EWA", "Polish"))
}

// Shift+Alt+/
/* func main() {
	fmt.Println(Hello("apple"))
} */
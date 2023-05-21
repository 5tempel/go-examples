package iteration

const numberOfRepeats = 5

// accept a character and return it repeated numberOfRepeates times
func RepeatN(c string) string {
	var repeatedC string
	for i := 0; i < numberOfRepeats; i++ {
		repeatedC += c
	}
	return repeatedC
}

// accept a character and return it repeated 5 times
func Repeat(c string) string {
	var repeated string // decalre string variable
	for i := 0; i < 5; i++ {
		repeated = repeated + c
	}
	return repeated
}

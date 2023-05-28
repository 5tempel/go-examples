
package iteration
import (
	"fmt"
	"errors"
	"log"
)


const numberOfRepeats = 5

// accept a character and return it repeated numberOfRepeates times
func RepeatN(c string) string {
	var repeatedC string
	validated, err := CharValidation(c)
	if err != nil {
		// %v
		//repeatedC = fmt.Sprintf("%s",err) 
		log.Fatal(fmt.Sprintf("%s",err) )
	} else {
		for i := 0; i < numberOfRepeats; i++ {
			repeatedC += validated
		}
	}
	return repeatedC
}

// accept a character and validates if it's a single character
func CharValidation(c string) (validated string, err error) {

	length := len(c)

	if length == 0 {
		err = errors.New("No character was provided") 
	}
	if length > 1 {
		err = errors.New("This is not a character")
	}

	return c, err
}

// accept a character and return it repeated 5 times
func Repeat(c string) string {	
	var repeated string // decalre string variable
	for i := 0; i < 5; i++ {
		repeated = repeated + c
	}
	return repeated
	
}

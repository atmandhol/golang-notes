package main

import (
	"fmt"
)

var (
	// Basic types are bool, string, int8,16,32,64 and float32,64
	Python           bool    = true           // Default is false
	AwesomeLanguage  string  = "Python Rules" // Defaults to ""
	OneOfTen         int8    = 10
	BigOne           int64            // Defaults to 0
	BigOneWithPoints float64 = 3.1415 // Defaults to 0
)

func main() {

	// Type conversion is pretty straight forward T(v)
	// := is the assignment operator
	// In Go, := is for declaration + assignment, whereas = is for assignment only.
	BigOneWithNoPoints := int64(BigOneWithPoints)
	// is same as var BigOneWithNoPoints int64 = int64(BigOneWithPoints)

	fmt.Println(OneOfTen)
	fmt.Println(BigOne)
	fmt.Println(BigOneWithPoints)
	fmt.Println(BigOneWithNoPoints)
}

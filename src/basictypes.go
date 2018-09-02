package main

import (
	"fmt"
)

var (
	Python bool = true	// Default is false
	AwesomeLanguage string = "Python Rules" // Defaults to ""
	OneOfTen int8 = 10
	BigOne int64 // Defaults to 0
	BigOneWithPoints float64 = 3.1415  // Defaults to 0
)

func main() {

	// Type conversion is pretty straight forward T(v)

	BigOneWithNoPoints := int64(BigOneWithPoints)

	fmt.Println(OneOfTen)
	fmt.Println(BigOne)
	fmt.Println(BigOneWithPoints)
	fmt.Println(BigOneWithNoPoints)
}

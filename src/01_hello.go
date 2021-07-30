package main

// fmt is the package used for input/output functions i.e. PrintLn
import "fmt"

var fact string = "Python Rules!"

func add(x int, y int) int {
	return x + y
}

func someMath(x, y int) (int, int) {
	return x + y, x - y
}

// Program start running in package main
// // is used for single line comment
// /* */ is used for multiline comment
func main() {
	fmt.Println(fact)
	fmt.Println(someMath(5, 6))
}

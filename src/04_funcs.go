package main

import "fmt"

// Creating a function by using func keyword
/*
Format is

func <name>(arg1 <type of arg1>, arg2, arg3 <type of arg2,3>) (x return-type) {

}
*/
func add(x, y float32) float32 {
	return x + y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// if nothing is specified after return, named outputs are returned
	// in this case, x and y are named outputs and they are returned in that order
	return
}

func main() {
	fmt.Println("Sum of 2 + -6.5 is", add(2, -6.5))
}

package main

import "fmt"

var fact string = "Python Rules!"

func add(x int, y int) int {
	return x + y
}

func someMath(x, y int) (int, int) {
	return x + y, x - y
}

func main() {
	fmt.Println(fact)
	fmt.Println(someMath(5,6))
}

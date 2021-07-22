package main

import "fmt"

func add(x, y float32) float32 {
	return x + y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("Sum of 2 + -6.5 is", add(2, -6.5))
}

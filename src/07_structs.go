package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	x := Vertex{1, 2}
	x1 := Vertex{}
	x2 := Vertex{X: 1}

	fmt.Println(x)
	fmt.Println(x1)
	fmt.Println(x2)

	// Here, point_x is a pointer to x which can be used for pass by reference
	point_x := &x
	// Here a new object y is created
	y := x

	fmt.Println(point_x.X)
	x.X = 12
	fmt.Println(y.X)
	fmt.Println(point_x.X)

}

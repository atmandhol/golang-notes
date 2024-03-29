package main

import (
	"fmt"
	"math"
)

// An interface type is defined as a set of method signatures.
// A value of interface type can hold any value that implements those methods.
// so a perfect example of interface is cache machine
// it can be implemented by localcache, redis and memcached

/*
One of the most common interface defined by fmt is Stringer
type Stringer interface {
    String() string
}
*/

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v

	fmt.Println(a.Abs())
}

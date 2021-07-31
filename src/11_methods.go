package main

import (
	"fmt"
	"math"
)

// Go does not have classes, but you can add methods to type structs

/*
* Normal func format
func <name>(arg1 <type of arg1>, arg2, arg3 <type of arg2,3>) (x return-type) {
}

* Method func format
func (<var> <type>) <name>(arg1 <type of arg1>, arg2, arg3 <type of arg2,3>) (x return-type) {
}
*/
type Coordinate struct {
	lat, lon float32
}

func (c Coordinate) flip_and_new() Coordinate {
	return Coordinate{c.lon, c.lat}
}

func (c *Coordinate) flip_current() *Coordinate {
	t := c.lat
	c.lat = c.lon
	c.lon = t
	return c
}

// Can be used on non-struct types too

type Float32 float32

func (f Float32) square() Float32 {
	return f * f
}

func main() {

	atl := Coordinate{33.123, -84.234}
	fmt.Println(atl)
	fmt.Println(atl.flip_and_new())
	fmt.Println(atl)
	// Print pointer to atl here cause flip_current returns &address
	fmt.Println(*atl.flip_current())
	fmt.Println(atl)

	f1 := Float32(math.Pi)
	fmt.Println(f1)
	fmt.Println(f1.square())

}

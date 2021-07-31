package main

import "fmt"

// Arrays
/*
var <var>[size]type{values,}
*/

func main() {
	cities := [3]string{"Atlanta", "NYC", "LA"}
	fmt.Println(cities)

	//	An array has a fixed size.
	// A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.
	// In practice, slices are much more common than arrays.
	// slice is a reference to the underlying array

	city_slice := cities[1:2]
	fmt.Println(city_slice)

	// len = size of the slice
	fmt.Println(len(city_slice))
	// cap = size of the array under it, starting from first elem of slice
	fmt.Println(cap(city_slice))

	// this is a slice literal. it creates an array underneath and then creates a slice that refs it
	states := []string{"GA", "NY", "CA"}
	fmt.Println(states)

	var nils []int
	fmt.Println("length:", len(nils), "capacity", cap(nils))

	// you can use make function to make a slice
	// format is make([]type, len, cap)
	make_val := make([]int, 5, 10)
	fmt.Println(make_val)
}

package main

import "fmt"

func main() {
	// You can use make (map[keytype]valtype) to create a new map
	m := make(map[string]string)
	n := map[string]string{"foo": "bar"}
	m["foo"] = "bar"
	n["zha"] = "gabar"
	fmt.Println(m, n)

	// for over map
	for key, val := range n {
		fmt.Println(key, val)
	}

	// delete a key
	delete(n, "zha")
	fmt.Println(m, n)

	// check if key exists
	elem, ok := n["zha"]
	if ok {
		fmt.Println(elem)
	} else {
		fmt.Println("Elem not found!")
	}

}

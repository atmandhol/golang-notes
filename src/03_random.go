package main

import (
	"fmt"
	"math"

	// For random functions, use math/rand
	"math/rand"
	// Well for time stuff
	"time"
)

func main() {
	// Way to get the unix epoch timestamp
	fmt.Println(time.Now().Unix())
	// Set the seed to timestamp
	rand.Seed(time.Now().Unix())
	// Get the random integer with set seed
	fmt.Println("My favorite number is ", rand.Intn(10))
	// Functions/objects name that starts with cap letter is an exported function
	// used to make it publically accessible. for e.g. math.Pi <- Pi being exported here
	fmt.Println(math.Pi)
}

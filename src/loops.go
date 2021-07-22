package main

import (
	"fmt"
	"math"
	"runtime"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum", sum)

	if ans := math.Sqrt(16); ans < 10 {
		fmt.Println("Fancy ass if statement where you can run something first")
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS X")
	case "linux":
	default:
		fmt.Println("Something else")
	}
}

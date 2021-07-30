package main

import (
	"fmt"
	"math"

	// Runtime gives you info about your application runtime env
	"runtime"
)

func main() {
	sum := 0
	/*
		for init; end condition; counter change {
			// init and counter change are not needed
		}
	*/
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum", sum)

	if ans := math.Sqrt(16); ans < 10 {
		fmt.Println("Fancy ass if statement where you can run something first")
	}

	/*
		switch <var> := condition; <var> {
		case <val1>:
			...
		case <val2>:
		default:
			...
		}
	*/

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS X")
	case "linux":
	default:
		fmt.Println("Something else")
	}
}

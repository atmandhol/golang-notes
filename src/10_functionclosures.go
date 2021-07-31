package main

import (
	"fmt"
)

// in this case, every object of adder (e.g. pos, neg) stores the state in sum
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// fibonacci closure function
func fibonacci() func() int {
	curr := 0
	next := 0
	return func() int {
		if curr == 0 && next == 0 {
			next = 1
			return 0
		} else {
			old_curr := curr
			curr = next
			next += old_curr
			return next
		}
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

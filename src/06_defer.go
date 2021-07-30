package main

import "fmt"

func deferred_greeting() string {
	return "A late Hello!"
}

func main() {
	var greet = deferred_greeting()
	// defer can be used to call a function that will execute after the control
	// goes out of the scope where defer was invoked
	// in this case, when main() ends
	defer fmt.Println(greet)
	fmt.Println("Thing that came after will print first")
}

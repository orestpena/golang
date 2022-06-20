package main

import "fmt"

// in this example we don't save the function in a variable, it is just simply defined.
func main() {
	x := func(l int, b int) int {
		return l * b
	}(20, 30)
	fmt.Printf("%T \n", x)
	fmt.Println(x)
}

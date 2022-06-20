package main

import "fmt"

//anonymous functions
/*
 * An anonymous function is a function that is declared without any named identifier to refer to it.
 * They can accept inputs and return outputs, just as standard functions do.
 * They can be used for containing functionality that need not be named and possibly for short-term use.
 */
// function inside function

func main() {
	x := func(l int, b int) int {
		return l * b
	}
	fmt.Printf("%T \n", x)
	fmt.Println(x(20, 30))
}

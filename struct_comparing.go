package main

import "fmt"

/*
 comparing structs
	* structs of the same type can be compared using Go's equality operators.
	*    ==     !=
*/

type s1 struct {
	x int
}

type s2 struct {
	x int
}

func main() {
	c := s1{x: 5}
	c1 := s2{x: 5}
	if c == c1 {
		fmt.Println("yes")
	}
	// this will provide an error. It is not possible to compare structs of different types.  One is of s1 and the other is of s2.
	// this will result an a compile time error:  invalid operation: c == c1 (mismatched types s1 and s2)
}

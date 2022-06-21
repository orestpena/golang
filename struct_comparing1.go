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

func main() {
	c := s1{x: 5}
	c1 := s1{x: 6}
	c2 := s1{x: 5}
	if c != c1 {
		fmt.Println("c and c1 have different values")
	}
	if c == c2 {
		fmt.Println("c is same as c2")
	}
}

package main

import "fmt"

func main() {
	s := 100
	var ptr *string = &s
	//var ptr *int = &s
	fmt.Println(s)
	*ptr += 100
	fmt.Println(s)
	// this produces an error due to mismatch between string and int
	// the correct answer is commented out. Comment the var, and uncomment the 2nd var.
}

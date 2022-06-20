package main

import "fmt"

/*
naming convention for functions
	* must begin with a letter.
	* can have any number of additional letters and symbols.
	* cannot contain spaces.
	* case-sensitive.
*/

func addNumbers(a int, b int) int {
	sum := a + b
	return sum
}

func main() {
	sumOfNumbers := addNumbers(2, 3)
	fmt.Println(sumOfNumbers)
}

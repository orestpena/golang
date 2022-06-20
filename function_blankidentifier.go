package main

import "fmt"

func f() (int, int) {
	return 42, 53
}

/*
func main() {
	a, b := f()
	fmt.Println(a, b)
}
*/

/*
func main() {
	v := f()
	fmt.Println(v)
	// this will provide an error b/c it sends back 2 ints
	// so this is why you use the blank identifier '_'
}
*/

func main() {
	v, _ := f()
	fmt.Println(v)
	// this is using the blank identifier.
}

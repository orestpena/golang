package main

import "fmt"

func doSomething(int, int) int {
	//return "2"    // returns error b/c it is a string and it is expecting an int
	return 2
}

func main() {
	fmt.Println(doSomething(1, 2))
}

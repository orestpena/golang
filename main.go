package main

import "fmt"

// to install go: https://go.dev/doc/install

func main() {
	fmt.Println("Hello World!")
	// To execute this program:  go run main.go
	/*
		var a, b int = 10, 10
		fmt.Println(a <= b)
		fmt.Println(a > b)
	*/
	//var a string = "cat"
	var a int = 12
	var b int = 12
	fmt.Println(a == b)

	a = 12
	fmt.Println(a <= b)

	a = 20
	fmt.Println(a <= b)

	b = 100
	fmt.Println(a <= b)

	c := 0
	fmt.Println(a <= c)
}

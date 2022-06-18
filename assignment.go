package main

import "fmt"

func main() {
	var x, y float64 = 27.9, 7.0
	x -= y
	fmt.Println(x)
	x += y
	fmt.Println(x)
}

package main

import "fmt"

func main() {

	var x, y int = 12, 25
	z := x ^ y
	fmt.Println(z)

	x = 212
	z = x << 1
	fmt.Println(z)

	x = 212
	z = x >> 2
	fmt.Println(z)

	x, y = 100, 90
	fmt.Println(x & y)
	fmt.Println(x | y)

	x, y = 100, 90
	fmt.Println((x + y) >> 2)
}

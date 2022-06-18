package main

import "fmt"

func main() {
	var a bool = false
	result := 10 > 50
	fmt.Println(!(a && result))
}

package main

import "fmt"

func main() {
	arr := []int{10, 20, 30, 40, 50}
	for index, value := range arr {
		fmt.Println(index, "=>", value)
	}
	fmt.Println()
	// if you do not need the index, but only the value use the underscore
	for _, value := range arr {
		fmt.Println(value)
	}
}

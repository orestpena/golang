package main

import "fmt"

func main() {
	var fruit string = "grapes"
	if fruit == "apples" {
		fmt.Println("Fruit is apple")
	} else {
		fmt.Println("Fruit is not apple")
	}

	if fruit == "apple" {
		fmt.Println("I love apples")
	} else if fruit == "orange" {
		fmt.Println("Oranges are not apples")
	} else {
		fmt.Println("no appetite")
	}
}

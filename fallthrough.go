package main

import "fmt"

func main() {
	var i int = 10
	switch i {
	case -5:
		fmt.Println("-5")
	case 10:
		fmt.Println("10")
		fallthrough
	case 20:
		fmt.Println("20")
		fallthrough
	case 30:
		fmt.Println("30")
	default:
		fmt.Println("default")
	}
}

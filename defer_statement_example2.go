package main

import "fmt"

func greetings() (x, y string) {
	x = "hello "
	y = "world"
	return
	//or
	//return x, y
}

func main() {
	fmt.Print(greetings())
}

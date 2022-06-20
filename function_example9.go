package main

import (
	"fmt"
	"strings"
)

func main() {
	x := func(s string) {
		fmt.Println(strings.ToLower(s))
	}("RacheL")
	fmt.Printf("%T \n", x)
}

// this program has an error, it has no value and used as value.

package main

import "fmt"

func print(n int) {
	if n == -6 {
		return
	}
	// base case
	fmt.Printf("%d ", n*n)
	print(n - 1)
}

func main() {
	print(2)
}

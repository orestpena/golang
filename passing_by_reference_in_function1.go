package main

import "fmt"

// slices and maps are passed by reference by default.

func modify(s []int) {
	s[0] = 100

}

func main() {
	slice := []int{10, 20, 30}
	fmt.Println(slice)
	modify(slice)
	fmt.Println(slice)
}

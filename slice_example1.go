package main

import "fmt"

func main() {
	arr := []int{10, 20, 30, 40, 50}
	slice := make([]int, 10)
	num := copy(slice, arr)
	fmt.Println(slice)
	fmt.Println(num)
}

package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 30, 90, 100}
	new_slice := append(arr[:3], arr[4:]...)
	fmt.Println(new_slice)
}

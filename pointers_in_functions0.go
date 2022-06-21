package main

import "fmt"

func modify(numbers ...int) {
	for i := range numbers {
		numbers[i] -= 5
	}
}

func main() {
	arr := []int{10, 20, 30}
	fmt.Println(arr)
	modify(arr...)
	fmt.Println(arr)
}

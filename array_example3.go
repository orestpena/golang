package main

import "fmt"

func main() {

	arr := [10]int{10, 20, 30, 40, 50}
	fmt.Println(arr)
	fmt.Println(len(arr))

	fmt.Println(arr[0])
	fmt.Println(arr[2])
	fmt.Println(arr[4])
	fmt.Println(arr[8])
	//fmt.Println(arr[10]) //this is out of bounds the max is 9
	fmt.Println(arr[9])

	arr1 := [5]string{"a", "b", "c"}
	for index, element := range arr1 {
		fmt.Println(index, ">=", element)
	}
}

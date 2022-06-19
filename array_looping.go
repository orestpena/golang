package main

import "fmt"

func main() {

	var grades [5]int = [5]int{90, 80, 70, 80, 97}
	fmt.Println(grades)

	for i := 0; i < len(grades); i++ {
		fmt.Println(grades[i])
	}

	for index, element := range grades {

		fmt.Println(index, "=>", element)
	}
}

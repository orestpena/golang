package main

import "fmt"

func main() {

	var fruits [5]string = [5]string{"apples", "oranges", "grapes", "mango", "papaya"}
	fmt.Println(fruits[3])
	//the index options are between 0-4 for an array of size 5.

	var grades [5]int = [5]int{90, 80, 70, 80, 97}
	fmt.Println(grades)
	//changing the data in an array index
	grades[1] = 100
	fmt.Println(grades)
}

package main

import "fmt"

func main() {
	var grades [5]int
	fmt.Println(grades)
	//initialize the array, so it will only have zeroes
	var fruits [3]string
	fmt.Println(fruits)
	//initialize the array, with 3 empty strings

	/*
		//now to initialize
		var grades [3]int = [3]int{10,20,30}
		fmt.Println(grades)

		//short hand
		grades := [3]int{10,20,30}
		fmt.Println(grades)

	    // short hand without specifying the array size
		grades := [...]int{10,20,30}
		fmt.Println(grades)
	*/

	//initialized array
	var fruit [2]string = [2]string{"apples", "oranges"}
	fmt.Println(fruit)

	marks := [3]int{10, 20, 30}
	fmt.Println(marks)

	names := [...]string("Rachel", "Phoebe", "Monica")
	fmt.Println(names)
}

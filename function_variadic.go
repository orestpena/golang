package main

import "fmt"

// variadic functions
// * function that accepts variable number of arguements.
// * it is possible to pass a varying number of arguments of the same type as
//   referenced in the function signature.
// * to declare a variadic function, the type of the final parameter is preceded by an ellipsis "..."
// * Example - fmt.Println method

func sumNumbers(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func main() {
	fmt.Println(sumNumbers) //this will print an address
	fmt.Println(sumNumbers())
	fmt.Println(sumNumbers(10))
	fmt.Println(sumNumbers(10, 20))
	fmt.Println(sumNumbers(10, 20, 30, 40, 50))
}

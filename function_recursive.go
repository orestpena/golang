package main

// perform a factorial of 5*4*3*2*1
import "fmt"

func factorial(n int) int {
	if n == 1 {
		//fmt.Println(n)
		//fmt.Println("inside if statement of 1")
		return 1
	}
	//fmt.Println(n)
	return n * factorial(n-1)
}

func main() {
	n := 5
	result := factorial(n)
	fmt.Println("Factorial of", n, "is :", result)
}

/*
     factorial(5) =120
return 5 * factorial(4) = 120
          ^
		  |
	return 4 * factorial(3) = 24
	          ^
			  |
		return 3 * factorial(2) = 6
		          ^
				  |
			return 2 * factorial(1) = 2
			          ^
					  |
					  1
*/

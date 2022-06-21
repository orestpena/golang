package main

import "fmt"

/*
passing by reference in functions
  * Go supports pointers, allowing you to pass references to values within your program.

  * In call by reference/pointer, the address of the variable is passed into the function call as the actual parameter.

  * All the operations in the function are performed on the value stored at the address of the actual parameter
*/

func modify(s *string) {
	*s = "world"

}

func main() {
	a := "hello"
	fmt.Println(a)
	modify(&a)
	fmt.Println(a)
}

/*
  memory address  |  memory
a a   0x0301      |   "hello"  --> "world"
      0x0302      |
      0x0303      |
      0x0304      |

when the value gets passed to a function it moves the value information
to another memory location and modifies that one. It does not modify the original.

*/

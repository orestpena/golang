package main

import "fmt"

/*
passing by value in functions
	* Function is called by directly passing the value of the variable as an argument.
	* the parameter is copied into another location of your memory.
	* So, when accessing or mmodifying the variable within your function, only the copy is accessed or modified,
	  and the original value is never modified.
	* All basic types (int, float, bool, string, array) are passed by value.
*/

func modify(a int) {
	a += 100
}

func main() {
	a := 10
	fmt.Println(a)
	modify(a)
	fmt.Println(a)
}

/*
  memory address |  memory
a     0x0301     |   10
      0x302      |
a     0x303      |   10 --> 110
      0x304      |

when the value gets passed to a function it moves the value information
to another memory location and modifies that one. It does not modify the original.

*/

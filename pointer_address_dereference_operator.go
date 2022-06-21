package main

import "fmt"

/*
pointers
	* A pointer is a variable that holds memory address of another variable.
	* They point where the memory is allocated and provide ways to find or even
	  change the value located at the memory location.

address and dereference operators
	& operator - The address of a variable can be obtained by preceding the name of a
	variable with an ampersand sign (&), known as address-of operator.

	* operator - It is known as the dereference operator. When placed before an address,
	it returns the value at that address.

	       									memory address | memory
	x := 77 ------------------------------>     0x0301     |  77
                                                0x0302     |

	&x = 0x0301      *0x0301 = 77
*/

func main() {
	i := 10
	fmt.Printf("%T %v \n", &i, &i)
	fmt.Printf("%T %v \n", *(&i), *(&i))
}

// the asterik in the answer *int is different than the derefernce operator
// *int is how pointers are expressed in golang

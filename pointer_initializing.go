package main

import "fmt"

/*
initializing a pointer
i.e.
   i := 10
   var ptr_i *int = &i

i.e.
   s := "hello"
   var ptr_s = &s

   variables are variables that store the memory address of another variable.
*/

func main() {
	s := "hello"
	var b *string = &s
	fmt.Println(b)
	var a = &s
	fmt.Println(a)
	c := &s
	fmt.Println(c)
}

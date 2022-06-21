package main

import "fmt"

/*
struct
	* user-defined data type.
	* a structure that groups together data elements.
	* provide a way to reference a series of grouped values through a single variable name.
	* used when it makes sense to group or associate two or more data variables.
*/

type Student struct {
	name   string
	rollNo int
	marks  []int
	grades map[string]int
}

func main() {
	var s Student
	fmt.Printf("%+v", s)
}

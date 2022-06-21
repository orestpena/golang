package main

import "fmt"

/*
struct - accessing fields
    you can access fields of a struct by using the . operator
    <variable_name>.<field_name>
*/

type Circle struct {
	x      int
	y      int
	radius int
}

func main() {
	var c Circle
	c.x = 5
	c.y = 5
	c.radius = 5
	fmt.Printf("%+v \n", c)
	// fmt.Printf(%+v \n", c.area)
	// The commented portion will create an error b/c c.area is undefinded. It is not part of the struct.
	// type Circle has no field or method area.
}

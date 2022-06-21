package main

import "fmt"

/*
  Methods
      * A method augments a function by adding an extra parameter section immediately after the 'func'
	    keyword that accepts a single arguement.
	  * This arguement is called a 'receiver'.
	  * A method is a function that has a defined receiver.
*/

type Circle struct {
	radius float64
	area   float64
}

func (c *Circle) calcArea() {
	c.area = 3.14 * c.radius * c.radius
}

func main() {
	c := Circle{radius: 5}
	c.calcArea()
	fmt.Printf("%+v", c)
}

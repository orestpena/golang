package main

import "fmt"

type Student struct {
	name   string
	rollNo int
}

func main() {
	st := Student{"Joe", 12}
	fmt.Printf("%+v", st)
}

package main

import "fmt"

type Student struct {
	name   string
	rollNo int
}

func main() {
	st := Student{
		name:   "Joe",
		rollNo: 12,
	}
	fmt.Printf("%+v", st)
}

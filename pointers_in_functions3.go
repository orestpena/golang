package main

import (
	"fmt"
	"strings"
)

func modify(s *string) {
	*s = strings.ToUpper(*s)
}

func main() {
	s := "hello"
	fmt.Println(s)
	modify(&s)
	fmt.Println(s)
}

// this program is an error, you can't pass [3]int as type *[3]int

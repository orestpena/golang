package main

import "fmt"

func modify(s map[string]int) {
	s["A"] = 100
}

func main() {
	ascii_codes := map[string]int{}
	ascii_codes["A"] = 65
	fmt.Println(ascii_codes)
	modify(ascii_codes)
	fmt.Println(ascii_codes)
}

// this program is an error, you can't pass [3]int as type *[3]int

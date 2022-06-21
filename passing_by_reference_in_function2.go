package main

import "fmt"

// slices and maps are passed by reference by default.
// map example

func modify(m map[string]int) {
	m["K"] = 75

}

func main() {
	ascii_codes := make(map[string]int)
	ascii_codes["A"] = 65
	ascii_codes["F"] = 70
	fmt.Println(ascii_codes)
	modify(ascii_codes)
	fmt.Println(ascii_codes)
}

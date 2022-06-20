package main

import (
	"fmt"
	"strings"
)

func printStrings(names ...string) (names_c []string) {
	//	names_c := []string{}
	for _, value := range names {
		names_c = append(names_c, strings.ToUpper(value))
	}
	return
}

func main() {
	result := printStrings("Joe", "Monica", "Gunther")
	fmt.Println(result)
}

//there is an error in this program, the error is there is no new variables on the left sideof :=   (line #9)
// overwrites the original names of names_c with nothing, and so it fails. to fix it just comment out line #9

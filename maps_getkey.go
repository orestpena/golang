package main

import "fmt"

func main() {

	codes := map[string]int{"en": 1, "fr": 2, "hi": 3}
	value, found := codes["en"]
	fmt.Println(found, value)
	value, found = codes["hh"]
	fmt.Println(found, value)

	//the first one, it is true and it shows the value of the key
	//the second is false, and since 0 is the default value of int, it just gives the value 0 since it doesn't exist.
}

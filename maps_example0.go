package main

import "fmt"

func main() {
	// this is an error b/c the map is not closed {}
	//var ascii_codes map[string]int

	//this is an error, for map variable being added outside incorrectly
	//var ascii_codes map[string]string{}
	//ascii_codes["A"] = 65
	//fmt.Println(ascii_codes)

	ascii_codes := map[string]int{}
	ascii_codes["A"] = 65
	_, found := ascii_codes["B"]
	if found {
		fmt.Println(("key B was not found"))
	} //this will produce no output b/c there is no B, it will just execute
}

package main

import "fmt"

// this program is to truncate a map.
// truncating means clearing all elements of a map
// truncate method 1 is to iterate over the map
func main() {

	codes := map[string]string{"en": "English", "fr": "French", "hi": "Hindi"}
	// using key and value will cause an error b/c value is not being used. So only yes key.
	//for key, value := range codes {
	for key := range codes {
		delete(codes, key)
	}
	fmt.Println(codes)
	//will produce a blank map

}

package main

import "fmt"

// this program is to truncate a map.
// truncating means clearing all elements of a map
// truncate method 2 is to iterate over the map
func main() {

	codes := map[string]string{"en": "English", "fr": "French", "hi": "Hindi"}

	codes = make(map[string]string)

	fmt.Println(codes)
	// The second method to truncate just overwrites the previous map to clear it.
}

package main

import "fmt"

func main() {
	// maps in go are implemented by hash tables equivalent to:
	// python dictionaries.
	// php associated arrays
	// java hash tables

	codes := map[string]string{"en": "English",
		"fr": "French"}
	fmt.Println(codes)

}

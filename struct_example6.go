package main

import "fmt"

type Movie struct {
	name   string
	rating float32
}

func getMovie(s string, r float32) (m Movie) {
	m = Movie{
		name:   s,
		rating: r,
	}
	return
}

func main() {
	mov := getMovie("xyz", 2.1)
	mov1 := getMovie("abc", 3.3)
	if mov.rating == mov1.rating || mov != mov1 {
		fmt.Println("condition met")
	} else if mov.rating == mov1.rating {
		fmt.Println("condition_2 met")
	}
}

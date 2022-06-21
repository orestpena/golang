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

func increaseRating(m *Movie) {
	m.rating += 1.0
}

func main() {
	mov := getMovie{name: "ABCD"}
	increaseRating(mov)
	fmt.Printf("%+v", mov)
}

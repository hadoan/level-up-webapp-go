package main

import "fmt"

func main() {

	episodeIV := Movie{
		Title:       "Star Wars: A New Hope",
		Rating:      5.0,
		ReleaseYear: 1977,
	}
	episodeIV.Actors = []string{
		"Mark Hamill",
		"Harrison Ford",
		"Carrie Fisher",
	}

	fmt.Println(episodeIV.Title, "has a rating of", episodeIV.Rating)
	fmt.Println(episodeIV.DisplayTitle())
}

type Movie struct {
	Actors      []string
	Rating      float32
	ReleaseYear int
	Title       string
}

func (movie Movie) DisplayTitle() string {
	return fmt.Sprintf("%s (%d)", movie.Title, movie.ReleaseYear)
}

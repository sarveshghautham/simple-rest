package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Entertainment is the struct to hold tv shows
type Entertainment struct {
	TVShows []TVShow `json:"tvShows"`
}

// TVShow struct holds the name and genre
type TVShow struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
}

func getMyShows(w http.ResponseWriter, r *http.Request) {
	entertainment := Entertainment{
		TVShows: []TVShow{
			{ID: 1, Name: "Friends", Genre: "Comedy"},
			{ID: 2, Name: "The Man in the High Castle", Genre: "Mystery"},
		},
	}

	json.NewEncoder(w).Encode(entertainment)
}

func handleRequests() {
	http.HandleFunc("/", getMyShows)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// TVShow struct holds the name and genre
type TVShow struct {
	Name  string `json:"name"`
	Genre string `json:"genre"`
}

func getMyShows(w http.ResponseWriter, r *http.Request) {
	tvShows := []TVShow{
		{Name: "Friends", Genre: "Comedy"},
		{Name: "The Man in the High Castle", Genre: "Mystery"},
	}

	json.NewEncoder(w).Encode(tvShows)
}

func handleRequests() {
	http.HandleFunc("/", getMyShows)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}

package handlers

import (
	"fmt"
	"goclass/helpers"
	"log"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Pong"))
}

func GetMovieGenreHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bytes, err := helpers.Fetch("/genre/movie/list")
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(200)
	w.Write(bytes)
}

func GetTVGenreHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bytes, err := helpers.Fetch("/genre/tv/list")
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(200)
	w.Write(bytes)
}

func GetMovieDetailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()
	id := query.Get("id")
	invalidID := len(id) <= 0
	if invalidID {
		w.WriteHeader(400)
		w.Write([]byte("you must choose a movie id"))
		return
	}
	bytes, err := helpers.Fetch(fmt.Sprintf("/movie/%s", id))
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(200)
	w.Write(bytes)
}

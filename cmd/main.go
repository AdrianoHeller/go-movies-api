package main

import (
	"fmt"
	"goclass/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.PingHandler)
	mux.HandleFunc("/movies", handlers.GetMovieGenreHandler)
	mux.HandleFunc("/movie", handlers.GetMovieDetailHandler)
	mux.HandleFunc("/tv", handlers.GetTVGenreHandler)
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

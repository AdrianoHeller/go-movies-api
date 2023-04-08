package main

import (
	"encoding/json"
	"goclass/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPing(t *testing.T) {
	t.Run("GET 200 Pong", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handlers.PingHandler)
		handler.ServeHTTP(rr, req)
		requestOk := rr.Code == http.StatusOK
		if !requestOk {
			t.Errorf("got %v want %v", rr.Code, http.StatusOK)
		}

		got := rr.Body.String()
		want := "Pong"
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestGetMovieGenre(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/movies", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetMovieGenreHandler)
	handler.ServeHTTP(rr, req)
	requestOk := rr.Code == http.StatusOK
	if !requestOk {
		t.Errorf("got %v want %v", rr.Code, http.StatusOK)
	}

	got := rr.Body.Bytes()

	type Genre struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	type Genres struct {
		Genres []Genre `json:"genres"`
	}

	var genres Genres

	err = json.Unmarshal(got, &genres)
	if err != nil {
		t.Fatal(err)
	}

	if len(genres.Genres) == 0 {
		t.Errorf("Expected len > 0, but got == 0")
	}

}

func TestGetTVGenre(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/tv", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetTVGenreHandler)
	handler.ServeHTTP(rr, req)
	requestOk := rr.Code == http.StatusOK
	if !requestOk {
		t.Errorf("got %v want %v", rr.Code, http.StatusOK)
	}

	got := rr.Body.Bytes()

	type Genre struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	type Genres struct {
		Genres []Genre `json:"genres"`
	}

	var genres Genres

	err = json.Unmarshal(got, &genres)
	if err != nil {
		t.Fatal(err)
	}

	if len(genres.Genres) == 0 {
		t.Errorf("Expected len > 0, but got == 0")
	}

}

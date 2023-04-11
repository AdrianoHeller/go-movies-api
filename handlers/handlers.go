package handlers

import (
	"fmt"
	"goclass/helpers"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	helpers.SetJSON(w)
	helpers.CheckValidMethod(w, r, http.MethodGet)
	helpers.WriteResponse(w, http.StatusOK, []byte("Pong"))
}

func GetMovieGenreHandler(w http.ResponseWriter, r *http.Request) {
	helpers.SetJSON(w)
	helpers.CheckValidMethod(w, r, http.MethodGet)
	bytes, err := helpers.Fetch("/genre/movie/list")
	if err != nil {
		helpers.WriteResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		return
	}
	helpers.WriteResponse(w, http.StatusOK, bytes)
}

func GetTVGenreHandler(w http.ResponseWriter, r *http.Request) {
	helpers.SetJSON(w)
	helpers.CheckValidMethod(w, r, http.MethodGet)
	bytes, err := helpers.Fetch("/genre/tv/list")
	if err != nil {
		helpers.WriteResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		return
	}
	helpers.WriteResponse(w, http.StatusOK, bytes)
}

func GetMovieDetailHandler(w http.ResponseWriter, r *http.Request) {
	helpers.SetJSON(w)
	helpers.CheckValidMethod(w, r, http.MethodGet)
	query := r.URL.Query()
	id := query.Get("id")
	invalidID := len(id) <= 0
	if invalidID {
		helpers.WriteResponse(w, http.StatusBadRequest, []byte("you must choose a movie id"))
		return
	}
	bytes, err := helpers.Fetch(fmt.Sprintf("/movie/%s", id))
	if err != nil {
		helpers.WriteResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		return
	}
	helpers.WriteResponse(w, http.StatusOK, bytes)
}

func GetPersonHandler(w http.ResponseWriter, r *http.Request) {
	helpers.SetJSON(w)
	helpers.CheckValidMethod(w, r, http.MethodGet)
	query := r.URL.Query()
	id := query.Get("id")
	invalidID := len(id) <= 0
	if invalidID {
		helpers.WriteResponse(w, http.StatusBadRequest, []byte("you must choose a person id"))
		return
	}
	bytes, err := helpers.Fetch(fmt.Sprintf("/person/%s", id))
	if err != nil {
		helpers.WriteResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		return
	}
	helpers.WriteResponse(w, http.StatusOK, bytes)
}

func GetPersonImagesHandler(w http.ResponseWriter, r *http.Request) {
	helpers.SetJSON(w)
	helpers.CheckValidMethod(w, r, http.MethodGet)
	query := r.URL.Query()
	id := query.Get("id")
	invalidID := len(id) <= 0
	if invalidID {
		helpers.WriteResponse(w, http.StatusBadRequest, []byte("you must choose a person image id"))
		return
	}
	bytes, err := helpers.Fetch(fmt.Sprintf("/person/%s/images", id))
	if err != nil {
		helpers.WriteResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		return
	}
	helpers.WriteResponse(w, http.StatusOK, bytes)
}

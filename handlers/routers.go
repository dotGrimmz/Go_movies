package handlers 

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewServer() *http.Server {
	router := mux.NewRouter()
	router.HandleFunc("/movie", PostNewMovie).Methods("POST")
	router.HandleFunc("/movie/{id}", FindMovieById).Methods("GET")
	router.HandleFunc("/movie/{id}", DeleteMovieById).Methods("DELETE")
	router.HandleFunc("/movie", UpdateMovieById).Methods("PUT")
	server := &http.Server{
		Handler: router,
		Addr: "127.0.0.1:8080",
	}

	return server
}
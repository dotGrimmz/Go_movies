package handlers 

import (
	// "encoding/json"
	"net/http"
	// "movies_api/entities"
	// "movies_api/repo"
	"github.com/gorilla/mux"
	"movies_api/service"
//"fmt"
)


func PostNewMovie(res http.ResponseWriter, req *http.Request) {
	 service.CreateMovie(res , req)
}

func FindMovieById(res http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	service.FindMovieId(res, req, id)
}


func DeleteMovieById(res http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	service.DeleteMovie(res, req, id)
}


func  UpdateMovieById(res http.ResponseWriter, req *http.Request) {
	service.UpdateMovie(res, req)
}

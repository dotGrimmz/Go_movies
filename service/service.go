package service

import (
	"encoding/json"
	"net/http"
	"movies_api/entities"
	"movies_api/repo"
)


func CreateMovie(res http.ResponseWriter, req *http.Request) ( http.ResponseWriter,  *http.Request){
	movieStruct := entities.Movie{}
	repo := repo.Repo{}

	err := json.NewDecoder(req.Body).Decode(&movieStruct);
	movieStruct.SetId()
	valid := movieStruct.ValidateMovie()
	if valid != "valid" {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(valid))
		return res, req
	}
	movie, err := repo.AddMovie(movieStruct)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return res, req
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	res.Write(movie)

	return res, req
}

func FindMovieId(res http.ResponseWriter, req *http.Request,  id string)( http.ResponseWriter,  *http.Request) {
	repo := repo.Repo{}
	movie := repo.FindMovieById(id)
	valid := movie.ValidateMovie()
	if valid != "valid" {
		res.WriteHeader(http.StatusNoContent)
		res.Write([]byte(valid))
		return res, req
	}
	movieBytes, err := json.Marshal(movie)
	if err != nil {
		panic(err)
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	res.Write(movieBytes)

	return res, req
}

func DeleteMovie(res http.ResponseWriter, req *http.Request,  id string)( http.ResponseWriter,  *http.Request) {
	repo := repo.Repo{}
	msg, sucess := repo.DeleteMovie(id)

	if !sucess {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(msg))
		return res, req
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(msg))
	return res, req

}

func UpdateMovie(res http.ResponseWriter, req *http.Request) ( http.ResponseWriter,  *http.Request) {
	repo := repo.Repo{}
	movieStruct := entities.Movie{}
	err := json.NewDecoder(req.Body).Decode(&movieStruct);
	if err != nil {
		panic(err)
	}

	valid := movieStruct.ValidateMovie()
	if valid != "valid" {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(valid))
		return res, req
	}

	movie := repo.UpdateMovieById(movieStruct)
	movieBytes, err := json.Marshal(movie)
	if err != nil {
		panic(err)
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(movieBytes)
	return res, req
}
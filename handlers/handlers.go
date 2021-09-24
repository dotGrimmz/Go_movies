package handlers 

import (
	"encoding/json"
	"net/http"
	"movies_api/entities"
	"movies_api/repo"
)


func PostNewMovie(res http.ResponseWriter, req *http.Request) {
	movieStruct := entities.Movie{}
	repo := repo.Repo{}

	err := json.NewDecoder(req.Body).Decode(&movieStruct);
	movieStruct.SetId()
	err = repo.AddMovie(movieStruct)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

}	
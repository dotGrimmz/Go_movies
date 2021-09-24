package repo

import (
	"movies_api/entities"
	"encoding/json"
	"io/ioutil"
)
type Repo struct {
	Movies []entities.Movie
}



func (r Repo) AddMovie (film entities.Movie)  error {
		repo, _ := getAllMovies()
		repo.Movies = append(repo.Movies, film)
		jsonFile, err := json.MarshalIndent(repo, "", " ")
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile("moviedb.json", jsonFile, 0644)
		if err != nil {
			panic(err)
		}
		return nil
	}


func getAllMovies() (Repo, error) {
	file, _ := ioutil.ReadFile("moviedb.json")
	data := Repo{}
	err := json.Unmarshal([]byte(file), &data)
	if err != nil {
		panic(err)
	}
	return data, err
}
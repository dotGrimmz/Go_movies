package repo

import (
	"movies_api/entities"
	"encoding/json"
	"io/ioutil"
	"fmt"
)
type Repo struct {
	Movies []entities.Movie
	MoviesMap map[string] entities.Movie
}



func (r Repo) AddMovie(film entities.Movie) ([]byte, error) {
		repo, _ := getAllMovies()
		repo.Movies = append(repo.Movies, film)
		moveFile, err := json.Marshal(film)
		jsonFile, err := json.MarshalIndent(repo, "", " ")
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile("moviedb.json", jsonFile, 0644)
		if err != nil {
			panic(err)
		}
		return moveFile, err
	}


func getAllMovies() (Repo, error) {
	file, _ := ioutil.ReadFile("moviedb.json")
	data := Repo{}
	moviemap := make(map[string] entities.Movie)
	err := json.Unmarshal([]byte(file), &data)
	if err != nil {
		panic(err)
	}

	for _, movieStruct := range data.Movies {
		moviemap[movieStruct.Id] = movieStruct
	}
	data.MoviesMap = moviemap

	return data, err
}


func (r Repo) FindMovieById(Id string) entities.Movie {
	repo, _ := getAllMovies()
	movie := repo.MoviesMap[Id]
	
	return movie

}


func (r Repo) DeleteMovie(Id string) (string, bool) {
	repo, _ := getAllMovies()
	movieToBeDeleted := repo.MoviesMap[Id]
	movieToBeDeleted.ValidateMovie()
	if movieToBeDeleted.ValidateMovie() != "valid" {
		return "movie does not exist, It my already be deleted", false
	}
	delete(repo.MoviesMap, Id)
	newMovieList := []entities.Movie{}
	for _, m := range repo.MoviesMap {
		newMovieList = append(newMovieList, m)
	}
	repo.Movies = newMovieList
	jsonFile, err := json.MarshalIndent(repo, "", " ")
		if err != nil {
			panic(err)
		}
	err = ioutil.WriteFile("moviedb.json", jsonFile, 0644)
		if err != nil {
			panic(err)
		}
		return "Movie has been deleted", true
}



func (r Repo) UpdateMovieById(film entities.Movie) entities.Movie {
	repo, _ := getAllMovies()
	movie := repo.MoviesMap[film.Id]
	for _, m := range repo.Movies {
		if(movie.Id == m.Id) {
			// p := &movie
			// *p = film
			fmt.Print(m)
		}
	}


	// I need to overwrite the exixting data... I just dont want to do it 
	// in a for range loop. or what other instance is theree?


	 	jsonFile, err := json.MarshalIndent(repo, "", " ")
		if err != nil {
 			panic(err)
 		}
 		err = ioutil.WriteFile("moviedb.json", jsonFile, 0644)
 		if err != nil {
 			panic(err)
 		}


	return movie
}
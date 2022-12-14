package main

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	title string `json:"ttle"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies[]Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item:= range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies [index+1:]...)
			break
		}
	}
}


func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _. item := range maovies {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func main(){
	r := mux.NewRouter()


	movies = append(movies, Movie{ID: "1", Isbn:"43827", Title:"Movie One", Director : &Director{Firstname:"Bob", Lastname:"Big"}})
	movies = append(movies, Movie{ID: "2", Isbn:"93828", Title:"Movie Two", Director : &Director{Firstname:"Jane", Lastname:"Ray"}})
	
	r.HandleFunc("/movies", getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
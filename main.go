package main

import (
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

func main(){
	r := mux.NewRouter()


	//movies = append(movies, Movie{ID: "1", Isbn:"43827", Title:"Movie One", Director : &Director{Firstname:"Bob", Lastname:"Big"}})
	
	
	r.HandleFunc("/movies", getMovie).Method("GET")
	r.HandleFunc("/movies/{id}", getMovie).Method("GET")
	r.HandleFunc("/movies", createMovie).Method("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Method("PUT")
	r.HandleFunc("/movies{id}",deleteMovie).Method("DELETE")

	fmt.Printf("starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000, r"))

}
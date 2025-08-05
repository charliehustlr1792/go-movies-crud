package main
import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"encoding/json"
	"strconv"
	"net/http"
	"math/rand"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

var movies []Movie


func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"]{
			movies=append(movies[:index], movies[index+1:]...)
			break;
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r:=mux.NewRouter()
	movies=append(movies,Movie{ID: "1", Isbn:"4321234",Title:"Fight Club", Director:&Director{FirstName: "David", LastName: "Fincher"}})
	movies=append(movies,Movie{ID: "2", Isbn: "1234567",Title:"The Godfather", Director:&Director{FirstName: "Francis", LastName: "Coppola"}})
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
    r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("movies/{id}",deleteMovie).Methods("DELETE")


	fmt.Printf("Starting server at port 8000 \n")
	log.Fatal(http.ListenAndServer((":8000"),r))

}
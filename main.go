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

func deleteMovie(w http.ResponseWriter,r *http.Request){
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

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for _,item :=range movies {
		if item.ID ==params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}


func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	_= json.NewDecoder(r.Body).Decode(&movie)//taking the json data from the request body and decoding it into the movie variable; _= part means ignoring the error returned by Decode
	movie.ID=strconv.Itoa(rand.Intn(1000000))//generates a random movie id rand.Intn(1000000) gives a random number between 0 to 999999 strconv.Itoa converts that into a string
	movies=append(movies, movie)//adds the movie
	json.NewEncoder(w).Encode(movie)//sends the json response
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json") //set json content type
	params:=mux.Vars(r) //get the parameters from the request
	//loop over the movies range
	//delete the movie with the id that you have sent
	//add a new movie - in the body that we send
	for index,item := range movies {
		if item.ID==params["id"]{
			movies=append(movies[:index], movies[index+1:]...)
			var movie Movie
			_=json.NewDecoder(r.Body).Decode(&movie)
			movie.ID=params["id"]//we are enduring the id in the url takes priority not the one in the request body
			movies=append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return 
		}
	}
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
	log.Fatal(http.ListenAndServe(":8000", r))

}
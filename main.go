package main

import(
    "fmt"               // printing
    "log"               // logging errors
    "encoding/json"     // encode data into json when sending to postman
    "math/rand"         // create random id
    "net/http"          // allows us to create a server in go
    "strconv"           // convert integer id to a string
    "github.com/gorilla/mux"
)

type Movie struct {
    ID string `json:"id"`
    Isbn string `json:"isbn"`
    Title string `json:"title"`
    Director *Director `json:"director"`
}

type Director struct {
    Firstname string `json:"firstname"`
    Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for _, item := range movies {
        if(item.ID == params["id"]){
            json.NewEncoder(w).Encode(item)
        }
    }
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range movies {
        if(item.ID == params["id"]){
            movies = append(movies[:index], movies[index + 1:]...)
            break
        }
    }
}

func main() {
    r := mux.NewRouter()

    movies = append(movies, Movie{ ID: "1", Isbn: "438227", Title: "Movie1", Director: &Director{ Firstname: "Ryan", Lastname: "Mello" }})
    movies = append(movies, Movie{ ID: "2", Isbn: "674291", Title: "Movie2", Director: &Director{ Firstname: "Josh", Lastname: "Uhler" }})
    movies = append(movies, Movie{ ID: "3", Isbn: "859350", Title: "Movie3", Director: &Director{ Firstname: "Parker", Lastname: "Nelson" }})
    movies = append(movies, Movie{ ID: "4", Isbn: "713529", Title: "Movie4", Director: &Director{ Firstname: "Brandon", Lastname: "Chamizo" }})

    r.HandleFunc("/movies", getMovies).Methods("GET")
    r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
    r.HandleFunc("/movies", createMovie).Methods("POST")
    r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
    r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

    fmt.Printf("Starting server at port 8000\n")
    log.Fatal(http.ListenAndServe(":8000", r))
}
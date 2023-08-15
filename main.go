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

// every movie has a director
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

func main() {
    
}
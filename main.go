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

}

type Director struct {

}

func main() {
    
}
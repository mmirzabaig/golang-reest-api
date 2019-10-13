package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// book struct (Model)
type Book struct {
	ID     string  `json: "id"`
	Isbn   string  `json: "isbn"`
	Title  string  `json: "title"`
	Author *Author `json: "id"`
}

// author struct
type Author struct {
	Firstname string `json: "firstname`
	Lastname  string `json: "lastname`
}

// get all books
func getBooks(w http.ResponseWriter, r *http.Request) {

}

func main() {
	fmt.Println("Hello World!")
	// init Router
	router := mux.NewRouter()

	// route handlers
	router.HandleFunc("/api/books", getBooks()).Methods("GET")
}

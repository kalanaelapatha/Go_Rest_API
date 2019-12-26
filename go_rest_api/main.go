package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//we use json request and responce and http protocols so that we must need to import those kind of packages to that projects
//we use to math lib to generate a random number when we use the Id
//and we need to use String Convert to convert the json to string

//Books Struct(Model)
type Books struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Tilte  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var books []Books //use that to mock data first -> 1 st Step

// Get All books

func getBooks(w http.ResponseWriter, route *http.Request) {

	w.Header().Set("Content-Type", "application/json") // third Step
	json.NewEncoder(w).Encode(books)                   // third Step

}

//Get Single Book by id
func getBook(w http.ResponseWriter, route *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(route) //get the params

	//Loops through books and find with id

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Books{})
}

// Create a book
func createBook(w http.ResponseWriter, route *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Books
	_ = json.NewDecoder(route.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock Id - Not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

//update a book

//its the cobination of the delete and create function also
//without change anythind jsu copy and past all of things under delete
//then get part of thing from create method
//after return it
func updateBook(w http.ResponseWriter, route *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(route) //get all params

	for index, item := range books {

		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			// from create method
			var book Books
			_ = json.NewDecoder(route.Body).Decode(&book)
			book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock Id - Not safe
			books = append(books, book)
			//end from create method

			return
		}
	}
	json.NewEncoder(w).Encode(books)

}

//delete a book
func deleteBook(w http.ResponseWriter, route *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(route) //get all params

	for index, item := range books {

		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)

}

func main() {
	//init  Router

	route := mux.NewRouter()

	// Mock Data - @todo - implement DB

	books = append(books, Books{ID: "1", Isbn: "234234", Tilte: "X Files", Author: &Author{FirstName: "Steve", LastName: "Harvey"}}) // second step
	books = append(books, Books{ID: "2", Isbn: "456567", Tilte: "Chemist", Author: &Author{FirstName: "jhone", LastName: "Doe"}})    // second step
	// Route Handlers / End Points

	route.HandleFunc("/api/books", getBooks).Methods("GET")
	route.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	route.HandleFunc("/api/books", createBook).Methods("POST")
	route.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	route.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	//setup the server

	log.Fatal(http.ListenAndServe(":8000", route))

}

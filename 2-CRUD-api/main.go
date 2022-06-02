package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
}

var books []Book

func main() {
	r := mux.NewRouter()

	// Dummy books
	books = append(books, Book{ID: "1", Title: "Butterflies",
		Author: &Author{FirstName: "Wra", LastName: "Lith"}})

	books = append(books, Book{ID: "2", Title: "Hurricanes",
		Author: &Author{FirstName: "John", LastName: "No"}})

	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{ID}", getBookById).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{ID}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{ID}", deleteBook).Methods("DELETE")

	fmt.Printf("Starting server at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Req params
	for _, item := range books {
		if item.ID == params["ID"] { // If id exist
			json.NewEncoder(w).Encode(item) // Write encoded JSON
			return
		}
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book                             // Create new var on Book struct
	_ = json.NewDecoder(r.Body).Decode(&book) // Appoint body of JSON request to new variable
	book.ID = strconv.Itoa(rand.Intn(100000)) // Random ID
	books = append(books, book)               // Ad created book to existing books
	json.NewEncoder(w).Encode(book)           // Write encoded JSON

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["ID"] {
			// Delete Book - Not DRY :(
			books = append(books[:index], books[index+1:]...)
			// Create Book - Not DRY :(
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["ID"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)            // Params of the request
	for index, item := range books { // Loop all books O(n)
		if item.ID == params["ID"] { // If id in params exist on books
			books = append(books[:index], books[index+1:]...) // Remove book
			break
		}
	}
	json.NewEncoder(w).Encode(books)

}

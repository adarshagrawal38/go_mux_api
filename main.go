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
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title  string `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(books)

}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params)
	for _, item := range books {
		if item.ID ==  params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = strconv.Itoa(rand.Intn(10000000))
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	r := mux.NewRouter()

	books = append(books, Book{ID: "123", Isbn: "12345", Title: "Book one", Author: &Author{FirstName: "Adatsh", LastName: "Agrawal"}})
	books = append(books, Book{ID: "124", Isbn: "23234324", Title: "Book two", Author: &Author{FirstName: "Rajesh", LastName: "Agrawal"}})

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
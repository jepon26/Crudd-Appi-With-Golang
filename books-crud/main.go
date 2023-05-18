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
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var books []Book

const (
	contentTypeJSON      = "application/json"
	headerContentTypeKey = "Content-Type"
)

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(headerContentTypeKey, contentTypeJSON)
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(headerContentTypeKey, contentTypeJSON)
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(headerContentTypeKey, contentTypeJSON)
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(headerContentTypeKey, contentTypeJSON)
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(headerContentTypeKey, contentTypeJSON)
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	books = append(books, Book{ID: "1", Isbn: "456789", Title: "Amazing Book", Author: &Author{FirstName: "Elvin", LastName: "Rueda"}})
	books = append(books, Book{ID: "2", Isbn: "123456", Title: "Real Book", Author: &Author{FirstName: "Javier", LastName: "Gomez"}})
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	fmt.Println("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}

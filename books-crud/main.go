package main


import (
	"fmt"
	"log"
	"encoding.json"
	"math.rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)	


type Book struct {
    ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}


	type Author struct {
	firstName string `json:"firstname"`
	lastName string `json:"lastname"`
}


var books[]Book


func getBooks(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(books)
}


func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"]{
			books = append(books[:index], books[index + 1:]...)
			break
		}
	}
}


func main(){
r := mux.NewRouter()



books = append(books, Book{ID: "1", Isbn: "456789", Title: "Amazing Book", Author: &Author{firstName: "Elvin", lastName: "Rueda"}})
books = append(books, Book{ID: "2", Isbn: "123456", Title: "Real Book", Author: &Author{firstName: "Javier", lastName: "Gomez"}})
r.HandleFunc("/books", getBooks).Methods("GET")
r.HandleFunc("/books/{id}", getBook).Methods("GET")
r.HandleFunc("/books", createBook).Methods("POST")
r.HandleFunc("/books/{id}", updatedBook).Methods("PUT")
r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

fmt.Println("Starting server at port 8080\n")
log.Fatal(http.ListenAndServe(":8080", r))
}


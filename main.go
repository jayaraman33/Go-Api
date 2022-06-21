package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// book struct

type Book struct {
	ID     string `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// type Todo struct {
// 	Title    string `json:"title"`
// 	SubTitle string `json:"subtitle"`
// 	Isdone   bool   `json:'"isdone"`
// }

var books []Book

// get all books

func getBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	json.NewEncoder(w).Encode(books)
}

// get one book

func getBook(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// add new book

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// update book

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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

// delete book


func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
	router := mux.NewRouter()

	books = append(books, Book{ID: "1", Isbn: "96880", Title: "Book One", Author: "Ram"})
	books = append(books, Book{ID: "2", Isbn: "82936", Title: "Book Two", Author: "Jayaraman"})
	books = append(books, Book{ID: "3", Isbn: "12345", Title: "Book Three", Author: "Vishnu"})
	books = append(books, Book{ID: "4", Isbn: "67890", Title: "Book Four", Author: "Raman"})

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("Post")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("Put")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}

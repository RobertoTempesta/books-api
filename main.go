package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Estrutura do Book
type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   int    `json:year`
}

// Base de dados em memória
var books = []Book{
	{ID: 1, Title: "12 Regras para a vida", Author: "Jordam Peterson", Year: 2017},
	{ID: 2, Title: "O Hobbit", Author: "J.R.R. Tolkien", Year: 1960},
	{ID: 3, Title: "1984", Author: "George Orwell", Year: 1949},
}

// GET todos os livros
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get single book by ID
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID Ivalido", http.StatusBadRequest)
		return
	}
	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

func main() {
	// Inicializar router
	router := mux.NewRouter()

	// Endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", router))
}

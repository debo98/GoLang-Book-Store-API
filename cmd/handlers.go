package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// setJSONContentType sets the Content-Type header to application/json.
func setJSONContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

// getBooksHandler handles GET requests to list all books.
func getBooksHandler(w http.ResponseWriter, _ *http.Request) {
	setJSONContentType(w)
	mu.Lock()
	defer mu.Unlock()
	json.NewEncoder(w).Encode(books)
}

// createBookHandler handles POST requests to add a new book.
func createBookHandler(w http.ResponseWriter, r *http.Request) {
	setJSONContentType(w)
	var newBook Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	mu.Lock()
	nextID++
	newBook.ID = nextID
	books = append(books, newBook)
	mu.Unlock()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

// deleteBookHandler handles DELETE requests to remove a book by ID.
func deleteBookHandler(w http.ResponseWriter, r *http.Request) {
	setJSONContentType(w)
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Missing book id in URL", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid book id", http.StatusBadRequest)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	index := -1
	for i, b := range books {
		if b.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	books = append(books[:index], books[index+1:]...)  // remove the book from the slice of books by creating a new slice
	w.WriteHeader(http.StatusNoContent)
}

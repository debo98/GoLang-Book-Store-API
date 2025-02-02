package main

import (
	"log"
	"net/http"
)

func main() {
	// Route for handling GET (view all books) and POST (add a new book)
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getBooksHandler(w, r)
		case http.MethodPost:
			createBookHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Route for handling DELETE requests for a specific book.
	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			deleteBookHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

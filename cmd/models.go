package main

import "sync"

// Book represents a book in our system.
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var (
	books  []Book     // in-memory storage for books
	nextID int        // used to generate unique IDs
	mu     sync.Mutex // ensures safe concurrent access
)

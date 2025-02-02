# Go Book API

A simple RESTful API written in Go to manage a list of books. The API supports adding a new book, retrieving all books, and deleting a book by its ID.

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Running the Code](#running-the-code)
- [API Endpoints](#api-endpoints)
  - [1. Get All Books](#1-get-all-books)
  - [2. Add a Book](#2-add-a-book)
  - [3. Delete a Book](#3-delete-a-book)

## Features

- **Add a Book:** Create a new book by providing its title and author.
- **View All Books:** Retrieve a list of all books stored in memory.
- **Delete a Book:** Remove a book by its unique ID.

## Technologies Used

- [Go](https://golang.org/) (Golang)
- Standard library packages:
  - `net/http`
  - `encoding/json`
  - `strconv`
  - `strings`
  - `sync`
  - `log`

## Running the Code

1. **Clone the repository:**

    ```bash
    git clone https://github.com/debo98/GoLang-ToDo-List.git
    cd GoLang-ToDo-List/cmd/
    ```

2. **Running the server**

```bash
./bookapi
```

This command will compile and start the server on port 8000. You should see a log message like:

```bash
Server running on http://localhost:8000
```

## API Endpoints

### 1. Get All Books

cURL:

```bash
curl -X GET http://localhost:8000/books
```

Description: Returns a JSON array of all books stored in memory.

Response:

- 200 OK with a JSON array of book objects.

Example Response:

```bash
[
    {
        "id": 1,
        "title": "Book Title",
        "author": "Author Name"
    },
    {
        "id": 2,
        "title": "Another Book",
        "author": "Another Author"
    }
]
```

### 2. Add a Book

cURL:

```bash
curl -X POST http://localhost:8000/books \
-H "Content-Type: application/json" \
-d '{"title": "Book Title", "author": "Author Name"}'
```

Description: Accepts a JSON payload to create a new book. The book is added to an in-memory list with a unique ID.

Response:

- 201 Created with the newly created book object (including its unique ID).

Example Response:

```bash
{
    "id": 3,
    "title": "New Book Title",
    "author": "New Author"
}
```

Error Handling:

- 400 Bad Request: If the request body is invalid or missing required fields.

### 3. Delete a Book

cURL:

```bash
curl -X DELETE http://localhost:8000/books/1
```

Description: Deletes the book with the specified ID.

URL Parameter:

- id – The unique identifier of the book to delete.

Response:

- 204 No Content: Indicates successful deletion with no response body.

Error Handling:

- 400 Bad Request: If the book ID is missing or not a valid integer.
- 404 Not Found: If no book with the provided ID exists.

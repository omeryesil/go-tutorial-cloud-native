package api

import (
	"encoding/json"
	"net/http"
)

//Book model
type Book struct {
	Title       string `json:"title"` //pascal case in Code, and camel case in JSON
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description"`
}

var mockedBooksRepo = map[string]Book{
	"11": Book{Title: "Cloud Native Go", Author: "Omer Yesil", ISBN: "11", Description: "Desc1"},
	"22": Book{Title: "Cloud Native Java", Author: "Omer Yesil", ISBN: "22", Description: "Desc2"},
	"33": Book{Title: "Cloud Native C#", Author: "Omer Yesil", ISBN: "33", Description: "Desc3"},
	"44": Book{Title: "Cloud Native C++", Author: "Omer Yesil", ISBN: "44", Description: "Desc4"},
	"55": Book{Title: "Cloud Native Python", Author: "Omer Yesil", ISBN: "55", Description: "Desc5"},
}

//ToJSON is used dto convert Book to JSON
func (b Book) ToJSON() []byte {
	toJSON, err := json.Marshal(b)

	if err != nil {
		panic(err)
	}

	return toJSON
}

func fromJSON(data []byte) Book {
	book := Book{}

	err := json.Unmarshal(data, &book)

	if err != nil {
		panic(err)
	}

	return book
}

//BooksHandleFunc returns list of books
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {

	switch method := r.Method; method {
	case http.MethodGet:
		books := AllBooks()
		writeJSON(w, books)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method"))
	}

}

//BookHandleFunc returns specific book based on given ISBN number
func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		isbn := r.URL.Query()["isbn"][0]
		book := mockedBooksRepo[isbn]
		writeJSON(w, []Book{book})

	case http.MethodPost:

	case http.MethodDelete:
	}

	booksInJSON, err := json.Marshal(mockedBooksRepo)

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	w.Write(booksInJSON)
}

//AllBooks returns list of books
func AllBooks() []Book {
	books := []Book{}

	for _, b := range mockedBooksRepo {
		books = append(books, b)
	}

	return books
}

func writeJSON(w http.ResponseWriter, books []Book) {
	booksInJSON, err := json.Marshal(books)

	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(booksInJSON)
}

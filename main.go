package main

import (
	"fmt"
	"github.com/omeryesil/go-tutorial-cloud-native/api"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Listening on http://localhost" + port())

	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", apiEcho)

	http.HandleFunc("/api/books", api.BooksHandleFunc)
	http.HandleFunc("/api/books/", api.BookHandleFunc)

	err := http.ListenAndServe(port(), nil)

	if err != nil {
		log.Fatal("ERROR:", err)
		os.Exit(1)
	}
}

func port() string {
	port := os.Getenv("NATIVE_APP_PORT")

	if len(port) == 0 {
		return ":8080"
	}

	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Hello Cloud Native Go")
}

func apiEcho(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]

	w.Header().Add("Content-Type", "text/plain")

	fmt.Fprintf(w, message)
}

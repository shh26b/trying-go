package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBookPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", vars["title"], vars["page"])
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Books list")
}

func main() {
	r := mux.NewRouter()

	book := r.PathPrefix("/books").Subrouter()

	book.HandleFunc("/", GetBooks).Methods("GET")
	book.HandleFunc("/{title}/page/{page}", GetBookPage).Methods("GET")

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":5000", r)
}

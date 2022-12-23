package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shh26b/trying-go/http_server/todo"
)

type RootHandler struct {
	Msg string
}

func (h *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v \n", r.Method, r.URL)

	if r.URL.Path == "/todos" {
		todoHandler := &todo.TodoHandler{}
		todoHandler.ServeHTTP(w, r)
		return
	}

	fmt.Fprintf(w, h.Msg)
	fmt.Fprintf(w, "\nYou requested with uri: %s\n", r.URL)
}

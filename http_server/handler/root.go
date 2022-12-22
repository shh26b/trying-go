package handler

import (
	"fmt"
	"log"
	"net/http"
)

type RootHandler struct {
	Msg string
}

func (h *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v \n", r.Method, r.URL)

	fmt.Fprintf(w, h.Msg)
	fmt.Fprintf(w, "\nYou requested with uri: %s\n", r.URL)
}

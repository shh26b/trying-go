package utils

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func Fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Render(w *http.ResponseWriter, file string, data interface{}) {
	dir, err := os.Getwd()
	Fatal(err)

	dir = fmt.Sprintf("%v/tmpl/%v.tmpl", dir, file)

	t, err := template.ParseFiles(dir)
	Fatal(err)

	err = t.Execute(*w, data)
	Fatal(err)
}

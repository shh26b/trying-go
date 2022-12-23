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

// func todoToJson(t todo) string {
// 	return fmt.Sprintf("%v,", t.String())
// }

// func todosToJson(todos []todo) string {
// 	l := len(todos)
// 	s := "["
// 	for _, todo := range todos[:l-1] {
// 		s += todoToJson(todo)
// 	}
// 	s += fmt.Sprintf("\t%v]", todos[l-1].String())
// 	return s
// }

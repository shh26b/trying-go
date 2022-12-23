package todo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shh26b/trying-go/http_server/utils"
)

type TodoHandler struct {
	todos []todo
}

type Data struct {
	Todos []todo
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

type context struct {
	h *TodoHandler
	w *http.ResponseWriter
	r *http.Request
}

func listPost(ctx context) {
	utils.Render(ctx.w, "todo/list", Data{ctx.h.todos})
}

type td struct {
	Todo string
}

func createPost(ctx context) {
	var p td
	err := json.NewDecoder(ctx.r.Body).Decode(&p)
	if err != nil {
		fmt.Println(err)
		http.Error(*ctx.w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Print(p.Todo, "Here")
}

func (h *TodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.todos = append(h.todos, todo{"11", "Hello word", true})
	h.todos = append(h.todos, todo{"12", "Hello word 2", false})

	if r.Method == "POST" {
		createPost(context{h, &w, r})
	} else if r.Method == "GET" {
		listPost(context{h, &w, r})
	}
}

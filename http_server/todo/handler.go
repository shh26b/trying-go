package todo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/shh26b/go-code/http_server/utils"
)

type TodoHandler struct {
	Todos []Todo
}

type context struct {
	h *TodoHandler
	w *http.ResponseWriter
	r *http.Request
}

func listPost(ctx context) {
	utils.Render(ctx.w, "todo/list", TodoHandler{ctx.h.Todos})
}

func createPost(ctx context) {
	t := createTodo{}
	err := json.NewDecoder(ctx.r.Body).Decode(&t)
	if err != nil {
		log.Println(err)
		http.Error(*ctx.w, err.Error(), http.StatusBadRequest)
		return
	}
	err = t.validate()
	if err != nil {
		log.Println(err)
		http.Error(*ctx.w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(t.Todo)
	ctx.h.Todos = append(ctx.h.Todos, Todo{"1", t.Todo, true})
	fmt.Println(ctx.h.Todos)
}

func (h *TodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		createPost(context{h, &w, r})
	} else if r.Method == "GET" {
		listPost(context{h, &w, r})
	}
}

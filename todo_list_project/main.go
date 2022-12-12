package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/thedevsaddam/renderer"
	mgov2 "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/shh26b/trying-go/todo_list_project/todo"
	"github.com/shh26b/trying-go/todo_list_project/utils"
)

var rnd *renderer.Render
var db *mgov2.Database

func init() {
	rnd = renderer.New()
	sess, err := mgov2.Dial(utils.DbHostName)
	utils.CheckErr(err)
	sess.SetMode(mgov2.Monotonic, true)
	db = sess.DB(utils.DbName)
}

func todoHandlers() http.Handler {
	rg := chi.NewRouter()
	rg.Group(func(r chi.Router) {
		r.Get("/", fetchTodos)
		r.Post("/", createTodo)
		r.Put("/{id}", updateToto)
		r.Delete("/{id}", deleteTodo)
	})
	return rg
}

func trimId(r *http.Request) string {
	return strings.TrimSpace(chi.URLParam(r, "id"))
}

func updateToto(w http.ResponseWriter, r *http.Request) {
	id := trimId(r)
	if !bson.IsObjectIdHex(id) {
		rnd.JSON(w, http.StatusBadRequest, renderer.M{
			"message": "The id is invalid",
		})
		return
	}

	t := todo.Todo{}
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		rnd.JSON(w, http.StatusProcessing, err)
		return
	}

	if t.Title == "" {
		rnd.JSON(w, http.StatusBadRequest, renderer.M{
			"message": "The title is required",
		})
		return
	}

	if err := db.C(utils.CollectionName).Update(
		bson.M{"_id": bson.ObjectIdHex(id)},
		bson.M{"title": t.Title, "completed": t.Completed},
	); err != nil {
		rnd.JSON(w, http.StatusProcessing, renderer.M{
			"message": "Failed to update todo",
			"error":   err,
		})
		return
	}
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id := trimId(r)
	if !bson.IsObjectIdHex(id) {
		rnd.JSON(w, http.StatusBadRequest, renderer.M{
			"message": "The title is required",
		})
		return
	}

	if err := db.C(utils.CollectionName).RemoveId(bson.ObjectIdHex(id)); err != nil {
		rnd.JSON(w, http.StatusProcessing, renderer.M{
			"message": "Failed to delete todo",
		})
		return
	}
	rnd.JSON(w, http.StatusOK, renderer.M{
		"message": "Todo deleted successfully",
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := rnd.Template(w, http.StatusOK, []string{"static/home.tpl"}, nil)
	utils.CheckErr(err)
}

func fetchTodos(w http.ResponseWriter, r *http.Request) {
	todos := []todo.TodoModel{}

	if err := db.C(utils.CollectionName).Find(bson.M{}).All(&todos); err != nil {
		rnd.JSON(w, http.StatusProcessing, renderer.M{
			"message": "Failed to fetch todos",
			"error":   err,
		})
		return
	}
	todoList := []todo.Todo{}

	for _, t := range todos {
		todoList = append(todoList, todo.Todo{
			ID:        t.ID.Hex(),
			Title:     t.Title,
			Completed: t.Completed,
			CreateAt:  t.CreateAt,
		})
	}
	rnd.JSON(w, http.StatusOK, renderer.M{"data": todoList})
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	t := todo.Todo{}

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		rnd.JSON(w, http.StatusProcessing, err)
		return
	}

	if t.Title == "" {
		rnd.JSON(w, http.StatusBadRequest, renderer.M{
			"message": "The title is required",
		})
		return
	}

	tm := todo.TodoModel{
		ID:        bson.ObjectId(t.ID),
		Title:     t.Title,
		Completed: false,
		CreateAt:  time.Now(),
	}

	if err := db.C(utils.CollectionName).Insert(&tm); err != nil {
		rnd.JSON(w, http.StatusProcessing, renderer.M{
			"message": "Failed to save todo",
		})
		return
	}

	rnd.JSON(w, http.StatusCreated, renderer.M{
		"message": "Todo create successfully",
		"todo_id": tm.ID.Hex(),
	})
}

func main() {
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Mount("/todo", todoHandlers())

	srv := &http.Server{
		Addr:         utils.Port,
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	go func() {
		log.Println("Listening on port", utils.Port)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Listen:%s\n", err)
		}
	}()

	<-stopChan
	log.Println("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
	defer cancel()
	log.Println("Server gracefully stopped")
}

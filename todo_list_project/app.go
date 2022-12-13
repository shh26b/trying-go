package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shh26b/trying-go/todo_list_project/core"
	"github.com/shh26b/trying-go/todo_list_project/todo"
)

func HomeHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Hello World",
	})
}

func start() {
	app := gin.Default()

	core.ConnectMongo()

	app.GET("/", HomeHandler)
	todo.Init(app)

	if err := app.Run(core.Port); err != nil {
		log.Fatal(err)
	}
}

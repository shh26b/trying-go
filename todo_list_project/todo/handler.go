package todo

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shh26b/go-code/todo_list_project/core"
	"go.mongodb.org/mongo-driver/mongo"
)

var todoColl *mongo.Collection = core.GetColl(core.DB, "todo")

func getTodos(ctx *gin.Context) {
	c, cancelCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelCtx()

	x, err := todoColl.Find(c, nil)

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "todo",
	})
}

func Init(app *gin.Engine) {
	r := app.Group("/api/todos")

	r.GET("/", getTodos)
}

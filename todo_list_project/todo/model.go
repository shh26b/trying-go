package todo

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	TodoModel struct {
		ID        bson.ObjectId `bson:"_id,omitempty"`
		Title     string        `bson:"title"`
		Completed bool          `bson:"completed"`
		CreateAt  time.Time     `bson:"createAt"`
	}
	Todo struct {
		ID        string    `json:"id"`
		Title     string    `json:"title"`
		Completed bool      `json:"completed"`
		CreateAt  time.Time `json:"createAt"`
	}
)

package todo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title"`
	Completed bool               `bson:"completed"`
	CreateAt  time.Time          `bson:"createAt"`
}

type TodoResponse struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreateAt  time.Time `json:"createAt"`
}

// user := bson.M{"foo": "bar", "hello": "world", "pi": 3.14159}
// result, err := db.Collection("todo").InsertOne(
// 	context.TODO(),
// 	user,
// )
// checkE(err)
// fmt.Println(result.InsertedID)

// users := []interface{}{
// 	bson.M{"fullName": "User 2", "age": 25},
// 	bson.M{"fullName": "User 3", "age": 20},
// 	bson.M{"fullName": "User 4", "age": 28},
// }
// results, err := db.Collection("todo").InsertMany(context.TODO(), users)
// checkE(err)
// fmt.Println(results.InsertedIDs)

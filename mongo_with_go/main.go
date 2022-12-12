package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MongoURI = "mongodb://localhost:27017"

func checkE(err error) {
	if err != nil {
		log.Fatal("MongoDB error:", err)
	}
}

func main() {
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(MongoURI),
	)
	checkE(err)
	db := client.Database("with-go")

	user := bson.M{"foo": "bar", "hello": "world", "pi": 3.14159}
	result, err := db.Collection("todo").InsertOne(
		context.TODO(),
		user,
	)
	checkE(err)
	fmt.Println(result.InsertedID)

	users := []interface{}{
		bson.M{"fullName": "User 2", "age": 25},
		bson.M{"fullName": "User 3", "age": 20},
		bson.M{"fullName": "User 4", "age": 28},
	}
	results, err := db.Collection("todo").InsertMany(context.TODO(), users)
	checkE(err)
	fmt.Println(results.InsertedIDs)

}

package core

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func checkE(err error) {
	if err != nil {
		log.Fatal("MongoDB error:", err)
	}
}

func ConnectMongo() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoURI))
	checkE(err)

	ctx, cancelCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelCtx()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	return client.Database(DBName)
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
}

var DB *mongo.Database = ConnectMongo()

func GetColl(db *mongo.Database, cName string) *mongo.Collection {
	collection := db.Collection(cName)
	return collection
}

package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(dBName, mongoURI string) error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	db := client.Database(dBName)

	Mongo = mongoInstance{
		Client: client,
		DB:     db,
		Ctx:    ctx,
	}
	return nil
}

func Disconnect() {
	if err := Mongo.Client.Disconnect(Mongo.Ctx); err != nil {
		log.Fatal(err)
	}
}

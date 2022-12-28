package book

import "go.mongodb.org/mongo-driver/mongo"

type (
	msa = map[string]interface{}
)

var (
	bookColl *mongo.Collection
)

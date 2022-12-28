package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Mongo mongoInstance
)

type mongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
	Ctx    context.Context
}

type (
	msa   = map[string]interface{}
	ctxT  = context.Context
	collT = mongo.Collection
)

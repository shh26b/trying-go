package db

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getObjectID(id string) (_id primitive.ObjectID, err error) {
	return primitive.ObjectIDFromHex(id)
}

func CreateOne(ctx ctxT, coll *collT, data any) (status int, err error) {
	res, err := coll.InsertOne(ctx, data)
	if err != nil {
		return 500, err
	}
	res2 := coll.FindOne(ctx, &msa{"_id": res.InsertedID})
	err = res2.Decode(data)
	if err != nil {
		return 500, err
	}
	return 201, nil
}

func GetMany(ctx ctxT, coll *collT, data any, filter msa) (status int, err error) {
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return 500, err
	}
	err = cursor.All(ctx, data)
	if err != nil {
		return 500, err
	}
	return 200, nil
}

func GetOne(ctx ctxT, coll *collT, data any, id string) (status int, err error) {
	var _id primitive.ObjectID
	_id, err = getObjectID(id)
	if err != nil {
		return 400, err
	}
	res := coll.FindOne(ctx, &msa{"_id": _id})
	err = res.Err()
	if err != nil {
		status = 500
		if err == mongo.ErrNoDocuments {
			status = 404
		}
		return status, err
	}
	err = res.Decode(data)
	if err != nil {
		return 500, err
	}
	return 200, nil
}

func DeleteOne(ctx ctxT, coll *collT, id string) (status int, err error) {
	var _id primitive.ObjectID
	_id, err = getObjectID(id)
	if err != nil {
		return 400, err
	}
	res, err := coll.DeleteOne(ctx, &msa{"_id": _id})
	if err != nil {
		return 500, err
	}
	if res.DeletedCount < 1 {
		err := errors.New("Record can't found with id: " + id)
		return 404, err
	}
	return 204, nil
}

func UpdateOne(ctx ctxT, coll *collT, data any, id string) (status int, err error) {
	var _id primitive.ObjectID
	_id, err = getObjectID(id)
	if err != nil {
		return 400, err
	}
	res := coll.FindOneAndUpdate(
		ctx,
		msa{"_id": _id}, msa{"$set": data},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)
	err = res.Err()
	if err != nil {
		status = 500
		if err == mongo.ErrNoDocuments {
			status = 404
		}
		return status, err
	}
	err = res.Decode(data)
	if err != nil {
		return 500, err
	}
	return 200, nil
}

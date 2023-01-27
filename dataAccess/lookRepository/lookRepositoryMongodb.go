package lookRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LookRepositoryMongodb struct {
	looksCollection *mongo.Collection
}

func NewLookRepositoryMongodb(database *mongo.Database) *LookRepositoryMongodb {
	return &LookRepositoryMongodb{
		looksCollection: database.Collection("looks"),
	}
}

func (lookRepositoryMongodb *LookRepositoryMongodb) FindLooks() []Look {
	var looks []Look

	cur, err := lookRepositoryMongodb.looksCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	cur.Decode(looks)

	return looks
}

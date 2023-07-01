package lookRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LookBson struct {
	_id    string `bson:"_id,omitempty"`
	Colour string `bson:"colour,omitempty"`
	Brand  string `bson:"brand,omitempty"`
}

type LookRepositoryMongodb struct {
	looksCollection *mongo.Collection
}

func NewLookRepositoryMongodb(database *mongo.Database) *LookRepositoryMongodb {
	return &LookRepositoryMongodb{
		looksCollection: database.Collection("looks"),
	}
}

func (lookRepositoryMongodb *LookRepositoryMongodb) FindLooks() []Look {
	var looksBson []LookBson
	looks := make([]Look, 0)

	cursor, err := lookRepositoryMongodb.looksCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	cursor.Decode(looksBson)
	/*
		if err = cursor.All(context.TODO(), &looksBson); err != nil {
			panic(err)
		}
	*/
	for _, lookBson := range looksBson {
		looks = append(looks, *NewLook(lookBson._id, lookBson.Colour, lookBson.Brand))
	}

	return looks
}

package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	database   = "favorites"
	collection = "users_history"
)

type HistoryMongo struct {
	db *mongo.Client
}

func NewHistoryMongo(db *mongo.Client) *HistoryMongo {
	return &HistoryMongo{
		db: db,
	}
}

func (historyMongo HistoryMongo) GetAll() {
	favoritesCollection := historyMongo.db.Database(database).Collection(collection)

	cursor, err := favoritesCollection.Find(context.TODO(), bson.D{})

	if err != nil {
		panic(err)
	}

	var results []bson.M
	// check for errors in the conversion
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	// display the documents retrieved
	fmt.Println("displaying all results in a collection")
	for _, result := range results {
		fmt.Println(result)
	}
}

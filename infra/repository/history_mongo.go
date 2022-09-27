package repository

import (
	"context"
	"github.com/WenLopes/recent/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	database   = "favorites"
	collection = "users_history"
)

type HistoryMongo struct {
	collection *mongo.Collection
}

func NewHistoryMongo(db *mongo.Client) *HistoryMongo {
	historyCollection := db.Database(database).Collection(collection)

	return &HistoryMongo{
		collection: historyCollection,
	}
}

func (historyMongo HistoryMongo) GetAll() []domain.History {

	cursor, err := historyMongo.collection.Find(context.TODO(), bson.D{})

	if err != nil {
		panic(err)
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	var histories []domain.History

	for _, result := range results {
		history := new(domain.History)
		history.UserId = result["user_id"].(string)
		if result["histories"] != nil {
			dado := result["histories"].(bson.A)
			for _, d := range dado {
				novodado := d.(bson.M)
				history.Id = int(novodado["id"].(int32))
				history.HistoryType = novodado["key_addressing_type"].(string)
				history.HistoryType = novodado["type"].(string)
			}
		}
		histories = append(histories, *history)
	}
	return histories
}

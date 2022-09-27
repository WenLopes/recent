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

type UsersHistoryMongo struct {
	collection *mongo.Collection
}

func NewUsersHistoryMongo(db *mongo.Client) *UsersHistoryMongo {
	historyCollection := db.Database(database).Collection(collection)

	return &UsersHistoryMongo{
		collection: historyCollection,
	}
}

//TODO: 1. Receber user ID como parâmetro
//TODO: 2. Modificar retorno
func (historyMongo UsersHistoryMongo) GetAllHistories() []domain.History {

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

package main

import (
	"context"
	"fmt"
	"github.com/WenLopes/recent/infra/repository"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
)

func main() {

	usersHistoryRepository := repository.NewHistoryMongo(initMongoClient())
	usersHistoryRepository.GetAll()

	router := mux.NewRouter()

	router.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(200)
		writer.Write([]byte("Hello, world!"))
	}).Methods("GET")

	fmt.Printf("Api pronta para receber requisições na porta 16092 🏆\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func initMongoClient() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://192.168.0.20:27051"))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	return client
}

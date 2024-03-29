package main

import (
	"context"
	"fmt"
	"github.com/WenLopes/recent/app/api"
	"github.com/WenLopes/recent/app/api/handlers"
	"github.com/WenLopes/recent/infra/repository"
	"github.com/WenLopes/recent/usecases/users_history"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
)

func main() {

	//TODO: Pesquisar boas práticas. Usar camelCase ou snake_case nos packages?
	router := mux.NewRouter()
	server := api.NewServer(initHandlers())
	server.Router(router)

	fmt.Printf("Api pronta para receber requisições na porta 16093 🏆\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func initHandlers() *api.Handlers {
	h := new(api.Handlers)
	h.AllHandler = initAllHandler()
	return h
}

func initAllHandler() *handlers.AllHandler {
	usersHistoryRepository := repository.NewUsersHistoryMongo(initMongoClient())
	usersHistoryService := users_history.NewService(usersHistoryRepository)
	return handlers.NewAllHandler(usersHistoryService)
}

func initMongoClient() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://192.168.0.2:27051"))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	return client
}

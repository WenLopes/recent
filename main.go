package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(200)
		writer.Write([]byte("Hello, world!"))
	}).Methods("GET")

	fmt.Printf("Api pronta para receber requisições na porta 16092 🏆\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}

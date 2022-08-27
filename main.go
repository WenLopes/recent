package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func main() {
	//router := mux.NewRouter()
	//
	//router.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Header().Set("Content-Type", "application/json")
	//	writer.WriteHeader(200)
	//	writer.Write([]byte("Hello, world!"))
	//}).Methods("GET")
	//
	//fmt.Printf("Api pronta para receber requisições 🏆\n")
	//log.Fatal(http.ListenAndServe(":8080", router))
	ctx := context.Background()
	consume(ctx)
}

func consume(ctx context.Context) {
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9093"},
		Topic:     "pix_transaction_new-history",
		GroupID:   "favorites",
		Partition: 0,
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))
	}
}

func configQueFunciona() {
	topic := "pix_transaction_new-history"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9093", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}

	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}

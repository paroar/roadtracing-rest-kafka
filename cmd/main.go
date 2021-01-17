package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/paroar/roadtracing-rest-kafka/internal/handlers"
	"github.com/paroar/roadtracing-rest-kafka/internal/kafka"
)

func main() {
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Printf("Listening on port: %s", port)
	http.HandleFunc("/", handlers.PositionHandler)
	kafka.NewProducer()
	panic(http.ListenAndServe(port, nil))
}

package main

import (
	"fmt"
	"gopkg.in/redis.v4"
	"log"
	"net/http"
)

var client = getClient()

func main() {
	router := NewRouter()

	fmt.Println("Starting server...")

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}

func getClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	fmt.Println("Connected to server")

	return client
}

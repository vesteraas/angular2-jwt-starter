package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	fmt.Println("Starting server...")

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}

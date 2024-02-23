package main

import (
	"log"
)

func main() {
	store, err := NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}
	server := NewAPIServer(":2000", store)
	err = server.Run()

	if err != nil {
		log.Println("Error starting server:", err)
		log.Fatal(err)
	}
}

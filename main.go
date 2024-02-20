package main

import (
	"log"
)

func main() {
	server := NewAPIServer(":2000")
	err := server.Run()

	if err != nil {
		log.Println("Error starting server:", err)
		log.Fatal(err)
	}
}

package main

import (
	"log"

	"github.com/calvinnle/go-chat/db"
)

func main() {
	_, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Connected to the database successfully.")
}

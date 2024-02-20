package main

import (
	"fmt"
	"log"
)

func main() {
	store, err := NewPostgressStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	server := NewAPIServer(":3000", store)
	server.Run()
	fmt.Println("Starting server...")
}

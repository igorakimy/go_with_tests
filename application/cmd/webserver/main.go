package main

import (
	poker "github.com/igorakimy/go_with_tests/application"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {

	store, closeFunc, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer closeFunc()

	server, err := poker.NewPlayerServer(store)
	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000: %v", err)
	}
}

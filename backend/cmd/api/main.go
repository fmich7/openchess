package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rekjef/openchess/internal"
	"github.com/rekjef/openchess/internal/config"
	"github.com/rekjef/openchess/internal/live_storage"
	"github.com/rekjef/openchess/internal/storage"
	"github.com/rekjef/openchess/internal/types"
)

func run(context context.Context) error {
	defer context.Done()
	config, err := config.LoadConfig(config.DEV)
	if err != nil {
		return err
	}

	// setup db
	credentials, err := storage.GetPostgressCredentials()
	if err != nil {
		return err
	}
	store, err := storage.NewPostgressStore(credentials)
	if err != nil {
		return err
	}

	if err := store.Init(); err != nil {
		return err
	}

	liveStore := live_storage.NewRAMStore()

	server := types.NewServer(config, store, liveStore)

	router := mux.NewRouter()
	internal.AddRoutes(router, server)

	handler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	)(router)

	if err = http.ListenAndServe(":"+config.Port, handler); err != nil {
		return err
	}
	log.Printf("Server is running on port: %s\n", config.Port)

	return nil
}

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

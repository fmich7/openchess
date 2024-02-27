package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rekjef/openchess/internal/config"
	"github.com/rekjef/openchess/internal/database"
	"github.com/rekjef/openchess/internal/routes"
	"github.com/rekjef/openchess/pkg/utils"
)

func run(context context.Context, env string) error {
	defer context.Done()
	// logger
	logger := utils.NewLogger("console")

	// load env
	config := config.NewEnv(logger)
	if err := config.LoadENV("dev"); err != nil {
		return err
	}

	// setup db
	store, err := database.NewPostgressStore(config)
	if err != nil {
		return err
	}

	// init if empty db
	if err := store.Init(); err != nil {
		return err
	}

	// server stuff
	mux := mux.NewRouter()
	routes.AddRoutes(mux, logger, store)

	port := config.GetEnv("PORT")
	logger.Info.Printf("Server is running on port: %s\n", port)

	if err = http.ListenAndServe(":"+port, mux); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(context.Background(), "dev"); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

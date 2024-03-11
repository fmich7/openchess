package routes

import (
	"github.com/gorilla/mux"
	"github.com/rekjef/openchess/internal/api/auth"
	"github.com/rekjef/openchess/internal/api/handlers"
	"github.com/rekjef/openchess/internal/database"
	"github.com/rekjef/openchess/pkg/utils"
)

func AddRoutes(mux *mux.Router, logger *utils.Logger, store database.Storage) {
	mux.HandleFunc("/login", handlers.HandleLogin(store))
	mux.HandleFunc("/account", handlers.HandleAccount(logger, store))
	// protect this in future
	mux.HandleFunc("/account/{id}", handlers.HandleAccountByID(logger, store))
	mux.HandleFunc("/game", handlers.HandleManagingChessGame(store))
	mux.HandleFunc("/game/{id}", auth.WithJWTAuth(handlers.HandleChessGame(store), store))
	mux.Handle("/whoami", auth.WhoAmI(logger))
	mux.Handle("/", handlers.HandleNotFound(logger))
}

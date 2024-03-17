package routes

import (
	"github.com/gorilla/mux"
	"github.com/rekjef/openchess/internal/api/auth"
	"github.com/rekjef/openchess/internal/api/handlers"
	"github.com/rekjef/openchess/internal/types"
	"github.com/rekjef/openchess/pkg/utils"
)

func AddRoutes(
	mux *mux.Router,
	logger *utils.Logger,
	store types.Storage,
	liveGameStore types.LiveGameStorage,
) {
	mux.Handle("/whoami", auth.WhoAmI(logger))
	mux.HandleFunc("/login", handlers.HandleLogin(store))
	mux.HandleFunc("/account", handlers.HandleAccount(logger, store))
	mux.HandleFunc("/account/{id}", handlers.HandleAccountByID(logger, store))

	mux.HandleFunc("/game", handlers.HandleManagingChessGame(store, liveGameStore))
	mux.HandleFunc("/game/{id}", auth.WithJWTAuth(handlers.HandleChessGame(store), store))

	mux.HandleFunc("/live/{id}", handlers.HandleLiveChessGameByID(liveGameStore))

	mux.Handle("/", handlers.HandleNotFound(logger))
}

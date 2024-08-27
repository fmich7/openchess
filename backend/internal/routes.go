package internal

import (
	"github.com/gorilla/mux"
	"github.com/rekjef/openchess/internal/auth"
	"github.com/rekjef/openchess/internal/handlers"
	"github.com/rekjef/openchess/internal/types"
)

func AddRoutes(
	mux *mux.Router,
	server *types.Server,
) {
	mux.Handle("/whoami", auth.WhoAmI())
	mux.HandleFunc("/login", handlers.HandleLogin(server.Store))
	mux.HandleFunc("/account", handlers.HandleAccount(server.Store))
	mux.HandleFunc("/account/{id}", handlers.HandleAccountByID(server.Store))
	mux.HandleFunc("/profile/{id}", handlers.HandleProfile(server.Store))

	mux.HandleFunc("/game", handlers.HandleManagingChessGame(server.Store, server.LiveStore))
	mux.HandleFunc("/game/{id}", auth.WithJWTAuth(handlers.HandleChessGame(server.Store), server.Store))

	mux.HandleFunc("/live/{id}", handlers.HandleLiveChessGameByID(server.LiveStore, server.Store))

	mux.HandleFunc("/leaderboard", handlers.HandleLeaderboard(server.Store))
	mux.HandleFunc("/stats", handlers.HandleStats(server.Store))

	mux.Handle("/", handlers.HandleNotFound())
}

package handlers

import (
	"net/http"

	"github.com/rekjef/openchess/internal/api"
	"github.com/rekjef/openchess/internal/types"
	"github.com/rekjef/openchess/pkg/utils"
)

// Get game data from FINISHED GAMES STORAGE
func getGameByID(id int, w http.ResponseWriter, store types.Storage) error {
	game, err := store.GetChessGameByID(id)
	if err != nil {
		return err
	}
	return utils.Encode(w, http.StatusOK, game)
}

// Update game in FINISHED GAMES STORAGE
func updateGameByID(id int, w http.ResponseWriter, store types.Storage) error {
	game, err := store.GetChessGameByID(id)
	if err != nil {
		return err
	}
	game.GameFEN = "UPDATED"
	game.GameStatus = "UPDATED?"
	store.UpdateChessGame(game)

	return utils.Encode(w, http.StatusOK, "game has been updated")
}

// HANDLE: /game/{id}
func HandleChessGame(store types.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := api.GetID(r)
		if err != nil {
			api.SendError(w, http.StatusBadRequest, err)
			return
		}

		switch r.Method {
		case "GET":
			err := getGameByID(id, w, store)
			api.SendError(w, http.StatusBadRequest, err)
		case "PATCH":
			err := updateGameByID(id, w, store)
			api.SendError(w, http.StatusBadRequest, err)
		default:
			api.MethodNotAllowed(w, r)
		}
	}
}

package handlers

import (
	"net/http"

	"github.com/rekjef/openchess/internal/api"
	"github.com/rekjef/openchess/internal/types"
	"github.com/rekjef/openchess/pkg/utils"
)

// Get all finished chess games
func getGamesData(w http.ResponseWriter, r *http.Request, store types.Storage) error {
	games, err := store.GetChessGames()
	if err != nil {
		return err
	}

	return utils.Encode(w, http.StatusOK, games)
}

// Create new game
func createNewGame(store types.Storage, liveGameStore types.LiveGameStorage, w http.ResponseWriter, r *http.Request) error {
	var gameReq types.CreateGameRequest
	if err := utils.Decode[types.CreateGameRequest](r, &gameReq); err != nil {
		return err
	}

	// store game in db
	game := types.NewChessGame(gameReq)
	if err := store.CreateChessGame(game); err != nil {
		return err
	}
	// store as live game
	if err := liveGameStore.AddGame(*game); err != nil {
		return err
	}

	return utils.Encode(w, http.StatusOK, game)
}

// HANDLE: /game
func HandleManagingChessGame(store types.Storage, liveGameStore types.LiveGameStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			if err := getGamesData(w, r, store); err != nil {
				api.SendError(w, http.StatusBadRequest, err)
			}
		case "POST":
			if err := createNewGame(store, liveGameStore, w, r); err != nil {
				api.SendError(w, http.StatusBadRequest, err)
			}
		default:
			api.MethodNotAllowed(w, r)
		}
	}
}

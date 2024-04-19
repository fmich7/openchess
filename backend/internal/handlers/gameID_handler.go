package handlers

import (
	"net/http"

	"github.com/rekjef/openchess/internal/types"
	"github.com/rekjef/openchess/internal/utils"
)

// Get game data from FINISHED GAMES STORAGE
func getGameByID(id int, w http.ResponseWriter, store types.Storage) error {
	game, err := store.GetChessGameByID(id)
	if err != nil {
		return err
	}
	return utils.Encode(w, http.StatusOK, game)
}

// HANDLE: /game/{id}
func HandleChessGame(store types.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := utils.GetID(r)
		if err != nil {
			utils.SendError(w, http.StatusBadRequest, err)
			return
		}

		switch r.Method {
		case "GET":
			err := getGameByID(id, w, store)
			utils.SendError(w, http.StatusBadRequest, err)
		default:
			utils.MethodNotAllowed(w, r)
		}
	}
}

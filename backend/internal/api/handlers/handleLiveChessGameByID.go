package handlers

import (
	"net/http"

	"github.com/rekjef/openchess/internal/api"
	"github.com/rekjef/openchess/internal/types"
	"github.com/rekjef/openchess/pkg/utils"
)

// Get LiveGame by ID
func getLiveGame(id int, store types.LiveGameStorage, w http.ResponseWriter) error {
	game, err := store.GetGame(id)
	if err != nil {
		return err
	}

	return utils.Encode(w, http.StatusOK, game)
}

// Update live game state
func updateLiveGameState(id int, store types.LiveGameStorage, w http.ResponseWriter, r *http.Request) error {
	options := new(types.GameUpdateOptions)
	if err := utils.Decode[types.GameUpdateOptions](r, options); err != nil {
		return err
	}

	_, err := store.GetGame(id)
	if err != nil {
		return err
	}

	store.UpdateGame(id, *options)

	return utils.Encode(w, http.StatusOK, *options)
}

// HANDLE: /live/{id}
func HandleLiveChessGameByID(liveGameStore types.LiveGameStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := api.GetID(r)
		if err != nil {
			api.SendError(w, http.StatusBadRequest, err)
			return
		}

		switch r.Method {
		case "GET":
			if err := getLiveGame(id, liveGameStore, w); err != nil {
				api.SendError(w, http.StatusBadRequest, err)
			}
		case "PUT":
			if err := updateLiveGameState(id, liveGameStore, w, r); err != nil {
				api.SendError(w, http.StatusBadRequest, err)
			}
		default:
			api.MethodNotAllowed(w, r)
		}
	}
}

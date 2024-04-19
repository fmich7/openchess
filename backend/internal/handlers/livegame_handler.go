package handlers

import (
	"net/http"

	"github.com/rekjef/openchess/internal/types"
	"github.com/rekjef/openchess/internal/utils"
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
func updateLiveGameState(
	id int, liveGameStore types.LiveGameStorage, store types.Storage,
	w http.ResponseWriter, r *http.Request,
) error {
	options := new(types.GameUpdateOptions)
	if err := utils.Decode[types.GameUpdateOptions](r, options); err != nil {
		return err
	}

	game, err := liveGameStore.GetGame(id)
	if err != nil {
		return err
	}

	if err := liveGameStore.UpdateGame(id, *options, store); err != nil {
		return err
	}

	type UpdateState struct {
		FEN string `json:"fen"`
	}

	return utils.Encode[UpdateState](w, http.StatusOK, UpdateState{FEN: game.Engine.FEN()})
}

// HANDLE: /live/{id}
func HandleLiveChessGameByID(
	liveGameStore types.LiveGameStorage,
	store types.Storage,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := utils.GetID(r)
		if err != nil {
			utils.SendError(w, http.StatusBadRequest, err)
			return
		}

		switch r.Method {
		case "GET":
			if err := getLiveGame(id, liveGameStore, w); err != nil {
				utils.SendError(w, http.StatusBadRequest, err)
			}
		case "PUT":
			if err := updateLiveGameState(id, liveGameStore, store, w, r); err != nil {
				utils.SendError(w, http.StatusBadRequest, err)
			}
		default:
			utils.MethodNotAllowed(w, r)
		}
	}
}

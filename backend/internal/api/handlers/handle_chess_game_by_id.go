package handlers

import (
	"errors"
	"net/http"

	"github.com/rekjef/openchess/internal/api"
	"github.com/rekjef/openchess/internal/database"
	"github.com/rekjef/openchess/pkg/utils"
)

func getGameByID(id int, w http.ResponseWriter, r *http.Request, store database.Storage) error {
	game, err := store.GetChessGameByID(id)
	if err != nil {
		return err
	}
	return utils.Encode(w, http.StatusOK, game)
}

func updateGameByID(id int, w http.ResponseWriter, r *http.Request, store database.Storage) error {

	return utils.Encode(w, http.StatusOK, "game has been updated")
}

func HandleChessGame(store database.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := api.GetID(r)
		if err != nil {
			api.SendError(w, http.StatusBadRequest, err)
			return
		}

		switch r.Method {
		case "GET":
			err := getGameByID(id, w, r, store)
			api.SendError(w, http.StatusBadRequest, err)
		case "PATCH":
			err := updateGameByID(id, w, r, store)
			api.SendError(w, http.StatusBadRequest, err)
		default:
			api.SendError(w, http.StatusMethodNotAllowed, errors.New("method not allowed "+r.Method))
		}
	}
}

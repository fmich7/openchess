package handlers

import (
	"net/http"

	"github.com/rekjef/openchess/internal/api"
	"github.com/rekjef/openchess/internal/database"
	"github.com/rekjef/openchess/internal/types"
	"github.com/rekjef/openchess/pkg/utils"
)

func createNewGame(w http.ResponseWriter, r *http.Request, store database.Storage) error {
	var gameReq types.CreateGameRequest
	if err := utils.Decode[types.CreateGameRequest](r, &gameReq); err != nil {
		return err
	}

	game := types.NewChessGame(gameReq)
	// write game to db
	if err := store.CreateChessGame(game); err != nil {
		return err
	}

	return utils.Encode(w, http.StatusOK, game)
}

func getGamesData(w http.ResponseWriter, r *http.Request, store database.Storage) error {
	games, err := store.GetChessGames()
	if err != nil {
		return err
	}

	return utils.Encode(w, http.StatusOK, games)
}

func HandleManagingChessGame(store database.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			err := getGamesData(w, r, store)
			api.SendError(w, http.StatusBadRequest, err)
		case "POST":
			err := createNewGame(w, r, store)
			api.SendError(w, http.StatusBadRequest, err)
		default:
			api.MethodNotAllowed(w, r)
		}
	}
}

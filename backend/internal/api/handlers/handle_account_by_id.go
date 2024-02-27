package handlers

import (
	"errors"
	"net/http"

	"github.com/rekjef/openchess/internal/api"
	"github.com/rekjef/openchess/internal/database"
	"github.com/rekjef/openchess/pkg/utils"
)

func handleDeleteAccountByID(w http.ResponseWriter, r *http.Request, store database.Storage) error {
	id, err := api.GetID(r)
	if err != nil {
		return err
	}

	if err := store.DeleteAccount(id); err != nil {
		return err
	}

	return utils.Encode(w, http.StatusOK, map[string]int{"deleted": id})
}

func handleGetAccountByID(w http.ResponseWriter, r *http.Request, store database.Storage) error {
	id, err := api.GetID(r)
	if err != nil {
		return err
	}

	account, err := store.GetAccountByID(id)
	if err != nil {
		return err
	}
	return utils.Encode(w, http.StatusOK, account)
}

func HandleAccountByID(logger *utils.Logger, store database.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			err := handleGetAccountByID(w, r, store)
			api.SendError(w, http.StatusBadRequest, err)
		case "DELETE":
			err := handleDeleteAccountByID(w, r, store)
			api.SendError(w, http.StatusBadRequest, err)
		default:
			api.SendError(w, http.StatusMethodNotAllowed, errors.New("method not allowed "+r.Method))
		}
	}
}

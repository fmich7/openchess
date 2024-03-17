package handlers

import (
	"net/http"

	"github.com/rekjef/openchess/internal/api"
	"github.com/rekjef/openchess/internal/types"
	"github.com/rekjef/openchess/pkg/utils"
)

// Delete account by ID
func deleteAccountByID(w http.ResponseWriter, r *http.Request, store types.Storage) error {
	id, err := api.GetID(r)
	if err != nil {
		return err
	}

	if err := store.DeleteAccount(id); err != nil {
		return err
	}

	return utils.Encode(w, http.StatusOK, map[string]int{"deleted": id})
}

// Get account by ID
func getAccountByID(w http.ResponseWriter, r *http.Request, store types.Storage) error {
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

// HANDLE: /account/{id}
func HandleAccountByID(logger *utils.Logger, store types.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			err := getAccountByID(w, r, store)
			api.SendError(w, http.StatusBadRequest, err)
		case "DELETE":
			err := deleteAccountByID(w, r, store)
			api.SendError(w, http.StatusBadRequest, err)
		default:
			api.MethodNotAllowed(w, r)
		}
	}
}

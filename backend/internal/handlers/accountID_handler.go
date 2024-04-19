package handlers

import (
	"net/http"

	"github.com/rekjef/openchess/internal/types"
	"github.com/rekjef/openchess/internal/utils"
)

// Delete account by ID
func deleteAccountByID(w http.ResponseWriter, r *http.Request, store types.Storage) error {
	id, err := utils.GetID(r)
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
	id, err := utils.GetID(r)
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
func HandleAccountByID(store types.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			err := getAccountByID(w, r, store)
			utils.SendError(w, http.StatusBadRequest, err)
		case "DELETE":
			err := deleteAccountByID(w, r, store)
			utils.SendError(w, http.StatusBadRequest, err)
		default:
			utils.MethodNotAllowed(w, r)
		}
	}
}

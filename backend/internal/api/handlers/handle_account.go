package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rekjef/openchess/internal/api"
	"github.com/rekjef/openchess/internal/database"
	"github.com/rekjef/openchess/internal/types"
	"github.com/rekjef/openchess/pkg/utils"
)

func handleGetAccount(w http.ResponseWriter, store database.Storage) error {
	accounts, err := store.GetAccounts()
	if err != nil {
		return err
	}

	return utils.Encode(w, http.StatusOK, accounts)
}

func handleCreateAccount(w http.ResponseWriter, r *http.Request, store database.Storage) error {
	accountRequest := new(types.CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(accountRequest); err != nil {
		return err
	}
	defer r.Body.Close()

	if len(accountRequest.Password) < 3 {
		return utils.Encode(w, http.StatusBadRequest, api.ApiError{Error: "invalid password"})
	}

	account, err := types.NewAccount(
		accountRequest.FirstName,
		accountRequest.LastName,
		accountRequest.Nickname, accountRequest.Password)

	if err != nil {
		return err
	}

	if err := store.CreateAccount(account); err != nil {
		return err
	}

	return utils.Encode(w, http.StatusOK, account)
}

func HandleAccount(logger *utils.Logger, store database.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			err := handleGetAccount(w, store)
			api.SendError(w, http.StatusBadRequest, err)
		case "POST":
			err := handleCreateAccount(w, r, store)
			api.SendError(w, http.StatusBadRequest, err)
		default:
			api.SendError(w, http.StatusMethodNotAllowed, errors.New("method not allowed "+r.Method))
		}

	}
}

package handlers

import (
	"net/http"

	"github.com/rekjef/openchess/internal/account"
	"github.com/rekjef/openchess/internal/types"
	"github.com/rekjef/openchess/internal/utils"
)

// Get all accounts data
func getAccounts(w http.ResponseWriter, store types.Storage) error {
	accounts, err := store.GetAccounts()
	if err != nil {
		return err
	}

	return utils.Encode(w, http.StatusOK, accounts)
}

// Create account
func createAccount(w http.ResponseWriter, r *http.Request, store types.Storage) error {
	accountRequest := new(types.CreateAccountRequest)
	if err := utils.Decode[types.CreateAccountRequest](r, accountRequest); err != nil {
		return err
	}
	defer r.Body.Close()

	if len(accountRequest.Password) < 3 {
		return utils.Encode(w, http.StatusBadRequest, utils.ApiError{Error: "invalid password"})
	}

	acc, err := account.NewAccount(
		accountRequest.FirstName,
		accountRequest.LastName,
		accountRequest.Nickname, accountRequest.Password)

	if err != nil {
		return err
	}

	id, err := store.CreateAccount(*acc)
	if err != nil {
		return err
	}
	acc.ID = id

	return utils.Encode(w, http.StatusOK, acc)
}

// HANDLE: /account
func HandleAccount(store types.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			err := getAccounts(w, store)
			utils.SendError(w, http.StatusBadRequest, err)
		case "POST":
			err := createAccount(w, r, store)
			utils.SendError(w, http.StatusBadRequest, err)
		default:
			utils.MethodNotAllowed(w, r)
		}

	}
}

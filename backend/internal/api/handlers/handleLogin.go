package handlers

import (
	"errors"
	"net/http"

	"github.com/rekjef/openchess/internal/api"
	"github.com/rekjef/openchess/internal/api/auth"
	"github.com/rekjef/openchess/internal/types"
	"github.com/rekjef/openchess/pkg/utils"
)

// Login user, check credentials, return TOKEN, ID
func loginUser(w http.ResponseWriter, r *http.Request, store types.Storage) error {
	loginReq := new(types.LoginRequest)
	if err := utils.Decode[types.LoginRequest](r, loginReq); err != nil {
		return err
	}

	acc, err := store.GetAccountByNickname(loginReq.Nickname)
	if err != nil {
		return api.SendError(w, http.StatusUnauthorized, errors.New("not authenticated"))
	}

	if !acc.ComparePasswords(loginReq.Password) {
		return api.SendError(w, http.StatusUnauthorized, errors.New("not authenticated"))
	}

	tokenString, err := auth.CreateJWT(acc)
	if err != nil {
		return err
	}

	resp := types.LoginResponse{
		Token: tokenString,
		ID:    acc.ID,
	}

	return utils.Encode(w, http.StatusOK, resp)
}

// HANDLE: /login
func HandleLogin(store types.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			err := loginUser(w, r, store)
			api.SendError(w, http.StatusBadRequest, err)
		default:
			api.MethodNotAllowed(w, r)
		}

	}
}

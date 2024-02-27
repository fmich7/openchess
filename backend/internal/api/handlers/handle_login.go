package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rekjef/openchess/internal/api"
	"github.com/rekjef/openchess/internal/api/auth"
	"github.com/rekjef/openchess/internal/database"
	"github.com/rekjef/openchess/internal/types"
	"github.com/rekjef/openchess/pkg/utils"
)

func handlePostLogin(w http.ResponseWriter, r *http.Request, store database.Storage) error {
	var loginReq types.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		return err
	}

	acc, err := store.GetAccountByNickname(loginReq.Nickname)
	if err != nil {
		return err
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

func HandleLogin(store database.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			err := handlePostLogin(w, r, store)
			api.SendError(w, http.StatusBadRequest, err)
		default:
			api.SendError(w, http.StatusMethodNotAllowed, errors.New("method not allowed "+r.Method))
		}

	}
}

package handlers

import (
	"net/http"

	"github.com/rekjef/openchess/internal/account"
	"github.com/rekjef/openchess/internal/types"
	"github.com/rekjef/openchess/internal/utils"
)

// HANDLE: /profile
func HandleProfile(store types.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			id, err := utils.GetID(r)
			if err != nil {
				utils.SendError(w, http.StatusBadRequest, err)
				return
			}
			profile, err := account.GetProfile(id, store)
			if err != nil {
				utils.SendError(w, http.StatusBadRequest, err)
				return
			}
			utils.Encode[account.Profile](w, http.StatusOK, profile)
		default:
			utils.MethodNotAllowed(w, r)
		}

	}
}

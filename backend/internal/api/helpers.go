package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rekjef/openchess/pkg/utils"
)

type ApiError struct {
	Error string `json:"error"`
}

func SendError(w http.ResponseWriter, status int, err error) error {
	if err != nil {
		return utils.Encode(w, status, ApiError{Error: err.Error()})
	}
	return nil
}

// type apiFunc func(http.ResponseWriter, *http.Request) error

// func MakeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if err := f(w, r); err != nil {
// 			// handle the error
// 			utils.Encode(w, http.StatusBadRequest, ApiError{Error: err.Error()})
// 		}
// 	}
// }

func PermissionDenied(w http.ResponseWriter) {
	utils.Encode(w, http.StatusForbidden, ApiError{Error: "permission denied"})
}

func GetID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid given id %s", idStr)
	}
	return id, nil
}

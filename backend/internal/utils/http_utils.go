package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ApiError struct {
	Error string `json:"error"`
}

func SendError(w http.ResponseWriter, status int, err error) error {
	if err != nil {
		return Encode(w, status, ApiError{Error: err.Error()})
	}
	return nil
}

// type apiFunc func(http.ResponseWriter, *http.Request) error

//	func MakeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
//		return func(w http.ResponseWriter, r *http.Request) {
//			if err := f(w, r); err != nil {
//				// handle the error
//				utils.Encode(w, http.StatusBadRequest, ApiError{Error: err.Error()})
//			}
//		}
//	}
func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	SendError(w, http.StatusMethodNotAllowed, errors.New("method not allowed "+r.Method))
}

func PermissionDenied(w http.ResponseWriter) {
	Encode(w, http.StatusForbidden, ApiError{Error: "permission denied"})
}

// requested game is not active
func NoActiveGameError(id int) error {
	return fmt.Errorf("there is no active game with give ID %d", id)
}

func GetID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid given id %s", idStr)
	}
	return id, nil
}

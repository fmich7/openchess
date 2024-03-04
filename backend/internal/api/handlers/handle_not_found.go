package handlers

import (
	"errors"
	"net/http"

	"github.com/rekjef/openchess/internal/api"
	"github.com/rekjef/openchess/pkg/utils"
)

func HandleNotFound(logger *utils.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Error.Println("Request URL not found: ", r.URL.Path)
		api.SendError(w, http.StatusNotFound, errors.New("path not handled"))
	}
}

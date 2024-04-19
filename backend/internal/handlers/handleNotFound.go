package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/rekjef/openchess/internal/utils"
)

func HandleNotFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request URL not found: ", r.URL.Path)
		utils.SendError(w, http.StatusNotFound, errors.New("path not handled"))
	}
}

package handlers

import (
	"net/http"

	"github.com/rekjef/openchess/internal/stats"
	"github.com/rekjef/openchess/internal/types"
	"github.com/rekjef/openchess/internal/utils"
)

type LeaderboardRangeRequest struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

// HANDLE: /stats
func HandleLeaderboard(store types.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// json: start: int, end: int (optional)
		case "GET":
			var requestData LeaderboardRangeRequest
			var leaderboard []types.Account

			// without range
			if err := utils.Decode[LeaderboardRangeRequest](r, &requestData); err != nil {
				leaderboard, err = stats.GetLeaderboard(store)
				if err != nil {
					utils.SendError(w, http.StatusInternalServerError, err)
					return
				}
			} else { // with range
				leaderboard, err = stats.GetRangedLeaderboard(
					requestData.Start,
					requestData.End,
					store,
				)
				if err != nil {
					utils.SendError(w, http.StatusInternalServerError, err)
					return
				}
			}

			utils.Encode[[]types.Account](w, http.StatusOK, leaderboard)
		default:
			utils.MethodNotAllowed(w, r)
		}

	}
}

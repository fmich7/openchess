package stats

import (
	"sort"

	"github.com/rekjef/openchess/internal/types"
)

func GetLeaderboard(store types.Storage) ([]types.Account, error) {
	leaderboard, err := store.GetAccounts()

	if err != nil {
		return leaderboard, err
	}

	sort.Slice(leaderboard, func(i, j int) bool {
		return leaderboard[i].Elo > leaderboard[j].Elo
	})

	return leaderboard, nil
}

func GetRangedLeaderboard(start int, end int, store types.Storage) ([]types.Account, error) {
	leaderboard, err := GetLeaderboard(store)
	if err != nil {
		return leaderboard, err
	}

	return leaderboard[start:min(end, len(leaderboard))], nil
}

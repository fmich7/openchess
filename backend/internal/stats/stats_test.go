package stats

import (
	"testing"

	"github.com/rekjef/openchess/internal/tests"
	"github.com/rekjef/openchess/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestGetLeaderboard(t *testing.T) {
	t.Run("empty acc table", func(t *testing.T) {
		store := tests.NewFakeDB()
		leaderboard, err := GetLeaderboard(store)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, len(leaderboard), 0)
	})

	t.Run("two acc - check if sorted", func(t *testing.T) {
		store := tests.NewFakeDB()
		// add acc to fakedb
		for i := 0; i < 2; i++ {
			_, err := store.CreateAccount(types.Account{ID: i, Elo: 1200 + i})
			if err != nil {
				t.Error(err)
			}
		}
		leaderboard, err := GetLeaderboard(store)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, len(leaderboard), 2)
		// check if order is correct
		assert.Greater(t, leaderboard[0].Elo, leaderboard[1].Elo)
	})

	t.Run("get given range from leaderboard", func(t *testing.T) {
		store := tests.NewFakeDB()
		// add acc to fakedb
		for i := 0; i < 4; i++ {
			_, err := store.CreateAccount(types.Account{ID: i, Elo: 1200 + i})
			if err != nil {
				t.Error(err)
			}
		}
		leaderboard, err := GetRangedLeaderboard(1, 3, store)
		if err != nil {
			t.Error(err)
		}
		// check size
		assert.Equal(t, len(leaderboard), 2)
		assert.Equal(t, leaderboard[0].Elo, 1202)
		assert.Equal(t, leaderboard[1].Elo, 1201)
	})
}

func TestCountUsers(t *testing.T) {
	store := tests.NewFakeDB()
	userCount, err := CountUsers(store)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, userCount, 0)
	_, err = store.CreateAccount(types.Account{})
	if err != nil {
		t.Error(err)
	}

	userCount, err = CountUsers(store)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, userCount, 1)
}

func TestCountPlayedGames(t *testing.T) {
	store := tests.NewFakeDB()
	gamesCount, err := CountPlayedGames(store)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, gamesCount, 0)
	_, err = store.CreateChessGame(types.ChessGame{})
	if err != nil {
		t.Error(err)
	}

	gamesCount, err = CountPlayedGames(store)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, gamesCount, 1)
}

package live_storage

import (
	"testing"

	"github.com/notnil/chess"
	"github.com/rekjef/openchess/internal/tests"
	"github.com/rekjef/openchess/internal/types"
	"github.com/stretchr/testify/assert"
)

// e RAMStore map[int]*types.LiveGame

func TestNewRamStore(t *testing.T) {
	liveStore := NewRAMStore()
	assert.Equal(t, liveStore, &RAMStore{})
}

func TestAddGame(t *testing.T) {
	liveStore := NewRAMStore()
	store := tests.NewFakeDB()
	game := types.ChessGame{ID: 0}
	if err := liveStore.AddGame(game, store); err != nil {
		t.Error(err)
	}
	// duplicate
	if err := liveStore.AddGame(game, store); err == nil {
		t.Error("game should be live")
	}
	assert.Equal(t, len(*liveStore), 1)
}

func TestGetGame(t *testing.T) {
	liveStore := NewRAMStore()
	store := tests.NewFakeDB()
	game := types.ChessGame{ID: 0, GameFEN: "asdfasf"}

	if _, err := liveStore.GetGame(0); err == nil {
		t.Error("game not live")
	}
	liveStore.AddGame(game, store)
	recievedGame, err := liveStore.GetGame(0)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, recievedGame.Details, game)
	assert.Equal(t, len(*liveStore), 1)
}

func TestUpdateGame(t *testing.T) {

}

func TestDeleteGame(t *testing.T) {
	liveStore := NewRAMStore()
	game := types.ChessGame{ID: 0, GameFEN: "asdfasf"}

	if err := liveStore.DeleteGame(0); err == nil {
		t.Error("should raise error on deleting game that does not exist")
	}
	(*liveStore)[game.ID] = types.NewLiveGame(game)
	if err := liveStore.DeleteGame(0); err != nil {
		t.Error(err)
	}
}

func TestEndGame(t *testing.T) {
	liveStore := NewRAMStore()
	store := tests.NewFakeDB()
	game := types.ChessGame{ID: 0, GameFEN: "asdfasf"}

	t.Run("no game to end", func(t *testing.T) {
		if err := liveStore.EndGame(0, chess.Draw, store); err == nil {
			t.Error("game not live")
		}
	})

	t.Run("end game, check updated state", func(t *testing.T) {

		if _, err := store.CreateChessGame(game); err != nil {
			t.Fatal(err)
		}

		if err := liveStore.AddGame(game, store); err != nil {
			t.Fatal(err)
		}

		if err := liveStore.EndGame(game.ID, chess.Draw, store); err != nil {
			t.Fatal(err)
		}

		updatedGame, err := store.GetChessGameByID(game.ID)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, updatedGame.GameStatus, types.GameEnded)
	})
}

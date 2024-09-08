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
	game := types.ChessGame{ID: 0, Time: 3000, WhiteTime: 3000, BlackTime: 3000}
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
	liveStore := NewRAMStore()
	store := tests.NewFakeDB()
	game := types.ChessGame{ID: 0, WhiteToMove: true, Time: 30, WhitePlayerID: 0}

	t.Run("no game to update", func(t *testing.T) {
		if err := liveStore.UpdateGame(2, types.GameUpdateOptions{}, store); err == nil {
			t.Error("game is not live")
		}
	})
	t.Run("resign", func(t *testing.T) {
		id, err := store.CreateChessGame(game)
		if err != nil {
			t.Error(err)
		}
		game.ID = id
		if err := liveStore.AddGame(game, store); err != nil {
			t.Error(err)
		}

		if err := liveStore.UpdateGame(game.ID, types.GameUpdateOptions{Resign: true}, store); err != nil {
			t.Error("game is not live")
		}

		finishedGame, err := store.GetChessGameByID(game.ID)
		if err != nil {
			t.Error("finished game should be saved in db")
		}
		assert.Equal(t, chess.BlackWon, finishedGame.GameOutcome)
	})

	t.Run("draw", func(t *testing.T) {
		id, err := store.CreateChessGame(game)
		if err != nil {
			t.Error(err)
		}
		game.ID = id
		if err := liveStore.AddGame(game, store); err != nil {
			t.Error(err)
		}

		if err := liveStore.UpdateGame(game.ID, types.GameUpdateOptions{OfferDraw: true}, store); err != nil {
			t.Error("game is not live")
		}

		finishedGame, err := store.GetChessGameByID(game.ID)
		if err != nil {
			t.Error("finished game should be saved in db")
		}
		assert.Equal(t, chess.Draw, finishedGame.GameOutcome)
	})

	t.Run("invalid move", func(t *testing.T) {
		id, err := store.CreateChessGame(game)
		if err != nil {
			t.Error(err)
		}
		game.ID = id
		if err := liveStore.AddGame(game, store); err != nil {
			t.Error(err)
		}

		if err := liveStore.UpdateGame(game.ID, types.GameUpdateOptions{Move: "xd", UserID: 0}, store); err == nil {
			t.Error("didn't raise error on invalid move")
		}
	})

	t.Run("valid move", func(t *testing.T) {
		id, err := store.CreateChessGame(game)
		if err != nil {
			t.Error(err)
		}
		game.ID = id
		if err := liveStore.AddGame(game, store); err != nil {
			t.Error(err)
		}

		if err := liveStore.UpdateGame(game.ID, types.GameUpdateOptions{Move: "a2a3", UserID: 0}, store); err != nil {
			t.Error("did raise error on valid move")
		}
	})

	t.Run("valid move but wrong user", func(t *testing.T) {
		id, err := store.CreateChessGame(game)
		if err != nil {
			t.Error(err)
		}
		game.ID = id
		if err := liveStore.AddGame(game, store); err != nil {
			t.Error(err)
		}

		if err := liveStore.UpdateGame(game.ID, types.GameUpdateOptions{Move: "a2a3", UserID: -1}, store); err == nil {
			t.Error("didn't raise an error on wrong user id")
		}
	})
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
		if err := liveStore.EndGame(0, store); err == nil {
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

		// TODO: check for draw
		if err := liveStore.EndGame(game.ID, store); err != nil {
			t.Fatal(err)
		}

		updatedGame, err := store.GetChessGameByID(game.ID)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, updatedGame.GameStatus, types.GameEnded)
	})
}

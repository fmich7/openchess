package database

import (
	"testing"

	"github.com/rekjef/openchess/internal/types"
)

func testEmptyStore(store *RAMStore, t *testing.T) {
	// GetGame
	if _, err := store.GetGame(0); err == nil {
		t.Fatal()
	}
	// DeleteGame
	if err := store.DeleteGame(0); err == nil {
		t.Fatal()
	}
	// UpdateGame
	if err := store.UpdateGame(0, types.GameUpdateOptions{Move: "asd"}); err == nil {
		t.Fatal()
	}
}

func testStoreAfterSeeding(store *RAMStore, t *testing.T) {
	// GetGame
	if _, err := store.GetGame(0); err != nil {
		t.Fatal(err)
	}
	// UpdateGame
	if err := store.UpdateGame(0, types.GameUpdateOptions{Move: "asd"}); err != nil {
		t.Fatal(err)
	}
	// DeleteGame
	if err := store.DeleteGame(0); err != nil {
		t.Fatal(err)
	}
}

func TestLiveGameStorage(t *testing.T) {
	store := NewRAMStore()
	testEmptyStore(store, t)

	// seed storage
	store.AddGame(types.ChessGame{ID: 0})

	// test after
	testStoreAfterSeeding(store, t)
}

package tests

import (
	"testing"

	"github.com/rekjef/openchess/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestNewFakeDB(t *testing.T) {
	storage := NewFakeDB()
	assert := assert.New(t)
	assert.Equal(storage.AccID, 0)
	assert.Equal(storage.GameID, 0)
	assert.IsType(map[int]types.Account{}, storage.Accounts)
	assert.IsType(map[int]types.ChessGame{}, storage.Games)
}

// -------------------
// - ACCOUNT SECTION -
// -------------------
func TestCreateAccount(t *testing.T) {
	storage := NewFakeDB()
	id, err := storage.CreateAccount(types.Account{})

	assert := assert.New(t)
	assert.Equal(err, nil)
	assert.Equal(id, 0)
	assert.Equal(storage.AccID, 1)
}

func TestUpdateAccount(t *testing.T) {
	storage := NewFakeDB()
	if err := storage.UpdateAccount(&types.Account{}); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteAccount(t *testing.T) {
	storage := NewFakeDB()
	if err := storage.DeleteAccount(0); err == nil {
		t.Fatal("return nil on deleting acc with invalid id")
	}
	//  mock
	storage.Accounts[storage.AccID] = types.Account{}
	accountsCount := len(storage.Accounts)
	if err := storage.DeleteAccount(storage.AccID); err != nil {
		t.Fatal(err)
	}
	assert.Greater(t, accountsCount, len(storage.Accounts))
}

func TestGetAccountByID(t *testing.T) {
	storage := NewFakeDB()
	if _, err := storage.GetAccountByID(0); err == nil {
		t.Fatal("got user with invalid id")
	}
	//  mock
	storage.Accounts[storage.AccID] = types.Account{}
	if _, err := storage.GetAccountByID(0); err != nil {
		t.Fatal(err)
	}
}

func TestGetAccountByNickname(t *testing.T) {
	storage := NewFakeDB()
	if _, err := storage.GetAccountByNickname("test"); err == nil {
		t.Fatal("no error on acc that does not exist")
	}
	//  mock
	acc := types.Account{Nickname: "test"}
	storage.Accounts[storage.AccID] = acc
	retrievedAcc, err := storage.GetAccountByNickname("test")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, acc, *retrievedAcc)
}

func TestGetAccounts(t *testing.T) {
	storage := NewFakeDB()
	assert := assert.New(t)

	// empty
	accounts, err := storage.GetAccounts()
	assert.Equal(len(accounts), 0)
	assert.Equal(err, nil)

	// add acc
	acc := types.Account{Nickname: "test"}
	storage.Accounts[storage.AccID] = acc

	accounts, err = storage.GetAccounts()
	assert.Equal(len(accounts), 1)
	assert.Equal(err, nil)

}

// ----------------------
// - CHESS GAME SECTION -
// ----------------------

func TestCreateChessGame(t *testing.T) {
	storage := NewFakeDB()
	id, err := storage.CreateChessGame(types.ChessGame{})

	assert.Equal(t, err, nil)
	assert.Equal(t, id, 0)
	assert.Equal(t, storage.GameID, 1)
}

func TestUpdateChessGame(t *testing.T) {
	storage := NewFakeDB()
	if err := storage.UpdateChessGame(types.ChessGame{ID: 3}); err == nil {
		t.Fatal("no error on game that does not exist")
	}

	TEST_FEN := "test"
	game := types.ChessGame{GameFEN: TEST_FEN}
	storage.Games[storage.GameID] = game
	if err := storage.UpdateChessGame(types.ChessGame{ID: 0, GameFEN: ""}); err != nil {
		t.Fatal(err)
	}

	game = storage.Games[0]
	assert := assert.New(t)
	assert.NotEqual(game.GameFEN, TEST_FEN)
}

func TestGetChessGameByID(t *testing.T) {
	storage := NewFakeDB()
	if _, err := storage.GetChessGameByID(0); err == nil {
		t.Fatal("got user with invalid id")
	}
	//  mock
	storage.Games[storage.GameID] = types.ChessGame{}
	if _, err := storage.GetChessGameByID(0); err != nil {
		t.Fatal(err)
	}
}

func TestChessGames(t *testing.T) {
	storage := NewFakeDB()
	assert := assert.New(t)

	// empty
	games, err := storage.GetChessGames()
	assert.Equal(len(games), 0)
	assert.Equal(err, nil)

	// add game
	game := types.ChessGame{}
	storage.Games[storage.GameID] = game

	games, err = storage.GetChessGames()
	assert.Equal(len(games), 1)
	assert.Equal(err, nil)

}

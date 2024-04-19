package tests

import (
	"fmt"

	"github.com/rekjef/openchess/internal/types"
)

type Account = types.Account
type ChessGame = types.ChessGame

type FakeDB struct {
	Accounts map[int]Account
	AccID    int
	Games    map[int]ChessGame
	GameID   int
}

func NewFakeDB() *FakeDB {
	return &FakeDB{
		Accounts: make(map[int]Account),
		Games:    make(map[int]ChessGame),
	}
}

// -------------------
// - ACCOUNT SECTION -
// -------------------
func (f *FakeDB) CreateAccount(acc Account) (int, error) {
	accID := f.AccID
	f.Accounts[accID] = acc
	f.AccID++
	return accID, nil
}

func (f *FakeDB) UpdateAccount(acc *Account) error {
	return nil
}

func (f *FakeDB) DeleteAccount(id int) error {
	if _, ok := f.Accounts[id]; !ok {
		return fmt.Errorf("invalid account id %d", id)
	}
	delete(f.Accounts, id)
	return nil
}
func (f *FakeDB) GetAccountByID(id int) (*Account, error) {
	if _, ok := f.Accounts[id]; !ok {
		return nil, fmt.Errorf("invalid account id %d", id)
	}
	acc := f.Accounts[id]
	return &acc, nil
}

func (f *FakeDB) GetAccountByNickname(nickname string) (*Account, error) {
	for _, acc := range f.Accounts {
		if acc.Nickname == nickname {
			return &acc, nil
		}
	}
	return nil, fmt.Errorf("user with nickname=%s does not exist", nickname)
}

func (f *FakeDB) GetAccounts() ([]*Account, error) {
	accounts := []*Account{}
	for _, acc := range f.Accounts {
		accounts = append(accounts, &acc)
	}
	return accounts, nil
}

// ----------------------
// - CHESS GAME SECTION -
// ----------------------
func (f *FakeDB) CreateChessGame(chessGame ChessGame) (int, error) {
	gameID := f.GameID
	f.Games[gameID] = chessGame
	f.GameID++
	return gameID, nil
}

func (f *FakeDB) GetChessGames() ([]*ChessGame, error) {
	games := []*ChessGame{}
	for _, game := range f.Games {
		games = append(games, &game)
	}
	return games, nil
}

func (f *FakeDB) GetChessGameByID(id int) (*ChessGame, error) {
	if _, ok := f.Games[id]; !ok {
		return nil, fmt.Errorf("there is no game with id %d", id)
	}
	game := f.Games[id]
	return &game, nil
}

func (f *FakeDB) UpdateChessGame(updatedGame ChessGame) error {
	if _, ok := f.Games[updatedGame.ID]; !ok {
		return fmt.Errorf("there is no game with id %d", updatedGame.ID)
	}
	f.Games[updatedGame.ID] = updatedGame
	return nil
}

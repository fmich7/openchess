package types

import (
	"github.com/notnil/chess"
	"github.com/rekjef/openchess/internal/config"
)

type Storage interface {
	// Account
	CreateAccount(Account) (int, error)
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]Account, error)
	GetAccountByID(int) (*Account, error)
	GetAccountByNickname(string) (*Account, error)

	// ChessGame
	CreateChessGame(ChessGame) (int, error)
	GetChessGames() ([]*ChessGame, error)
	GetChessGameByID(int) (*ChessGame, error)
	UpdateChessGame(ChessGame) error
	UpdatePlayerStats(int, UserStats) error
}
type LiveGameStorage interface {
	AddGame(ChessGame, Storage) error
	GetGame(int) (*LiveGame, error)
	MakeMove(id int, move *chess.Move) error
	UpdateGame(int, GameUpdateOptions, Storage) error
	DeleteGame(int) error
	EndGame(int, Storage) error
}
type Server struct {
	Config    config.Config
	Store     Storage
	LiveStore LiveGameStorage
}

func NewServer(
	config config.Config,
	store Storage,
	liveStore LiveGameStorage,
) *Server {
	return &Server{
		Config: config, Store: store, LiveStore: liveStore,
	}
}

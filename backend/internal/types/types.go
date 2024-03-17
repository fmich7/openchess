package types

type Storage interface {
	// Account
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
	GetAccountByNickname(string) (*Account, error)

	// ChessGame
	CreateChessGame(*ChessGame) error
	GetChessGames() ([]*ChessGame, error)
	GetChessGameByID(int) (*ChessGame, error)
	UpdateChessGame(*ChessGame) error
}

type LiveGameStorage interface {
	AddGame(ChessGame) error
	GetGame(int) (*LiveGame, error)
	UpdateGame(int, GameUpdateOptions) error
	DeleteGame(int) error
}

// type ApiServer struct {
// 	router        *mux.Router
// 	logger        *utils.Logger
// 	store         Storage
// 	liveGameStore LiveGameStorage
// }

package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rekjef/openchess/internal/config"
	"github.com/rekjef/openchess/internal/types"
)

// binds
type Account = types.Account
type ChessGame = types.ChessGame

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
}

type PostgressStore struct {
	db *sql.DB
}

func NewPostgressStore(config *config.Env) (*PostgressStore, error) {
	connStr := fmt.Sprintf(
		"user=%s dbname=%s password=%s sslmode=%s",
		config.GetEnv("POSTGRES_USER"),
		config.GetEnv("POSTGRES_DBNAME"),
		config.GetEnv("POSTGRES_PASSWORD"),
		config.GetEnv("POSTGRES_SSLMODE"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgressStore{db: db}, nil
}

func (s *PostgressStore) Init() error {
	if err := s.createAccountTable(); err != nil {
		return err
	}

	if err := s.createChessGameTable(); err != nil {
		return err
	}

	return nil
}

func (s *PostgressStore) createAccountTable() error {
	query := `create table if not exists account (
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		nickname varchar(50),
		encrypted_password varchar(100),
		elo serial,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgressStore) createChessGameTable() error {
	query := `CREATE TABLE IF NOT EXISTS chessGame (
		id serial primary key,
		host_id integer,
		white_player_id integer,
		black_player_id integer,
		white_to_move boolean,
		game_fen varchar(500),
		game_type varchar(50),
		game_status varchar(50),
		game_ended boolean,
		game_won_by_white boolean,
		is_ranked boolean,
		time integer,
		time_added integer,
		moves_count integer,
		move_history varchar(1000),
		created_at timestamp
	)`

	_, err := s.db.Exec(query)

	return err
}
func (s *PostgressStore) CreateChessGame(chessGame *ChessGame) error {
	query := `insert into chessGame
	(
		host_id,
		white_player_id,
		black_player_id,
		white_to_move,
		game_fen,
		game_type,
		game_status,
		game_ended,
		game_won_by_white,
		is_ranked,
		time,
		time_added,
		moves_count,
		move_history,
		created_at
	) 
	values (
		$1, $2, $3, $4, $5, $6, $7, $8, 
		$9, $10, $11, $12, $13, $14, $15
	)`

	_, err := s.db.Query(
		query,
		chessGame.HostID,
		chessGame.WhitePlayerID,
		chessGame.BlackPlayerID,
		chessGame.WhiteToMove,
		chessGame.GameFEN,
		chessGame.GameType,
		chessGame.GameStatus,
		chessGame.GameEnded,
		chessGame.GameWonByWhite,
		chessGame.IsRanked,
		chessGame.Time,
		chessGame.TimeAdded,
		chessGame.MovesCount,
		chessGame.MoveHistory,
		chessGame.CreatedAt,
	)
	if err != nil {
		return err
	}

	return err
}

func scanIntoGame(rows *sql.Rows) (*ChessGame, error) {
	game := new(ChessGame)
	err := rows.Scan(
		&game.ID,
		&game.HostID,
		&game.WhitePlayerID,
		&game.BlackPlayerID,
		&game.WhiteToMove,
		&game.GameFEN,
		&game.GameType,
		&game.GameStatus,
		&game.GameEnded,
		&game.GameWonByWhite,
		&game.IsRanked,
		&game.Time,
		&game.TimeAdded,
		&game.MovesCount,
		&game.MoveHistory,
		&game.CreatedAt,
	)

	return game, err
}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Nickname,
		&account.EncryptedPassword,
		&account.Elo,
		&account.CreatedAt)
	return account, err
}

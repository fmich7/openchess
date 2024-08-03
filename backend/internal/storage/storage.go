package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/rekjef/openchess/internal/types"
)

type Account = types.Account
type ChessGame = types.ChessGame

type PostgressStore struct {
	db *sql.DB
}

type PostgressCredentials struct {
	Host     string
	User     string
	Dbname   string
	Password string
	Sslmode  string
}

func GetPostgressCredentials() (PostgressCredentials, error) {
	cred := PostgressCredentials{}
	credFields := map[string]*string{
		"POSTGRES_HOST":     &cred.Host,
		"POSTGRES_USER":     &cred.User,
		"POSTGRES_DBNAME":   &cred.Dbname,
		"POSTGRES_PASSWORD": &cred.Password,
		"POSTGRES_SSLMODE":  &cred.Sslmode,
	}

	for key, value := range credFields {
		envVal, exists := os.LookupEnv(key)
		if !exists {
			return PostgressCredentials{}, errors.New("environment variable " + key + " not set")
		}
		*value = envVal
	}

	return cred, nil
}

func NewPostgressStore(cred PostgressCredentials) (*PostgressStore, error) {
	connStr := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s sslmode=%s",
		cred.Host,
		cred.User,
		cred.Dbname,
		cred.Password,
		cred.Sslmode,
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
		games_won serial,
		games_lost serial,
		games_played serial,
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
		color_to_move integer,
		game_fen varchar(500),
		game_type varchar(50),
		game_status varchar(50),
		game_ended boolean,
		game_outcome varchar(100),
		is_ranked boolean,
		time integer,
		time_added integer,
		white_time integer,
		black_time integer,
		moves_count integer,
		move_history varchar(1000),
		created_at timestamp
	)`

	_, err := s.db.Exec(query)

	return err
}

package storage

import (
	"database/sql"
	"fmt"

	"github.com/rekjef/openchess/internal/types"
)

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Nickname,
		&account.EncryptedPassword,
		&account.Elo,
		&account.GamesWon,
		&account.GamesLost,
		&account.GamesPlayed,
		&account.CreatedAt)
	return account, err
}

func (s *PostgressStore) CreateAccount(acc Account) (int, error) {
	query := `insert into account
	(
		first_name, last_name, nickname, encrypted_password, elo, 
		games_won, games_lost, games_played, created_at
	) 
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`

	row, err := s.db.Query(
		query,
		acc.FirstName,
		acc.LastName,
		acc.Nickname,
		acc.EncryptedPassword,
		acc.Elo,
		0, 0, 0,
		acc.CreatedAt,
	)

	if err != nil {
		return -1, err
	}

	for row.Next() {
		err = row.Scan(&acc.ID)

		return -1, err
	}

	return acc.ID, err
}

func (s *PostgressStore) UpdateAccount(*Account) error {
	return nil
}
func (s *PostgressStore) DeleteAccount(id int) error {
	_, err := s.db.Query("delete from account where id = $1", id)
	return err
}
func (s *PostgressStore) GetAccountByID(id int) (*Account, error) {
	rows, err := s.db.Query("select * from account where id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("user with ID=%d does not exist", id)
}

func (s *PostgressStore) GetAccountByNickname(nickname string) (*Account, error) {
	rows, err := s.db.Query("select * from account where nickname = $1", nickname)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("user with nickname=%s does not exist", nickname)
}

func (s *PostgressStore) GetAccounts() ([]Account, error) {
	rows, err := s.db.Query("select * from account")
	if err != nil {
		return nil, err
	}

	accounts := []Account{}
	for rows.Next() {
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, *account)
	}

	return accounts, nil
}

func (s *PostgressStore) UpdatePlayerStats(id int, stats types.UserStats) error {
	var updateQuery string

	if stats.GameLost == 1 {
		updateQuery = `
        UPDATE account 
        SET
            games_lost = games_lost + 1,
            games_played = games_played + 1
        WHERE
            id = $1`
	} else {
		updateQuery = `
        UPDATE account 
        SET
            games_won = games_won + 1,
            games_played = games_played + 1
        WHERE
            id = $1`
	}

	_, updateErr := s.db.Exec(updateQuery, id)
	return updateErr
}

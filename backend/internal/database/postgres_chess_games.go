package database

import "fmt"

func (s *PostgressStore) GetChessGames() ([]*ChessGame, error) {
	rows, err := s.db.Query("select * from chessGame")
	if err != nil {
		return nil, err
	}

	games := []*ChessGame{}
	for rows.Next() {
		game, err := scanIntoGame(rows)
		if err != nil {
			return nil, err
		}
		games = append(games, game)
	}

	return games, nil
}

func (s *PostgressStore) GetChessGameByID(id int) (*ChessGame, error) {
	rows, err := s.db.Query("select * from chessGame where id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoGame(rows)
	}

	return nil, fmt.Errorf("game with ID=%d does not exist", id)
}

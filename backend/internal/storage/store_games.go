package storage

import (
	"database/sql"
	"fmt"
)

func scanIntoGame(rows *sql.Rows) (*ChessGame, error) {
	game := new(ChessGame)
	err := rows.Scan(
		&game.ID,
		&game.HostID,
		&game.WhitePlayerID,
		&game.BlackPlayerID,
		&game.WhiteToMove,
		&game.ColorToMove,
		&game.GameFEN,
		&game.GameType,
		&game.GameStatus,
		&game.GameEnded,
		&game.GameOutcome,
		&game.IsRanked,
		&game.Time,
		&game.TimeAdded,
		&game.WhiteTime,
		&game.BlackTime,
		&game.MovesCount,
		&game.MoveHistory,
		&game.CreatedAt,
	)

	return game, err
}

func (s *PostgressStore) CreateChessGame(chessGame ChessGame) (int, error) {
	query := `insert into chessGame
	(
		host_id,
		white_player_id,
		black_player_id,
		white_to_move,
		color_to_move,
		game_fen,
		game_type,
		game_status,
		game_ended,
		game_outcome,
		is_ranked,
		time,
		time_added,
		white_time,
		black_time,
		moves_count,
		move_history,
		created_at
	) 
	values (
		$1, $2, $3, $4, $5, $6, $7, $8, 
		$9, $10, $11, $12, $13, $14, $15,
		$16, $17, $18 
	) RETURNING id`

	row, err := s.db.Query(
		query,
		chessGame.HostID,
		chessGame.WhitePlayerID,
		chessGame.BlackPlayerID,
		chessGame.WhiteToMove,
		chessGame.ColorToMove,
		chessGame.GameFEN,
		chessGame.GameType,
		chessGame.GameStatus,
		chessGame.GameEnded,
		chessGame.GameOutcome,
		chessGame.IsRanked,
		chessGame.Time,
		chessGame.TimeAdded,
		chessGame.Time,
		chessGame.Time,
		chessGame.MovesCount,
		chessGame.MoveHistory,
		chessGame.CreatedAt,
	)

	if err != nil {
		return -1, err
	}

	for row.Next() {
		if err := row.Scan(&chessGame.ID); err != nil {
			return -1, err
		}
	}

	return chessGame.ID, err
}

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

func (s *PostgressStore) UpdateChessGame(updatedGame ChessGame) error {
	updateQuery := `
	UPDATE chessGame 
	SET
		host_id = $1,
		white_player_id = $2,
		black_player_id = $3,
		white_to_move = $4,
		color_to_move = $5,
		game_fen = $6,
		game_type = $7,
		game_status = $8,
		game_ended = $9,
		game_outcome = $10,
		is_ranked = $11,
		time = $12,
		time_added = $13,
		white_time = $14,
		black_time = $15,
		moves_count = $16,
		move_history = $17,
		created_at = $18
	WHERE
		id = $19`

	_, updateErr := s.db.Query(
		updateQuery,
		updatedGame.HostID,
		updatedGame.WhitePlayerID,
		updatedGame.BlackPlayerID,
		updatedGame.WhiteToMove,
		updatedGame.ColorToMove,
		updatedGame.GameFEN,
		updatedGame.GameType,
		updatedGame.GameStatus,
		updatedGame.GameEnded,
		updatedGame.GameOutcome,
		updatedGame.IsRanked,
		updatedGame.Time,
		updatedGame.TimeAdded,
		updatedGame.WhiteTime,
		updatedGame.BlackTime,
		updatedGame.MovesCount,
		updatedGame.MoveHistory,
		updatedGame.CreatedAt,
		updatedGame.ID,
	)

	return updateErr
}

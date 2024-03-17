package database

import (
	"fmt"

	"github.com/rekjef/openchess/internal/types"
)

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

func (s *PostgressStore) UpdateChessGame(updatedGame *types.ChessGame) error {
	updateQuery := `
	UPDATE chessGame 
	SET
		host_id = $1,
		white_player_id = $2,
		black_player_id = $3,
		white_to_move = $4,
		game_fen = $5,
		game_type = $6,
		game_status = $7,
		game_ended = $8,
		game_won_by_white = $9,
		is_ranked = $10,
		time = $11,
		time_added = $12,
		moves_count = $13,
		move_history = $14,
		created_at = $15
	WHERE
		id = $16`

	_, updateErr := s.db.Query(
		updateQuery,
		updatedGame.HostID,
		updatedGame.WhitePlayerID,
		updatedGame.BlackPlayerID,
		updatedGame.WhiteToMove,
		updatedGame.GameFEN,
		updatedGame.GameType,
		updatedGame.GameStatus,
		updatedGame.GameEnded,
		updatedGame.GameWonByWhite,
		updatedGame.IsRanked,
		updatedGame.Time,
		updatedGame.TimeAdded,
		updatedGame.MovesCount,
		updatedGame.MoveHistory,
		updatedGame.CreatedAt,
		updatedGame.ID,
	)

	return updateErr
}

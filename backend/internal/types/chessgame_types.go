package types

import (
	"time"

	"github.com/notnil/chess"
)

type CreateGameRequest struct {
	ID            int    `json:"id"`
	HostID        int    `json:"hostID"`
	WhitePlayerID int    `json:"whitePlayerID"`
	BlackPlayerID int    `json:"blackPlayerID"`
	IsRanked      bool   `json:"isRanked"`
	Time          int    `json:"time"`
	TimeAdded     int    `json:"timeAdded"`
	GameType      string `json:"gameType"`
}

type Status string

const (
	GameStarted Status = "Game started"
	GameEnded   Status = "Game ended"
)

type ChessGame struct {
	ID            int           `json:"id"`
	HostID        int           `json:"host_id"`
	WhitePlayerID int           `json:"white_player_id"`
	BlackPlayerID int           `json:"black_player_id"`
	WhiteToMove   bool          `json:"white_to_move"`
	ColorToMove   chess.Color   `json:"color_to_move"`
	GameFEN       string        `json:"game_fen"`
	GameType      string        `json:"game_type"`
	GameStatus    Status        `json:"game_status"`
	GameEnded     bool          `json:"game_ended"`
	GameOutcome   chess.Outcome `json:"game_outcome"`
	IsRanked      bool          `json:"is_ranked"`
	Time          int           `json:"time"`
	TimeAdded     int           `json:"time_added"`
	WhiteTime     int           `json:"white_time"`
	BlackTime     int           `json:"black_time"`
	MovesCount    int           `json:"moves_count"`
	MoveHistory   string        `json:"move_history"`
	CreatedAt     time.Time     `json:"created_at"`
}

func NewChessGame(req CreateGameRequest) *ChessGame {
	white := chess.White
	return &ChessGame{
		HostID:        req.HostID,
		WhitePlayerID: req.WhitePlayerID,
		BlackPlayerID: req.BlackPlayerID,
		WhiteToMove:   true,
		ColorToMove:   white,
		GameFEN:       "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		GameType:      req.GameType,
		GameStatus:    GameStarted,
		GameEnded:     false,
		GameOutcome:   chess.NoOutcome,
		IsRanked:      req.IsRanked,
		Time:          req.Time,
		TimeAdded:     req.TimeAdded,
		WhiteTime:     req.Time,
		BlackTime:     req.Time,
		MovesCount:    0,
		MoveHistory:   "",
		CreatedAt:     time.Now().UTC(),
	}
}

func (g *ChessGame) SwitchColors() {
	g.WhiteToMove = !g.WhiteToMove
	if g.WhiteToMove {
		g.ColorToMove = chess.White
	} else {
		g.ColorToMove = chess.Black
	}
}

package types

import (
	"math/rand"
	"time"

	"github.com/notnil/chess"
)

type CreateGameRequest struct {
	ID         int  `json:"id"`
	HostID     int  `json:"hostID"`
	OpponentID int  `json:"opponentID"`
	IsRanked   bool `json:"isRanked"`
	Time       int  `json:"time"`
	TimeAdded  int  `json:"timeAdded"`
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

func NewChessGame(req CreateGameRequest) ChessGame {
	var gameType string
	if req.IsRanked {
		gameType = "Ranked"
	} else {
		gameType = "Unranked"
	}
	var WhitePlayerID, BlackPlayerID int = req.HostID, req.OpponentID
	if rand.Intn(2) == 1 {
		WhitePlayerID, BlackPlayerID = BlackPlayerID, WhitePlayerID
	}

	return ChessGame{
		HostID:        req.HostID,
		WhitePlayerID: WhitePlayerID,
		BlackPlayerID: BlackPlayerID,
		WhiteToMove:   true,
		ColorToMove:   chess.White,
		GameFEN:       "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		GameType:      gameType,
		GameStatus:    GameStarted,
		GameEnded:     false,
		GameOutcome:   chess.NoOutcome,
		IsRanked:      req.IsRanked,
		Time:          req.Time * 1000, // req.Time is in seconds, we need miliseconds
		TimeAdded:     req.TimeAdded * 1000,
		WhiteTime:     req.Time * 1000,
		BlackTime:     req.Time * 1000,
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

func (g *ChessGame) AddToHistory(move string) {
	if g.MoveHistory == "" {
		g.MoveHistory = move
	} else {
		g.MoveHistory += " " + move
	}
}

func (g ChessGame) UserIDToMove() int {
	if g.WhiteToMove {
		return g.WhitePlayerID
	}
	return g.BlackPlayerID
}

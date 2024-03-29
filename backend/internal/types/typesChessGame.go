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
	Started Status = "Game started"
	Ended   Status = "Game ended"
)

type ChessGame struct {
	ID             int       `json:"id"`
	HostID         int       `json:"hostID"`
	WhitePlayerID  int       `json:"whitePlayerID"`
	BlackPlayerID  int       `json:"blackPlayerID"`
	WhiteToMove    bool      `json:"whiteToMove"`
	GameFEN        string    `json:"gameFEN"`
	GameType       string    `json:"gameType"`
	GameStatus     string    `json:"gameStatus"`
	GameEnded      bool      `json:"gameEnded"`
	GameWonByWhite bool      `json:"gameWonByWhite"`
	IsRanked       bool      `json:"isRanked"`
	Time           int       `json:"time"`
	TimeAdded      int       `json:"timeAdded"`
	MovesCount     int       `json:"movesCount"`
	MoveHistory    string    `json:"moveHistory"`
	CreatedAt      time.Time `json:"createdAt"`
}

func NewChessGame(req CreateGameRequest) *ChessGame {
	return &ChessGame{
		HostID:         req.HostID,
		WhitePlayerID:  req.WhitePlayerID,
		BlackPlayerID:  req.BlackPlayerID,
		WhiteToMove:    true,
		GameFEN:        "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		GameType:       req.GameType,
		GameStatus:     string(Started),
		GameEnded:      false,
		GameWonByWhite: false,
		IsRanked:       req.IsRanked,
		Time:           req.Time,
		TimeAdded:      req.TimeAdded,
		MovesCount:     0,
		MoveHistory:    "",
		CreatedAt:      time.Now().UTC(),
	}
}

type LiveGame struct {
	Engine  *chess.Game
	Details ChessGame
}

type GameUpdateOptions struct {
	Move      string `json:"move"`
	Resign    bool   `json:"resign"`
	OfferDraw bool   `json:"draw"`
}

func (c *ChessGame) createEngineFromFEN(fenStr string) (*chess.Game, error) {
	fen, err := chess.FEN(fenStr)
	if err != nil {
		return nil, err
	}

	return chess.NewGame(fen), nil
}

func (c *ChessGame) Move(move string) error {
	engine, err := c.createEngineFromFEN(c.GameFEN)
	if err != nil {
		return err
	}

	if err := engine.MoveStr(move); err != nil {
		return err
	}

	c.WhiteToMove = !c.WhiteToMove
	c.GameFEN = engine.FEN()
	c.MovesCount++
	return nil
}

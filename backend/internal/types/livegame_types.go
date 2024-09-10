package types

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/notnil/chess"
)

type LiveGame struct {
	Engine         *chess.Game      `json:"-"`
	ColorChannel   chan chess.Color `json:"-"`
	EndGameChannel chan bool        `json:"-"`
	Details        ChessGame        `json:"details"`
}

type GameUpdateOptions struct {
	Move      string `json:"move"`
	Resign    bool   `json:"resign"`
	UserID    int    `json:"user_id"`
	OfferDraw bool   `json:"draw"`
}

type UpdatedState struct {
	FEN         string `json:"fen"`
	MoveHistory string `json:"move_history"`
	WhiteToMove bool   `json:"white_to_move"`
	WhiteTime   int    `json:"white_time"`
	BlackTime   int    `json:"black_time"`
}

func NewLiveGame(details ChessGame) *LiveGame {
	return &LiveGame{
		Engine:         chess.NewGame(),
		ColorChannel:   make(chan chess.Color, 1),
		EndGameChannel: make(chan bool, 1),
		Details:        details,
	}
}

func (g *LiveGame) StartGame() {
	ticker := time.NewTicker(100 * time.Millisecond)
	// start decrementing white time
	targetTime := &g.Details.WhiteTime
	fmt.Println(*targetTime)

	for {
		select {
		// switch targetTime after move
		case c := <-g.ColorChannel:
			if c == chess.White {
				targetTime = &g.Details.WhiteTime
			} else {
				targetTime = &g.Details.BlackTime
			}
		// timer
		case <-ticker.C:
			*targetTime -= 100
			if *targetTime <= 0 {
				return
			}
		// check for draw or surrender
		case end := <-g.EndGameChannel:
			if end {
				return
			}
		}
	}
}

func (g *LiveGame) CommitMove(move *chess.Move) error {
	err := g.Engine.Move(move)
	if err != nil {
		return err
	}
	g.Details.SwitchColors()

	g.Details.GameOutcome = g.Engine.Outcome()
	g.Details.AddToHistory(move.String())
	g.Details.MovesCount++
	g.Details.GameFEN = g.Engine.FEN()

	return nil
}

func (g *LiveGame) ComputeAIMove() *chess.Move {
	moves := g.Engine.ValidMoves()
	move := moves[rand.Int()%len(moves)]
	return move
}

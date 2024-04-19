package types

import (
	"time"

	"github.com/notnil/chess"
)

type LiveGame struct {
	Engine         *chess.Game      `json:"-"`
	ColorChannel   chan chess.Color `json:"-"`
	EndGameChannel chan bool        `json:"-"`
	Details        ChessGame        `json:"details"`
}

func NewLiveGame(details ChessGame) *LiveGame {
	return &LiveGame{
		Engine:         chess.NewGame(),
		ColorChannel:   make(chan chess.Color, 1),
		EndGameChannel: make(chan bool, 1),
		Details:        details,
	}
}

type GameUpdateOptions struct {
	Move      string `json:"move"`
	Resign    bool   `json:"resign"`
	OfferDraw bool   `json:"draw"`
}

func (g *LiveGame) StartGame() {
	ticker := time.NewTicker(100 * time.Millisecond)
	// start decrementing white time
	targetTime := &g.Details.WhiteTime

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

package database

import (
	"fmt"

	"github.com/notnil/chess"
	"github.com/rekjef/openchess/internal/types"
)

type RAMStore struct {
	games map[int]types.LiveGame
}

func NewRAMStore() *RAMStore {
	return &RAMStore{
		games: make(map[int]types.LiveGame, 0),
	}
}

// Add new game to memory
func (r *RAMStore) AddGame(details types.ChessGame) error {
	if _, ok := r.games[details.ID]; ok {
		return fmt.Errorf("game with ID %d already exists", details.ID)
	}

	newGame := types.LiveGame{Engine: chess.NewGame(), Details: details}
	r.games[details.ID] = newGame
	return nil
}

// Get status from memory
func (r *RAMStore) GetGame(id int) (*types.LiveGame, error) {
	if liveGame, ok := r.games[id]; ok {
		return &liveGame, nil
	}
	return nil, fmt.Errorf("there is no active game with ID %d", id)
}

// Update live state
func (r *RAMStore) UpdateGame(id int, options types.GameUpdateOptions) error {
	liveGame, ok := r.games[id]
	if !ok {
		return fmt.Errorf("there is no active game with ID %d", id)
	}

	if options.Resign {
		liveGame.Details.GameEnded = true
	} else if options.OfferDraw {
		liveGame.Details.GameStatus = "Draw"
	} else if options.Move != "" {
		liveGame.Details.MoveHistory = options.Move
	}
	r.games[id] = liveGame
	fmt.Println(r.games[id].Details.MoveHistory)
	return nil
}

// Delete game
func (r *RAMStore) DeleteGame(id int) error {
	if _, ok := r.games[id]; ok {
		delete(r.games, id)
		return nil
	}
	return fmt.Errorf("there is no active game with ID %d", id)
}

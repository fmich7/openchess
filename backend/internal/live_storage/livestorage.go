package live_storage

import (
	"errors"
	"fmt"
	"log"

	"github.com/notnil/chess"
	"github.com/rekjef/openchess/internal/types"
	"github.com/rekjef/openchess/internal/utils"
)

type Account = types.Account

type RAMStore map[int]*types.LiveGame

func NewRAMStore() *RAMStore {
	return &RAMStore{}
}

// Add new game to memory
func (r RAMStore) AddGame(details types.ChessGame, store types.Storage) error {
	id := details.ID
	if _, ok := r[id]; ok {
		return fmt.Errorf("game with ID %d already exists", id)
	}

	r[id] = types.NewLiveGame(details)
	if r[id] == nil {
		return fmt.Errorf("failed to create new LiveGame")
	}

	go func() {
		// make first move if computer is white
		if details.HostID != details.UserIDToMove() {
			r.MakeMove(id, r[id].ComputeAIMove())
		}
		if r[id] != nil {
			r[id].StartGame()
			defer r.EndGame(id, chess.BlackWon, store)
		} else {
			log.Println("LiveGame is nil, cannot start game")
		}
	}()

	return nil
}

// Get status from memory
func (r RAMStore) GetGame(id int) (*types.LiveGame, error) {
	if liveGame, ok := r[id]; ok {
		return liveGame, nil
	}
	return nil, utils.NoActiveGameError(id)
}

func (r RAMStore) MakeMove(id int, move *chess.Move) error {
	liveGame, ok := r[id]
	if !ok {
		return utils.NoActiveGameError(id)
	}
	details := &liveGame.Details

	if err := liveGame.CommitMove(move); err != nil {
		return err
	}
	r[id].ColorChannel <- details.ColorToMove

	return nil
}

// Update live state
func (r RAMStore) UpdateGame(id int, options types.GameUpdateOptions, store types.Storage) error {
	liveGame, ok := r[id]
	if !ok {
		return utils.NoActiveGameError(id)
	}

	details := &liveGame.Details
	engine := liveGame.Engine
	// fmt.Printf("%+v\n", engine.ValidMoves())
	if options.Resign {
		var color chess.Color
		if details.WhiteToMove {
			color = chess.White
		} else {
			color = chess.Black
		}

		engine.Resign(color)
	} else if options.OfferDraw {
		engine.Draw(chess.DrawOffer)
	} else if options.Move != "" {
		// check if move is invalid
		if options.UserID != details.UserIDToMove() {
			return errors.New("your id doesn't match with allowed one to move")
		}

		var move *chess.Move
		for _, elem := range engine.ValidMoves() {
			if options.Move == elem.String() {
				move = elem
			}
		}

		if err := r.MakeMove(id, move); err != nil {
			return err
		}
	}

	r[id] = liveGame
	// check if game has ended
	// *, 1-0, 0-1, 1/2-1/2
	engineOutcome := engine.Outcome()
	if engineOutcome != chess.NoOutcome {
		return r.EndGame(id, engineOutcome, store)
	}

	return nil
}

// Delete game
func (r RAMStore) DeleteGame(id int) error {
	if _, ok := r[id]; ok {
		delete(r, id)
		return nil
	}
	return utils.NoActiveGameError(id)
}

// Delete game from ram mem, add to db
func (r RAMStore) EndGame(id int, outcome chess.Outcome, store types.Storage) error {
	game, ok := r[id]
	if !ok {
		return utils.NoActiveGameError(id)
	}

	// modify game state
	game.EndGameChannel <- true
	details := game.Details
	details.GameEnded = true
	details.GameStatus = types.GameEnded
	details.GameOutcome = outcome

	if err := r.DeleteGame(id); err != nil {
		return err
	}

	// check whenever add new game to db or to update if exists
	if _, err := store.GetChessGameByID(details.ID); err != nil {
		_, err := store.CreateChessGame(details)
		return err
	}

	return store.UpdateChessGame(details)
}

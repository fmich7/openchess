package account

import (
	"github.com/rekjef/openchess/internal/types"
)

type Profile struct {
	Acc   Account
	Games []types.ChessGame
}

func GetProfile(id int, storage types.Storage) (Profile, error) {
	acc, err := storage.GetAccountByID(id)
	// no acc
	if err != nil {
		return Profile{}, err
	}
	allGames, err := storage.GetChessGames()
	if err != nil {
		return Profile{Acc: *acc, Games: make([]types.ChessGame, 0)}, nil
	}

	userGames := []types.ChessGame{}
	for _, game := range allGames {
		if game.WhitePlayerID == id || game.BlackPlayerID == id {
			userGames = append(userGames, *game)
		}
	}

	return Profile{Acc: *acc, Games: userGames}, nil
}

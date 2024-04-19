package account

import (
	"math/rand"
	"time"

	"github.com/rekjef/openchess/internal/types"
	"golang.org/x/crypto/bcrypt"
)

type Account = types.Account

func NewAccount(firstName, lastName, nickname, password string) (*Account, error) {
	encyptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Account{
		FirstName:         firstName,
		LastName:          lastName,
		Nickname:          nickname,
		EncryptedPassword: string(encyptedPassword),
		Elo:               rand.Intn(2500),
		CreatedAt:         time.Now().UTC(),
	}, nil
}

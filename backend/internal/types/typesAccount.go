package types

import (
	"math/rand"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthClaims struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	jwt.RegisteredClaims
}

type UserAuthInfo struct {
	ID int `json:"id"`
}

type Account struct {
	ID                int       `json:"id"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	Nickname          string    `json:"nickname"`
	EncryptedPassword string    `json:"-"`
	Elo               int       `json:"elo"`
	CreatedAt         time.Time `json:"createdAt"`
}

type LoginResponse struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
}

type LoginRequest struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Nickname  string `json:"nickname"`
	Password  string `json:"password"`
}

func (a *Account) ComparePasswords(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(pw)) == nil
}

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

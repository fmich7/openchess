package types

import (
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

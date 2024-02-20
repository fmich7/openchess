package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Nickname  string `json:"nickname"`
}

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Nickname  string    `json:"nickname"`
	Elo       int       `json:"elo"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewAccount(firstName, lastName, nickname string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Nickname:  nickname,
		Elo:       rand.Intn(2500),
		CreatedAt: time.Now().UTC(),
	}
}

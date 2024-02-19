package main

import "math/rand"

type Account struct {
	ID        int `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Nickname  string `json:"nickname"`
	Elo       int `json:"elo"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID: rand.Intn(10000),
		FirstName: firstName,
		LastName: lastName,
		Nickname: "carlsen",
		Elo: rand.Intn(2500),
	}
}
package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	acc, err := NewAccount("Bob", "Ross", "br11", "maslo")
	assert.Nil(t, err)
	fmt.Printf("%+v\n", acc)
}

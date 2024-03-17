package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	password := "maslo"
	acc, err := NewAccount("Bob", "Ross", "br11", password)
	assert.Nil(t, err)

	if password == acc.EncryptedPassword {
		t.Fatalf("password has not been hashed")
	} else if acc.ComparePasswords(password) == false {
		t.Fatalf("password hashes are not same")
	}
}

package account

import (
	"testing"

	"github.com/rekjef/openchess/internal/tests"
	"github.com/rekjef/openchess/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestGetProfile(t *testing.T) {
	storage := tests.NewFakeDB()

	t.Run("empty db", func(t *testing.T) {
		if _, err := GetProfile(0, storage); err == nil {
			t.Error("should raise an error on empty db")
		}
	})

	storage.CreateAccount(types.Account{ID: 0})

	t.Run("acc exists", func(t *testing.T) {
		profile, err := GetProfile(0, storage)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, profile.Acc.ID, 0)
	})
}

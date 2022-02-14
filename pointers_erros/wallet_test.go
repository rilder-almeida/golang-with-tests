package pointerserros

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit from wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		got, err := wallet.GetBalance()
		want := Bitcoin(10)

		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("withdraw from wallet", func(t *testing.T) {
		wallet := Wallet{Balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		got, err := wallet.GetBalance()
		want := Bitcoin(10)

		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("insufficient founds", func(t *testing.T) {
		wallet := Wallet{Balance: Bitcoin(20)}
		got := wallet.Withdraw(Bitcoin(100))

		assert.Error(t, got)
	})

}

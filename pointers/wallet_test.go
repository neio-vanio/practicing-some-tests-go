package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	verifyBalance := func(t *testing.T, wallet Wallet, expected Bitcoin) {
		t.Helper()
		result := wallet.Balance()

		if result != expected {
			t.Errorf("result %.2f, expected %.2f", result, expected)
		}
	}

	verifyError := func(t *testing.T, err error) {
		t.Helper()
		if err == nil {
			t.Error("expected an error, but none occurred")
		}
	}

	t.Run("Balance", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		expected := Bitcoin(10.00000000)
		verifyBalance(t, wallet, expected)
	})

	t.Run("Balance in string format", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		expected := Bitcoin(10.000000)
		verifyBalance(t, wallet, expected)
	})

	t.Run("Withdrawal", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(8.89512410)}

		wallet.Withdrawal(Bitcoin(0.78500000))
		expected := Bitcoin(8.1101241)
		verifyBalance(t, wallet, expected)
	})

	t.Run("Withdrawal with insufficient balance", func(t *testing.T) {
		startBalance := Bitcoin(1.80003200)
		wallet := Wallet{startBalance}
		err := wallet.Withdrawal(Bitcoin(57.980000))
		verifyError(t, err)
	})
}

package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	// helper function
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		if wallet.Balance != want {
			t.Errorf("Expected %s but got %s", want, wallet.Balance)
		}
	}

	asserError := func(t *testing.T, err error, err_msg_want string) {
		t.Helper()
		if err == nil {
			t.Fatal("Expected an error, but didn't get any!")
		}
		if err.Error() != err_msg_want {
			t.Errorf("Didn't got the error message as we wanted, we wanted %s but got %s", err_msg_want, err)
		}
	}

	assertNotError := func(t *testing.T, err error) {
		t.Helper()
		if err != nil {
			t.Fatalf("There shouldn't be any error, but go some error %v", err)
		}
	}

	// test cases
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{Balance: 0}
		wallet.Deposit(10)
		assertBalance(t, wallet, 10)
	})

	t.Run("withdraw (when with sufficient fund)", func(t *testing.T) {
		wallet := Wallet{Balance: 10}
		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, 0)
		assertNotError(t, err)
	})

	t.Run("withdraw (when with in-sufficient fund)", func(t *testing.T) {
		wallet := Wallet{Balance: 10}
		err := wallet.Withdraw(11) // it should return err in this scenario so catching it
		assertBalance(t, wallet, 10)
		asserError(t, err, InsufficientFundsError.Error())
	})

}

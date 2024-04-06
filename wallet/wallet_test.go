package wallet_test

import (
	"sync"
	"testing"

	"github.com/salawhaaat/bcwallet/wallet"
)

func TestConcurrentOperations(t *testing.T) {
	t.Run("MultipleDeposits", func(t *testing.T) {
		testMultipleDeposits(t)
	})

	t.Run("MultipleWithdrawals", func(t *testing.T) {
		testMultipleWithdrawals(t)
	})

	t.Run("MultipleDepositsAndWithdrawals", func(t *testing.T) {
		testMultipleDepositsAndWithdrawals(t)
	})
}

func testMultipleDeposits(t *testing.T) {
	w := wallet.New(wallet.Bitcoin(100))

	numDeposits := 1000
	depositAmount := wallet.Bitcoin(10)

	var wg sync.WaitGroup
	wg.Add(numDeposits)

	for i := 0; i < numDeposits; i++ {
		go func() {
			defer wg.Done()
			w.Deposit(depositAmount)
		}()
	}

	wg.Wait()

	finalBalance := w.Balance()
	expectedBalance := wallet.Bitcoin(100 + wallet.Bitcoin(numDeposits)*depositAmount)
	assertBalance(t, finalBalance, expectedBalance)
}

func testMultipleWithdrawals(t *testing.T) {
	w := wallet.New(wallet.Bitcoin(1000))

	numWithdrawals := 500
	withdrawalAmount := wallet.Bitcoin(2)

	var wg sync.WaitGroup
	wg.Add(numWithdrawals)

	for i := 0; i < numWithdrawals; i++ {
		go func() {
			defer wg.Done()
			w.Withdraw(withdrawalAmount)
		}()
	}

	wg.Wait()

	finalBalance := w.Balance()
	expectedBalance := wallet.Bitcoin(1000 - wallet.Bitcoin(numWithdrawals)*withdrawalAmount)
	assertBalance(t, finalBalance, expectedBalance)
}

func testMultipleDepositsAndWithdrawals(t *testing.T) {
	w := wallet.New(wallet.Bitcoin(100))

	numOperations := 1000
	depositAmount := wallet.Bitcoin(10)
	withdrawalAmount := wallet.Bitcoin(5)

	var wg sync.WaitGroup
	wg.Add(numOperations)

	for i := 0; i < numOperations; i++ {
		go func() {
			defer wg.Done()
			w.Deposit(depositAmount)
			w.Withdraw(withdrawalAmount)
		}()
	}

	wg.Wait()

	finalBalance := w.Balance()
	expectedBalance := wallet.Bitcoin(100 + wallet.Bitcoin(numOperations)*depositAmount - wallet.Bitcoin(numOperations)*withdrawalAmount)
	assertBalance(t, finalBalance, expectedBalance)
}

func assertBalance(t *testing.T, got, want wallet.Bitcoin) {
	t.Helper()
	if got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}

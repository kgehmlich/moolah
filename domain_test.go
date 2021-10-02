package main

import "testing"

func TestAccount(t *testing.T) {
	t.Run("Balance", func(t *testing.T) {
		t.Run("returns balance", func(t *testing.T) {
			expectedBalance := Money(123.45)
			acct := Account{balance: expectedBalance}
			returnedBalance := acct.Balance()
			if returnedBalance != expectedBalance {
				t.Fatalf("Expected %f, got %f", expectedBalance, returnedBalance)
			}
		})
	})

	t.Run("Credit", func(t *testing.T) {
		t.Run("with positive amount succeeds", func(t *testing.T) {
			expectedBalance := Money(123.45)
			acct := Account{balance: 0}
			err := acct.Credit(expectedBalance)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if acct.Balance() != expectedBalance {
				t.Fatalf("Expected %f, got %f", expectedBalance, acct.Balance())
			}
		})

		t.Run("with zero amount returns error", func(t *testing.T) {
			acct := Account{balance: 100}
			err := acct.Credit(0)
			if err == nil {
				t.Fatalf("Expected error, got nil")
			}
			if err != ErrNonPositiveAmount {
				t.Fatalf("Unexpected error: %s", err)
			}
			if acct.Balance() != 100 {
				t.Fatalf("Account balance changed")
			}
		})

		t.Run("with negative amount returns error", func(t *testing.T) {
			acct := Account{balance: 100}
			err := acct.Credit(-1)
			if err == nil {
				t.Fatalf("Expected error, got nil")
			}
			if err != ErrNonPositiveAmount {
				t.Fatalf("Unexpected error: %s", err)
			}
			if acct.Balance() != 100 {
				t.Fatalf("Account balance changed")
			}
		})
	})

	t.Run("Debit", func(t *testing.T) {
		t.Run("with positive amount succeeds", func(t *testing.T) {
			startingBalance := Money(123.45)
			debitAmount := Money(23.45)
			acct := Account{balance: startingBalance}
			err := acct.Debit(debitAmount)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			expectedBalance := startingBalance - debitAmount
			if acct.Balance() != expectedBalance {
				t.Fatalf("Expected %f, got %f", expectedBalance, acct.Balance())
			}
		})

		t.Run("with amount equal to balance succeeds", func(t *testing.T) {
			startingBalance := Money(123.45)
			debitAmount := startingBalance
			acct := Account{balance: startingBalance}
			err := acct.Debit(debitAmount)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			expectedBalance := Money(0)
			if acct.Balance() != expectedBalance {
				t.Fatalf("Expected %f, got %f", expectedBalance, acct.Balance())
			}
		})

		t.Run("with zero amount returns error", func(t *testing.T) {
			acct := Account{balance: 100}
			err := acct.Debit(0)
			if err == nil {
				t.Fatalf("Expected error, got nil")
			}
			if err != ErrNonPositiveAmount {
				t.Fatalf("Unexpected error: %s", err)
			}
			if acct.Balance() != 100 {
				t.Fatalf("Account balance changed")
			}
		})

		t.Run("with negative amount returns error", func(t *testing.T) {
			acct := Account{balance: 100}
			err := acct.Debit(-1)
			if err == nil {
				t.Fatalf("Expected error, got nil")
			}
			if err != ErrNonPositiveAmount {
				t.Fatalf("Unexpected error: %s", err)
			}
			if acct.Balance() != 100 {
				t.Fatalf("Account balance changed")
			}
		})

		t.Run("with amount greater than balance returns error", func(t *testing.T) {
			startingBalance := Money(123.45)
			debitAmount := Money(1000)
			acct := Account{balance: startingBalance}
			err := acct.Debit(debitAmount)
			if err == nil {
				t.Fatalf("Expected error, got nil")
			}
			if err != ErrInsufficientFunds {
				t.Fatalf("Unexpected error: %s", err)
			}
			if acct.Balance() != startingBalance {
				t.Fatalf("Account balance changed")
			}
		})
	})
}

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

func TestCategory(t *testing.T) {
	t.Run("has name", func(t *testing.T) {
		expectedName := "test name"
		cat := Category{Name: expectedName}
		if cat.Name != expectedName {
			t.Fatalf("Expected %s, got %s", expectedName, cat.Name)
		}
	})

	t.Run("Available", func(t *testing.T) {
		t.Run("returns available amount", func(t *testing.T) {
			expectedAvailable := Money(123.45)
			cat := Category{available: expectedAvailable}
			returnedAvailable := cat.Available()
			if returnedAvailable != expectedAvailable {
				t.Fatalf("Expected %f, got %f", expectedAvailable, returnedAvailable)
			}
		})
	})

	t.Run("Assign", func(t *testing.T) {
		t.Run("increases available by amount", func(t *testing.T) {
			startingAvailable := Money(123.45)
			assignAmount := Money(100)
			cat := Category{available: startingAvailable}
			err := cat.Assign(assignAmount)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			expectedAvailable := startingAvailable + assignAmount
			if cat.Available() != expectedAvailable {
				t.Fatalf("Expected %f, got %f", expectedAvailable, cat.Available())
			}
		})

		t.Run("with zero amount returns error", func(t *testing.T) {
			cat := Category{available: 100}
			err := cat.Assign(0)
			if err == nil {
				t.Fatalf("Expected error, got nil")
			}
			if err != ErrNonPositiveAmount {
				t.Fatalf("Unexpected error: %s", err)
			}
			if cat.Available() != 100 {
				t.Fatalf("Available amount changed")
			}
		})

		t.Run("with negative amount returns error", func(t *testing.T) {
			cat := Category{available: 100}
			err := cat.Assign(-1)
			if err == nil {
				t.Fatalf("Expected error, got nil")
			}
			if err != ErrNonPositiveAmount {
				t.Fatalf("Unexpected error: %s", err)
			}
			if cat.Available() != 100 {
				t.Fatalf("Available amount changed")
			}
		})
	})

	t.Run("Unassign", func(t *testing.T) {
		t.Run("decreases available by amount", func(t *testing.T) {
			startingAvailable := Money(123.45)
			unassignAmount := Money(200)
			cat := Category{available: startingAvailable}
			err := cat.Unassign(unassignAmount)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			expectedAvailable := startingAvailable - unassignAmount
			if cat.Available() != expectedAvailable {
				t.Fatalf("Expected %f, got %f", expectedAvailable, cat.Available())
			}
		})

		t.Run("with zero amount returns error", func(t *testing.T) {
			cat := Category{available: 100}
			err := cat.Unassign(0)
			if err == nil {
				t.Fatalf("Expected error, got nil")
			}
			if err != ErrNonPositiveAmount {
				t.Fatalf("Unexpected error: %s", err)
			}
			if cat.Available() != 100 {
				t.Fatalf("Available amount changed")
			}
		})

		t.Run("with negative amount returns error", func(t *testing.T) {
			cat := Category{available: 100}
			err := cat.Unassign(-1)
			if err == nil {
				t.Fatalf("Expected error, got nil")
			}
			if err != ErrNonPositiveAmount {
				t.Fatalf("Unexpected error: %s", err)
			}
			if cat.Available() != 100 {
				t.Fatalf("Available amount changed")
			}
		})
	})
}

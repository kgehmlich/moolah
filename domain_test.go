package main

import (
	"math"
	"testing"
)

func moneyEqual(a, b Money) bool {
	return math.Abs(float64(a-b)) < 1e-10
}

func TestAccount(t *testing.T) {
	t.Run("ID", func(t *testing.T) {
		t.Run("returns the account's permanent ID", func(t *testing.T) {
			expectedID := UniqueID("test_id")
			acct := Account{id: expectedID}
			returnedID := acct.ID()
			if returnedID != expectedID {
				t.Fatalf("Expected %s, got %s", expectedID, returnedID)
			}
		})
	})

	t.Run("Balance", func(t *testing.T) {
		t.Run("returns balance", func(t *testing.T) {
			expectedBalance := Money(123.45)
			acct := Account{balance: expectedBalance}
			returnedBalance := acct.Balance()
			if !moneyEqual(returnedBalance, expectedBalance) {
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
			if !moneyEqual(acct.Balance(), expectedBalance) {
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
			if !moneyEqual(acct.Balance(), 100) {
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
			if !moneyEqual(acct.Balance(), 100) {
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
			if !moneyEqual(acct.Balance(), expectedBalance) {
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
			if !moneyEqual(acct.Balance(), expectedBalance) {
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
			if !moneyEqual(acct.Balance(), 100) {
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
			if !moneyEqual(acct.Balance(), 100) {
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
			if !moneyEqual(acct.Balance(), startingBalance) {
				t.Fatalf("Account balance changed")
			}
		})
	})
}

func TestCategory(t *testing.T) {
	t.Run("ID", func(t *testing.T) {
		t.Run("returns the category's permanent ID", func(t *testing.T) {
			expectedID := UniqueID("test_id")
			cat := Category{id: expectedID}
			returnedID := cat.ID()
			if returnedID != expectedID {
				t.Fatalf("Expected %s, got %s", expectedID, returnedID)
			}
		})
	})

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
			if !moneyEqual(returnedAvailable, expectedAvailable) {
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
			if !moneyEqual(cat.Available(), expectedAvailable) {
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
			if !moneyEqual(cat.Available(), 100) {
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
			if !moneyEqual(cat.Available(), 100) {
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
			if !moneyEqual(cat.Available(), expectedAvailable) {
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
			if !moneyEqual(cat.Available(), 100) {
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
			if !moneyEqual(cat.Available(), 100) {
				t.Fatalf("Available amount changed")
			}
		})
	})
}

func TestBudget(t *testing.T) {
	t.Run("AddAccount", func(t *testing.T) {
		t.Run("creates new account", func(t *testing.T) {
			acctName := "test account"
			budget := Budget{}
			err := budget.AddAccount(acctName)
			if err != nil {
				t.Fatalf("Unespected error: %s", err)
			}
			if count := len(budget.Accounts()); count != 1 {
				t.Fatalf("Unexpected number of accounts. Expected %d, got %d", 1, count)
			}
			if n := budget.Accounts()[0].Name; n != acctName {
				t.Fatalf("Unexpected account name. Expected %s, got %s", acctName, n)
			}
		})

		t.Run("assigns a unique ID to the new account", func(t *testing.T) {
			acctName := "test account"
			budget := Budget{}
			err := budget.AddAccount(acctName)
			if err != nil {
				t.Fatalf("Unespected error: %s", err)
			}
			if id := budget.Accounts()[0].ID(); id == "" {
				t.Fatalf("Expected non-empty account ID")
			}
		})

		t.Run("returns error if name already exists", func(t *testing.T) {
			existingAcctName := "existing account"
			existingAcct := &Account{Name: existingAcctName}
			budget := Budget{
				accounts: []*Account{existingAcct},
			}

			err := budget.AddAccount(existingAcctName)
			if err == nil {
				t.Fatalf("Expected error, got nil")
			}
			if err != ErrDuplicateName {
				t.Fatalf("Unexpected error: %s", err)
			}
			if count := len(budget.Accounts()); count != 1 {
				t.Fatalf("Unexpected number of accounts. Expected %d, got %d", 1, count)
			}
		})
	})

	t.Run("TotalFunds", func(t *testing.T) {
		t.Run("returns total amount of money across all accounts", func(t *testing.T) {
			accts := []*Account{
				{balance: 111.11},
				{balance: 222.22},
			}
			expectedFunds := Money(333.33)

			budget := Budget{accounts: accts}

			if f := budget.TotalFunds(); !moneyEqual(f, expectedFunds) {
				t.Fatalf("Expected %f, got %f", expectedFunds, f)
			}
		})
	})

	t.Run("AddCategory", func(t *testing.T) {
		t.Run("creates new category", func(t *testing.T) {
			catName := "test category"
			budget := Budget{}
			err := budget.AddCategory(catName)
			if err != nil {
				t.Fatalf("Unespected error: %s", err)
			}
			if count := len(budget.Categories()); count != 1 {
				t.Fatalf("Unexpected number of categories. Expected %d, got %d", 1, count)
			}
			if n := budget.Categories()[0].Name; n != catName {
				t.Fatalf("Unexpected category name. Expected %s, got %s", catName, n)
			}
		})

		t.Run("assigns a unique ID to the new category", func(t *testing.T) {
			categoryName := "test category"
			budget := Budget{}
			err := budget.AddCategory(categoryName)
			if err != nil {
				t.Fatalf("Unespected error: %s", err)
			}
			if id := budget.Categories()[0].ID(); id == "" {
				t.Fatalf("Expected non-empty category ID")
			}
		})

		t.Run("returns error if name already exists", func(t *testing.T) {
			existingCategoryName := "existing account"
			existingCategory := &Category{Name: existingCategoryName}
			budget := Budget{
				categories: []*Category{existingCategory},
			}

			err := budget.AddCategory(existingCategoryName)
			if err == nil {
				t.Fatalf("Expected error, got nil")
			}
			if err != ErrDuplicateName {
				t.Fatalf("Unexpected error: %s", err)
			}
			if count := len(budget.Categories()); count != 1 {
				t.Fatalf("Unexpected number of categories. Expected %d, got %d", 1, count)
			}
		})
	})

	t.Run("UnassignedFunds", func(t *testing.T) {
		t.Run("returns unassigned funds", func(t *testing.T) {
			accts := []*Account{
				{balance: 111.11},
				{balance: 222.22},
			}
			categories := []*Category{
				{available: 100},
				{available: 200},
			}
			budget := Budget{
				accounts:   accts,
				categories: categories,
			}
			expectedUnassigned := Money(33.33)

			availableFunds := budget.UnassignedFunds()

			if !moneyEqual(availableFunds, expectedUnassigned) {
				t.Fatalf("Expected %f, got %f", expectedUnassigned, availableFunds)
			}
		})
	})
}

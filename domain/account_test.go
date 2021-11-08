package domain

import (
	"math"
	"testing"
)

func moneyEqual(a, b Money) bool {
	return math.Abs(float64(a-b)) < 1e-10
}

func TestAccount(t *testing.T) {
	t.Run("AddTransaction", func(t *testing.T) {
		t.Run("adds the transaction to the account", func(t *testing.T) {
			acct := Account{}
			acct.AddTransaction(
				&Transaction{
					amount: Money(123.45),
				})

			if l := len(acct.transactions); l != 1 {
				t.Fatalf("Expected 1 transaction, got %d", l)
			}

			if amt := acct.transactions[0].amount; !moneyEqual(amt, 123.45) {
				t.Fatalf("Expected 123.45, got %f", amt)
			}
		})
	})

	t.Run("Balance", func(t *testing.T) {
		t.Run("returns 0 when no transactions exist", func(t *testing.T) {
			acct := Account{}
			balance := acct.Balance()
			if balance != 0 {
				t.Fatalf("Expected 0, got %f", balance)
			}
		})

		t.Run("returns correct balance when transactions exist", func(t *testing.T) {
			acct := Account{
				transactions: []*Transaction{
					{amount: 123.45},
					{amount: 25.00},
					{amount: -14.75},
				},
			}

			balance := acct.Balance()
			if !moneyEqual(balance, 133.70) {
				t.Fatalf("Expected 133.70, got %f", balance)
			}
		})
	})

	t.Run("Name", func(t *testing.T) {
		t.Run("returns correct account name", func(t *testing.T) {
			acct := Account{name: "test name"}
			name := acct.Name()
			if name != "test name" {
				t.Fatalf("Expected 'test_name', got '%s'", name)
			}
		})
	})
}

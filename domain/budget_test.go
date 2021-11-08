package domain

import "testing"

func TestBudget(t *testing.T) {
	t.Run("AddAccount", func(t *testing.T) {
		t.Run("adds correct account", func(t *testing.T) {
			budget := Budget{}
			budget.AddAccount("account name")
			if l := len(budget.accounts); l != 1 {
				t.Fatalf("Expected 1 account, got %d", l)
			}
			if name := budget.accounts[0].name; name != "account name" {
				t.Fatalf("Expected 'account name', got '%s'", name)
			}
		})

		t.Run("returns error if name is empty", func(t *testing.T) {
			budget := Budget{}
			err := budget.AddAccount("")
			if err == nil {
				t.Fatal("Expected error, got nil")
			}
			if err != ErrMissingName {
				t.Fatalf("Expected ErrMissingName, got '%s'", err.Error())
			}
		})

		t.Run("returns error if name is only whitespace", func(t *testing.T) {
			budget := Budget{}
			err := budget.AddAccount("  ")
			if err == nil {
				t.Fatal("Expected error, got nil")
			}
			if err != ErrMissingName {
				t.Fatalf("Expected ErrMissingName, got '%s'", err.Error())
			}
		})

		t.Run("trims whitespace from name", func(t *testing.T) {
			budget := Budget{}
			err := budget.AddAccount(" account name ")
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if name := budget.accounts[0].name; name != "account name" {
				t.Fatalf("Expected 'account name', got '%s'", name)
			}
		})
	})

	t.Run("AddCategory", func(t *testing.T) {
		t.Run("adds correct category", func(t *testing.T) {
			budget := Budget{}
			err := budget.AddCategory("category name")
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if l := len(budget.categories); l != 1 {
				t.Fatalf("Expected 1 category, got %d", l)
			}
			if name := budget.categories[0].name; name != "category name" {
				t.Fatalf("Expected 'category name', got '%s'", name)
			}
		})

		t.Run("returns error if name is empty", func(t *testing.T) {
			budget := Budget{}
			err := budget.AddCategory("")
			if err == nil {
				t.Fatal("Expected error, got nil")
			}
			if err != ErrMissingName {
				t.Fatalf("Expected ErrMissingName, got '%s'", err.Error())
			}
		})

		t.Run("returns error if name is only whitespace", func(t *testing.T) {
			budget := Budget{}
			err := budget.AddCategory("  ")
			if err == nil {
				t.Fatal("Expected error, got nil")
			}
			if err != ErrMissingName {
				t.Fatalf("Expected ErrMissingName, got '%s'", err.Error())
			}
		})

		t.Run("trims whitespace from name", func(t *testing.T) {
			budget := Budget{}
			err := budget.AddCategory(" category name ")
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if name := budget.categories[0].name; name != "category name" {
				t.Fatalf("Expected 'category name', got '%s'", name)
			}
		})
	})

	t.Run("AddPayee", func(t *testing.T) {
		t.Run("adds correct payee", func(t *testing.T) {
			budget := Budget{}
			err := budget.AddPayee("payee name")
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if l := len(budget.payees); l != 1 {
				t.Fatalf("Expected 1 payee, got %d", l)
			}
			if name := budget.payees[0].name; name != "payee name" {
				t.Fatalf("Expected 'payee name', got '%s'", name)
			}
		})

		t.Run("returns error if name is empty", func(t *testing.T) {
			budget := Budget{}
			err := budget.AddPayee("")
			if err == nil {
				t.Fatal("Expected error, got nil")
			}
			if err != ErrMissingName {
				t.Fatalf("Expected ErrMissingName, got '%s'", err.Error())
			}
		})

		t.Run("returns error if name is only whitespace", func(t *testing.T) {
			budget := Budget{}
			err := budget.AddPayee("  ")
			if err == nil {
				t.Fatal("Expected error, got nil")
			}
			if err != ErrMissingName {
				t.Fatalf("Expected ErrMissingName, got '%s'", err.Error())
			}
		})

		t.Run("trims whitespace from name", func(t *testing.T) {
			budget := Budget{}
			err := budget.AddPayee(" payee name ")
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if name := budget.payees[0].name; name != "payee name" {
				t.Fatalf("Expected 'payee name', got '%s'", name)
			}
		})
	})
}

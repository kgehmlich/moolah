package domain

import (
	"fmt"
	"strings"
)

var ErrMissingName = fmt.Errorf("Name is required")

type MonthlyCategory struct {
	category  Category
	assigned  Money
	used      Money
	available Money
}

type MonthlyBudget struct {
	categories []*MonthlyCategory
	unassigned Money
}

type Budget struct {
	accounts   []*Account
	categories []*Category
	monthlies  []*MonthlyBudget
	payees     []*Payee
}

func (b *Budget) AddAccount(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return ErrMissingName
	}
	b.accounts = append(b.accounts, &Account{name: name})
	return nil
}

func (b *Budget) AddCategory(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return ErrMissingName
	}
	b.categories = append(b.categories, &Category{name})
	return nil
}

func (b *Budget) AddPayee(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return ErrMissingName
	}
	b.payees = append(b.payees, &Payee{name})
	return nil
}

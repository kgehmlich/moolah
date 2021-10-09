package main

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

var ErrDuplicateName = fmt.Errorf("that name already exists")
var ErrNonPositiveAmount = fmt.Errorf("amount must be greater than 0")

type UniqueID string
type Money float64

func generateUniqueID() UniqueID {
	id := strings.ReplaceAll(uuid.NewString(), "-", "")
	return UniqueID(id)
}

type Account struct {
	Name    string
	id      UniqueID
	balance Money
}

func (a *Account) ID() UniqueID {
	return a.id
}

func (a *Account) Balance() Money {
	return a.balance
}

func (a *Account) Credit(amt Money) error {
	if amt <= 0 {
		return ErrNonPositiveAmount
	}
	a.balance += amt
	return nil
}

func (a *Account) Debit(amt Money) error {
	if amt <= 0 {
		return ErrNonPositiveAmount
	}

	a.balance -= amt
	return nil
}

type Category struct {
	Name      string
	id        UniqueID
	available Money
}

func (c *Category) ID() UniqueID {
	return c.id
}

func (c *Category) Available() Money {
	return c.available
}

func (c *Category) Assign(amt Money) error {
	if amt <= 0 {
		return ErrNonPositiveAmount
	}
	c.available += amt
	return nil
}

func (c *Category) Unassign(amt Money) error {
	if amt <= 0 {
		return ErrNonPositiveAmount
	}
	c.available -= amt
	return nil
}

type Budget struct {
	accounts   map[UniqueID]*Account
	categories map[UniqueID]*Category
}

func (b *Budget) Accounts() []*Account {
	accts := []*Account{}
	for _, v := range b.accounts {
		accts = append(accts, v)
	}
	return accts
}

func (b *Budget) AddAccount(name string) error {
	for _, a := range b.accounts {
		if a.Name == name {
			return ErrDuplicateName
		}
	}
	id := generateUniqueID()
	newAcct := &Account{
		Name: name,
		id:   id,
	}
	b.accounts[id] = newAcct
	return nil
}

func (b *Budget) TotalFunds() Money {
	funds := Money(0)
	for _, a := range b.accounts {
		funds += a.Balance()
	}
	return funds
}

func (b *Budget) Categories() []*Category {
	cats := []*Category{}
	for _, v := range b.categories {
		cats = append(cats, v)
	}
	return cats
}

func (b *Budget) AddCategory(name string) error {
	for _, c := range b.categories {
		if c.Name == name {
			return ErrDuplicateName
		}
	}
	id := generateUniqueID()
	newCategory := &Category{
		Name: name,
		id:   id,
	}
	b.categories[id] = newCategory
	return nil
}

func (b *Budget) UnassignedFunds() Money {
	totalFunds := Money(0)
	for _, a := range b.accounts {
		totalFunds += a.Balance()
	}

	assignedFunds := Money(0)
	for _, c := range b.categories {
		assignedFunds += c.Available()
	}

	return totalFunds - assignedFunds
}

func (b *Budget) Deposit(amt Money, acctID UniqueID) {
	acct := b.accounts[acctID]
	acct.Credit(amt)
}

func (b *Budget) Spend(amt Money, acctID, categoryID UniqueID) {
	acct := b.accounts[acctID]
	acct.Debit(amt)

	cat := b.categories[categoryID]
	cat.Unassign(amt)
}

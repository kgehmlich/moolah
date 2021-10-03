package main

import "fmt"

var ErrNonPositiveAmount = fmt.Errorf("amount must be greater than 0")
var ErrInsufficientFunds = fmt.Errorf("insufficient funds")

type Money float64

type Creditable interface {
	Credit(amt Money) error
}

type Debitable interface {
	Debit(amt Money) error
}

type Account struct {
	Name    string
	balance Money
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
	if amt > a.balance {
		return ErrInsufficientFunds
	}

	a.balance -= amt
	return nil
}

type Category struct {
	Name      string
	available Money
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

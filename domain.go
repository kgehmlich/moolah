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

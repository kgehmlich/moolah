package main

import "fmt"

var InsufficientFundsErr = fmt.Errorf("Insufficient funds")

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
	a.balance += amt
	return nil
}

func (a *Account) Debit(amt Money) error {
	if amt > a.balance {
		return InsufficientFundsErr
	}

	a.balance -= amt
	return nil
}
